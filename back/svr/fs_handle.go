package fssvr

import (
	"fmt"

	"github.com/iopig/feed-scale/back/common"
	"github.com/iopig/feed-scale/interface/grpc/go_out/fsapi"
)

type pistyLog struct {
	PigstyId       int
	FedWeight      int64
	StyfirstWeight int64
	LastTime       int64
	StartTime      int64
	FeedCount      int32
	PigNum         uint32
	AverageWeight  uint32
	AdviseWeight   uint32
	StyName        string
	PigSpecies     string
	FarmId         int
}
type ScaleProcess struct {
	//TODO need  add mutex lock
	CurrentWeight  float64 //前一次称上的重量
	DevFirstWeight int
	PigstyData     map[int]*pistyLog
	DevId          string
	FarmId         int32
	LastTime       int64
	StartTime      int64
	FedWeight      float64 //从喂料开始 到当前时间，一共喂料的重量。
	CfgVersion     int32
	FeedCount      int32
}

func (sp *ScaleProcess) UploadRawInfo(in *fsapi.ScaleDevRawData, fedres *fsapi.ResHeader) (err error) {

	//TODO 增加合法性校验，设备是否合法，是否已经开始了一次的喂料过程，距离上次喂料时间有多长。
	//当前重量是否合理，时间是否合理。
	fedres.ErrCode = fsapi.ErrCode_ERR_SUCCESS
	fedres.ErrMsg = ""
	stm, err := common.MysqlDbHandle.Prepare(
		"insert  into feed_log(sty_id,fodder_id,advise_value,device_id,  current_weight,feed_weight,feed_time,cmd_type) " +
			"values(?,?,?,?,  ?,?,from_unixtime(?),?)")
	if err != nil {
		fmt.Println(err)
		fedres.ErrCode = fsapi.ErrCode_ERR_SUCCESS
		fedres.ErrMsg = "operate db prepare error!"
		return
	}
	defer stm.Close()

	CurrentWeight := int(in.CurrentWeight * 1000)

	res, err := stm.Exec(in.PigstyId, 0, 0, sp.DevId,
		CurrentWeight, 0, in.Timestamp/1000, in.FeedType)
	if err != nil {
		fmt.Println(err)
		fedres.ErrCode = fsapi.ErrCode_ERR_SUCCESS
		fedres.ErrMsg = "operate db exe error!"
		return
	}
	_, err = res.LastInsertId()
	if err != nil {
		fmt.Println("insert LastInsertId:", err)
		fedres.ErrCode = fsapi.ErrCode_ERR_SUCCESS
		fedres.ErrMsg = "operate db not insert error!"
		return
	}
	//TODO ,目前同步操作，数据量增大后，加入消息队列中异步计算。
	CalcRawInfo(in, sp)
	return
}
