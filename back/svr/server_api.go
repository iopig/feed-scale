package fssvr

import (
	"fmt"
	"strconv"
	"time"

	"github.com/iopig/feed-scale/back/common"
	"github.com/iopig/feed-scale/interface/grpc/go_out/fsapi"
	"golang.org/x/net/context"
)

/*const (
	FEED_CMD_TYPE_LOAD          = 1
	FEED_CMD_TYPE_CHOOSE_PIGSTY = 2
	FEED_CMD_TYPE_WEIGHT_INFO   = 3
)*/

type SvrApi struct {
	//TODO redis for every device id
	ScaleProcessMap map[string]*ScaleProcess
	PigSpecies      map[int]string
	RA              common.RecommendateAlgm
}

func (c *SvrApi) Init() {
	common.MysqlInit(common.MysqlConfig{
		ConnectStr:   "root:0pl,9okm@tcp(192.168.100.102:3306)/fodder?charset=utf8",
		MaxOpenConns: 2000,
		MaxIdleConns: 2000,
	})
	c.ScaleProcessMap = make(map[string]*ScaleProcess)
	c.RA.InitDays()
}

func (c *SvrApi) PadLogin(ctx context.Context, in *fsapi.DevInfoReq) (*fsapi.PigstyInfoRes, error) {

	var famers Farmers
	var pistyInfoRes fsapi.PigstyInfoRes

	//TODO check in.ReqHeader.Ts is right ？

	fminfo, err := famers.GetFarmerInfoByDev(in.ReqHeader.DevId)
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	pistyInfoRes.Farmer = fminfo.Name
	pistyInfoRes.PigFarmId = fminfo.Id
	pistyInfoRes.PigHouseInfo = make([]*fsapi.PigHouseInfo, 0, len(fminfo.PigHouseList))

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

/*func (c *SvrApi) LoadCmd(ctx context.Context, in *fsapi.LoadReq) (*fsapi.ResHeader, error) {

	var ResHeader fsapi.ResHeader
	Devid, err := strconv.Atoi(in.ReqHeader.DevId)
	if c.ScaleProcessMap[Devid] == nil {
		c.ScaleProcessMap[Devid] = &ScaleProcess{
			CurrentWeight: 0,
			FedWeight:     0,
			DevId:         Devid,
		}
	}
	c.ScaleProcessMap[Devid].LoadCmd(in, &ResHeader)
	//当前重量
	return &ResHeader, err
}
func (c *SvrApi) ChoosePigsty(ctx context.Context, in *fsapi.UploadDevDateReq) (*fsapi.CurrentFedRes, error) {
	var CurrentFedRes fsapi.CurrentFedRes

	Devid, err := strconv.Atoi(in.ReqHeader.DevId)
	if c.ScaleProcessMap[Devid] == nil {
		c.ScaleProcessMap[Devid] = &ScaleProcess{
			CurrentWeight: 0,
			FedWeight:     0,
			DevId:         Devid,
		}
	}
	c.ScaleProcessMap[Devid].ChoosePigsty(in, &CurrentFedRes)

	return &CurrentFedRes, err
}*/
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
