package fssvr

import (
	"fmt"
	"strconv"
	"time"

	"github.com/iopig/feed-scale/interface/grpc/go_out/fsapi"
	"golang.org/x/net/context"
)

type SvrApi struct {
	//TODO redis for every device id
	ScaleProcessMap map[int]*ScaleProcess
}

func (c *SvrApi) Init() {
	c.ScaleProcessMap = make(map[int]*ScaleProcess)
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
		for _, U := range v.PigstyInfoList {
			var pigstyInfo fsapi.PigstyInfo
			pigHouseInfo.PigstyInfo = append(pigHouseInfo.PigstyInfo, &pigstyInfo)
			pigstyInfo.PigstyId = U.PigstyId
			pigstyInfo.PigNum = U.PigNum
			pigstyInfo.AverageWeight = U.AverageWeight
			//TODO add advise algorithm
			pigstyInfo.AdviseWeight = 5000
			//TODO
			pigstyInfo.LastFed = uint32(time.Now().Unix()) - 3600
			pigstyInfo.PigId = append(pigstyInfo.PigId, "1")
		}

	}

	return &pistyInfoRes, err
}
func (c *SvrApi) LoadCmd(ctx context.Context, in *fsapi.LoadReq) (*fsapi.ResHeader, error) {

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
func (c *SvrApi) ChoosePigsty(ctx context.Context, in *fsapi.ChoosePigstyReq) (*fsapi.CurrentFedRes, error) {
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
}
func (c *SvrApi) UploadRawInfo(ctx context.Context, in *fsapi.ChoosePigstyReq) (*fsapi.CurrentFedRes, error) {
	var CurrentFedRes fsapi.CurrentFedRes

	Devid, err := strconv.Atoi(in.ReqHeader.DevId)
	if c.ScaleProcessMap[Devid] == nil {
		c.ScaleProcessMap[Devid] = &ScaleProcess{
			CurrentWeight: 0,
			FedWeight:     0,
			DevId:         Devid,
		}
	}
	c.ScaleProcessMap[Devid].UploadRawInfo(in, &CurrentFedRes)

	return &CurrentFedRes, err
}
