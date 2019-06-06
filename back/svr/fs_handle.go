package fssvr

import (
	"fmt"
	"strconv"

	"github.com/iopig/feed-scale/back/common"
	"github.com/iopig/feed-scale/interface/grpc/go_out/fsapi"
)

type ScaleProcess struct {
	//TODO need  add mutex lock
	CurrentWeight int
	PigstyId      int
	FedWeight     int
	DevId         int
}

const (
	FEED_CMD_TYPE_LOAD          = 1
	FEED_CMD_TYPE_CHOOSE_PIGSTY = 2
	FEED_CMD_TYPE_WEIGHT_INFO   = 3
)

func (sp *ScaleProcess) LoadCmd(in *fsapi.LoadReq, resHeader *fsapi.ResHeader) (err error) {

	//TODO 增加合法性校验，设备是否合法，是否已经开始了一次的喂料过程，距离上次喂料时间有多长。
	//当前重量是否合理，时间是否合理。

	resHeader.ErrCode = fsapi.ErrCode_ERR_SUCCESS
	resHeader.ErrMsg = ""
	stm, err := common.MysqlDbHandle.Prepare(
		"insert  into feed_log(sty_id,fodder_id,advise_value,device_id, current_weight, feed_weight,feed_time,cmd_type) " +
			"values(?,?,?,?,?,?,?,?)")
	if err != nil {
		fmt.Println(err)
		resHeader.ErrCode = fsapi.ErrCode_ERR_SUCCESS
		resHeader.ErrMsg = "operate db prepare error!"
		return
	}
	defer stm.Close()
	CurrentWeight, err := common.WeightToNumber(in.CurrentWeight)

	res, err := stm.Exec(0, 0, 0, in.ReqHeader.DevId,
		CurrentWeight, 0, in.ReqHeader.Ts, FEED_CMD_TYPE_LOAD)
	if err != nil {
		fmt.Println(err)
		resHeader.ErrCode = fsapi.ErrCode_ERR_SUCCESS
		resHeader.ErrMsg = "operate db exe error!"
		return
	}
	_, err = res.LastInsertId()
	if err != nil {
		fmt.Println("insert LastInsertId:", err)
		resHeader.ErrCode = fsapi.ErrCode_ERR_SUCCESS
		resHeader.ErrMsg = "operate db not insert error!"
		return
	}
	sp.CurrentWeight = CurrentWeight
	return
}

func (sp *ScaleProcess) ChoosePigsty(in *fsapi.ChoosePigstyReq, fedres *fsapi.CurrentFedRes) (err error) {

	//TODO 增加合法性校验，设备是否合法，是否已经开始了一次的喂料过程，距离上次喂料时间有多长。
	//当前重量是否合理，时间是否合理。
	resHeader := fedres.ResHeader
	resHeader.ErrCode = fsapi.ErrCode_ERR_SUCCESS
	resHeader.ErrMsg = ""
	stm, err := common.MysqlDbHandle.Prepare(
		"insert  into feed_log(sty_id,fodder_id,advise_value,device_id,  current_weight,feed_weight,feed_time,cmd_type) " +
			"values(?,?,?,?,  ?,?,from_unixtime(?),?)")
	if err != nil {
		fmt.Println(err)
		resHeader.ErrCode = fsapi.ErrCode_ERR_SUCCESS
		resHeader.ErrMsg = "operate db prepare error!"
		return
	}
	defer stm.Close()

	CurrentWeight, err := common.WeightToNumber(in.CurrentWeight)

	res, err := stm.Exec(in.PigstyId, 0, 0, in.ReqHeader.DevId,
		CurrentWeight, 0, in.ReqHeader.Ts, FEED_CMD_TYPE_CHOOSE_PIGSTY)
	if err != nil {
		fmt.Println(err)
		resHeader.ErrCode = fsapi.ErrCode_ERR_SUCCESS
		resHeader.ErrMsg = "operate db exe error!"
		return
	}
	_, err = res.LastInsertId()
	if err != nil {
		fmt.Println("insert LastInsertId:", err)
		resHeader.ErrCode = fsapi.ErrCode_ERR_SUCCESS
		resHeader.ErrMsg = "operate db not insert error!"
		return
	}
	sp.FedWeight = sp.CurrentWeight - CurrentWeight
	sp.PigstyId, err = strconv.Atoi(in.PigstyId)
	fedres.FedWeight = strconv.Itoa(CurrentWeight)
	fedres.PigstyId = in.PigstyId
	sp.CurrentWeight = CurrentWeight
	return
}

func (sp *ScaleProcess) UploadRawInfo(in *fsapi.ChoosePigstyReq, fedres *fsapi.CurrentFedRes) (err error) {

	//TODO 增加合法性校验，设备是否合法，是否已经开始了一次的喂料过程，距离上次喂料时间有多长。
	//当前重量是否合理，时间是否合理。
	resHeader := fedres.ResHeader
	resHeader.ErrCode = fsapi.ErrCode_ERR_SUCCESS
	resHeader.ErrMsg = ""
	stm, err := common.MysqlDbHandle.Prepare(
		"insert  into feed_log(sty_id,fodder_id,advise_value,device_id,  current_weight,feed_weight,feed_time,cmd_type) " +
			"values(?,?,?,?,  ?,?,from_unixtime(?),?)")
	if err != nil {
		fmt.Println(err)
		resHeader.ErrCode = fsapi.ErrCode_ERR_SUCCESS
		resHeader.ErrMsg = "operate db prepare error!"
		return
	}
	defer stm.Close()

	CurrentWeight, err := common.WeightToNumber(in.CurrentWeight)

	res, err := stm.Exec(in.PigstyId, 0, 0, in.ReqHeader.DevId,
		CurrentWeight, 0, in.ReqHeader.Ts, FEED_CMD_TYPE_CHOOSE_PIGSTY)
	if err != nil {
		fmt.Println(err)
		resHeader.ErrCode = fsapi.ErrCode_ERR_SUCCESS
		resHeader.ErrMsg = "operate db exe error!"
		return
	}
	_, err = res.LastInsertId()
	if err != nil {
		fmt.Println("insert LastInsertId:", err)
		resHeader.ErrCode = fsapi.ErrCode_ERR_SUCCESS
		resHeader.ErrMsg = "operate db not insert error!"
		return
	}
	sp.FedWeight = sp.CurrentWeight - CurrentWeight
	sp.PigstyId, err = strconv.Atoi(in.PigstyId)
	fedres.FedWeight = strconv.Itoa(CurrentWeight)
	fedres.PigstyId = in.PigstyId
	sp.CurrentWeight = CurrentWeight
	return
}
