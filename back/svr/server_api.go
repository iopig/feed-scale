package fssvr

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/iopig/feed-scale/back/common"
	fspub "github.com/iopig/feed-scale/back/pigpub"
	"github.com/iopig/feed-scale/interface/grpc/go_out/fsapi"
	"golang.org/x/net/context"
)

/*const (
	FEED_CMD_TYPE_LOAD          = 1
	FEED_CMD_TYPE_CHOOSE_PIGSTY = 2
	FEED_CMD_TYPE_WEIGHT_INFO   = 3
)*/

type SvrApi struct {
	//TODO 这些缓存数据应该放在redis中
	ScaleProcessMap map[string]*ScaleProcess
	//PigSpecies      map[int]string
	RA common.RecommendateAlgm
	//下方猪场配置延时时间，单位为秒
	StyCfgDelay int64
}

func (c *SvrApi) Init(mysqlUrl string) {
	common.MysqlInit(common.MysqlConfig{
		//ConnectStr:   "root:0pl,9okm@tcp(192.168.100.102:3306)/fodder?charset=utf8",
		ConnectStr:   mysqlUrl,
		MaxOpenConns: 2000,
		MaxIdleConns: 2000,
	})
	c.ScaleProcessMap = make(map[string]*ScaleProcess)
	c.StyCfgDelay = 3600
	c.RA.InitDays()
	dbint, _ := common.GetOneIntegerFromDB("select sty_cfg_delay from sys_config")
	if dbint != 0 {
		c.StyCfgDelay = dbint
	}
}

func (c *SvrApi) PadLogin(ctx context.Context, in *fsapi.DevInfoReq) (*fsapi.PigstyInfoRes, error) {

	var famers Farmers
	var pistyInfoRes fsapi.PigstyInfoRes
	var sp *ScaleProcess

	var rh fsapi.ReqHeader

	pistyInfoRes.ReqHeader = &rh

	//TODO check in.ReqHeader.Ts is right ？

	if c.ScaleProcessMap[in.ReqHeader.DevId] == nil {
		sp = &ScaleProcess{
			CurrentWeight: 0,
			FedWeight:     0,
			LastTime:      time.Now().Unix(),
			DevId:         in.ReqHeader.DevId,
		}
		c.ScaleProcessMap[in.ReqHeader.DevId] = sp

	} else {
		sp = c.ScaleProcessMap[in.ReqHeader.DevId]
		if in.ConfigVersion != 0 {
			if c.ScaleProcessMap[in.ReqHeader.DevId].CfgVersion != in.ConfigVersion {

			} else if time.Now().Unix()-c.ScaleProcessMap[in.ReqHeader.DevId].LastTime < c.StyCfgDelay {
				return nil, errors.New("Not enough idle time for farmer config !")
			}
		}
	}

	fminfo, err := famers.GetFarmerInfoByDev(in.ReqHeader.DevId)
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	pistyInfoRes.Farmer = fminfo.Name
	pistyInfoRes.PigFarmId = fminfo.Id
	pistyInfoRes.PigHouseInfo = make([]*fsapi.PigHouseInfo, 0, len(fminfo.PigHouseList))
	pistyInfoRes.ConfigVersion = fminfo.Version
	sp.CfgVersion = fminfo.Version
	farmid, _ := strconv.Atoi(fminfo.Id)
	c.ScaleProcessMap[in.ReqHeader.DevId].FarmId = int32(farmid)

	for _, v := range fminfo.PigHouseList {
		var pigHouseInfo fsapi.PigHouseInfo
		pistyInfoRes.PigHouseInfo = append(pistyInfoRes.PigHouseInfo, &pigHouseInfo)
		pigHouseInfo.HouseId = v.HouseId
		pigHouseInfo.PigstyInfo = make([]*fsapi.PigstyInfo, 0, len(v.PigstyInfoList))

		pigAge := (time.Now().Unix() - int64(v.PigAge)) / 86400
		pigHouseInfo.AGE = strconv.FormatInt(pigAge, 10)

		for _, U := range v.PigstyInfoList {

			var pigstyInfo fsapi.PigstyInfo
			pigHouseInfo.PigstyInfo = append(pigHouseInfo.PigstyInfo, &pigstyInfo)
			pigstyInfo.PigstyId = U.PigstyId
			pigstyInfo.PigNum = U.PigNum
			pigstyInfo.AverageWeight = U.AverageWeight
			//TODO add advise algorithm
			pigstyInfo.AdviseWeight = c.RA.GatAdviseWeight(int(U.PigNum), int(pigAge))
			//TODO

			pigstyInfo.LastFed = uint64(time.Now().Unix())
			pigstyInfo.PigId = append(pigstyInfo.PigId, "1")
			pigstyInfo.StyName = U.StyName
			pigstyInfo.PigSpecies = U.SpeciesName
		}

	}

	return &pistyInfoRes, err
}

func (c *SvrApi) UploadRawInfo(ctx context.Context, in *fsapi.UploadDevDateReq) (*fsapi.ResHeader, error) {
	fmt.Println("upload raw info func is called ")
	var CurrentFedRes fsapi.ResHeader

	//devId, err := strconv.Atoi(in.ReqHeader.DevId)
	if c.ScaleProcessMap[in.ReqHeader.DevId] == nil {
		c.ScaleProcessMap[in.ReqHeader.DevId] = &ScaleProcess{
			CurrentWeight: 0,
			FedWeight:     0,
			DevId:         in.ReqHeader.DevId,
		}
	}
	for _, v := range in.DevRawData {
		err := c.ScaleProcessMap[in.ReqHeader.DevId].UploadRawInfo(v, &CurrentFedRes)
		if err != nil {
			fmt.Println(err)
			return nil, nil
		}
	}

	return &CurrentFedRes, nil
}
func (c *SvrApi) GetLast(ctx context.Context, in *fsapi.UploadDevDateReq) (*fsapi.ResHeader, error) {
	fmt.Println("upload raw info func is called ")
	var CurrentFedRes fsapi.ResHeader

	//devId, err := strconv.Atoi(in.ReqHeader.DevId)
	if c.ScaleProcessMap[in.ReqHeader.DevId] == nil {
		c.ScaleProcessMap[in.ReqHeader.DevId] = &ScaleProcess{
			CurrentWeight: 0,
			FedWeight:     0,
			LastTime:      time.Now().Unix(),
			DevId:         in.ReqHeader.DevId,
		}
	}
	//
	for _, v := range in.DevRawData {
		err := c.ScaleProcessMap[in.ReqHeader.DevId].UploadRawInfo(v, &CurrentFedRes)
		if err != nil {
			fmt.Println(err)
			return nil, nil
		}
	}
	c.ScaleProcessMap[in.ReqHeader.DevId].LastTime = time.Now().Unix()
	//

	return &CurrentFedRes, nil
}

func (c *SvrApi) GetHistoryFeedLog(in *fsapi.HistoryReq, res fsapi.FsPad_GetHistoryFeedLogServer) error {
	var famers Farmers
	var farmer fspub.Farmer

	famers.GetFarmerByDev(in.ReqHeader.DevId, &farmer)
	farm_id, err := strconv.Atoi(farmer.Id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	pPigstyInfoList, err := GetPigstysHistoryBySty(farm_id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	for _, v := range *pPigstyInfoList {
		res.Send(v)
	}

	return err
}
