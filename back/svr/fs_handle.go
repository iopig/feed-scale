package fssvr

import (
	"fmt"

	"github.com/iopig/feed-scale/back/common"
	"github.com/iopig/feed-scale/interface/grpc/go_out/fsapi"
)

type pistyLog struct {
	PigstyId       int
	FedWeight      int
	StyfirstWeight int
	LastTime       int
	StartTime      int
}
type ScaleProcess struct {
	//TODO need  add mutex lock
	CurrentWeight  int
	DevFirstWeight int
	PigstyData     map[int]*pistyLog
	DevId          int
	LastTime       int
	StartTime      int
	FedWeight      int //从喂料开始 到当前时间，一共喂料的重量。

}

/*func (sp *ScaleProcess) LoadCmd(in *fsapi.LoadReq, resHeader *fsapi.ResHeader) (err error) {

	//TODO 增加合法性校验，设备是否合法，是否已经开始了一次的喂料过程，距离上次喂料时间有多长。
	//当前重量是否合理，时间是否合理。

	resHeader.ErrCode = fsapi.ErrCode_ERR_SUCCESS
	resHeader.ErrMsg = ""
	stm, err := common.MysqlDbHandle.Prepare(
		"insert  into feed_log(sty_id,fodder_id,advise_value,device_id, current_weight, feed_weight,feed_time,cmd_type) " +
			"values(?,?,?,?,?,?,from_unixtime(?),?)")
	if err != nil {
		fmt.Println(err)
		resHeader.ErrCode = fsapi.ErrCode_ERR_SUCCESS
		resHeader.ErrMsg = "operate db prepare error!"
		return
	}
	defer stm.Close()
	CurrentWeight, err := common.WeightToNumber(in.CurrentWeight)

	res, err := stm.Exec(0, 0, 0, in.ReqHeader.DevId,
		CurrentWeight, 0, in.ReqHeader.Ts, fsapi.FeedType_LOAD)
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
	sp.DevFirstWeight = CurrentWeight
	sp.StartTime = int(time.Now().Unix())
	sp.LastTime = int(time.Now().Unix())
	return
}

func (sp *ScaleProcess) chooseStyLog(in *fsapi.UploadDevDateReq, fedres *fsapi.CurrentFedRes) (err error) {

	stm, err := common.MysqlDbHandle.Prepare(
		"insert  into feed_log(sty_id,fodder_id,advise_value,device_id,  current_weight,feed_weight,feed_time,cmd_type) " +
			"values(?,?,?,?,  ?,?,from_unixtime(?),?)")
	if err != nil {
		fmt.Println(err)
		fmt.Println("operate db prepare error!")
		return
	}
	defer stm.Close()

	CurrentWeight, err := common.WeightToNumber(in.DevRawData[1].CurrentWeight)

	res, err := stm.Exec(in.DevRawData[0].PigstyId, 0, 0, in.ReqHeader.DevId,
		CurrentWeight, 0, in.ReqHeader.Ts, fsapi.FeedType_CHOOSE_PIGSTY)
	if err != nil {
		fmt.Println(err)
		fmt.Println("operate db exe error!")
		return
	}
	_, err = res.LastInsertId()
	if err != nil {
		fmt.Println("insert LastInsertId:", err)
		fmt.Println("operate db not insert error!")
		return
	}

	return
}*/

//calcScaleData 函数完成以下计算：
//1.计算推荐重量
//2.计算已经喂的重量
//3.计算饲料盆里还剩余多少重量
/*func (sp *ScaleProcess) calcScaleData(in *fsapi.UploadDevDateReq, fedres *fsapi.CurrentFedRes) (err error) {

	//TODO 当前重量是否合理，时间是否合理。

	//计算重量

	PigstyId, err := strconv.Atoi(in.DevRawData[0].PigstyId)
	CurrentWeight, err := common.WeightToNumber(in.DevRawData[1].CurrentWeight)

	if sp.PigstyData[PigstyId] == nil {
		sp.PigstyData[PigstyId] = &pistyLog{
			PigstyId:       PigstyId,
			FedWeight:      0,
			StyfirstWeight: CurrentWeight,
			LastTime:       int(time.Now().Unix()),
			StartTime:      int(time.Now().Unix()),
		}
	}
	sp.PigstyData[PigstyId].LastTime = int(time.Now().Unix())
	sp.FedWeight = sp.PigstyData[PigstyId].StyfirstWeight
	fedres.FedWeight = strconv.Itoa(sp.FedWeight)

	fedres.PigstyId = in.DevRawData[0].PigstyId
	sp.CurrentWeight = CurrentWeight
	sp.LastTime = int(time.Now().Unix())
	return
}*/

/*func (sp *ScaleProcess) ChoosePigsty(in *fsapi.UploadDevDateReq, fedres *fsapi.CurrentFedRes) (err error) {

	//TODO 增加合法性校验，设备是否合法，是否已经开始了一次的喂料过程，距离上次喂料时间有多长。
	//当前重量是否合理，时间是否合理。
	resHeader := fedres.ResHeader
	resHeader.ErrCode = fsapi.ErrCode_ERR_SUCCESS
	resHeader.ErrMsg = ""

	//记录原始数据
	err = sp.chooseStyLog(in, fedres)
	if err != nil {
		resHeader.ErrCode = fsapi.ErrCode_ERR_FAILED
		resHeader.ErrMsg = "db operator fail"
		return
	}

	//计算重量

	err = sp.calcScaleData(in, fedres)
	if err != nil {
		resHeader.ErrCode = fsapi.ErrCode_ERR_FAILED
		resHeader.ErrMsg = "calc scale data  fail"
		return
	}

	//记录 “计算结果” 到数据

	return
}*/

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

	CurrentWeight := int(in.CurrentWeight)

	res, err := stm.Exec(in.PigstyId, 0, 0, sp.DevId,
		CurrentWeight, 0, in.Timestamp, in.FeedType)
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
	// sp.FedWeight = sp.CurrentWeight - CurrentWeight
	//sp., err = strconv.Atoi(in.PigstyId)

	//fedres = strconv.Itoa(CurrentWeight)
	//fedres = in.DevRawData.PigstyId
	// sp.FedWeight = CurrentWeight
	// PigstyId, err := strconv.Atoi(in.PigstyId)
	// sp.PigstyData[PigstyId].PigstyId = PigstyId
	// sp.CurrentWeight = CurrentWeight
	// sp.LastTime = int(time.Now().Unix())
	return
}
