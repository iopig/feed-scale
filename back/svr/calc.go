package fssvr

import (
	"fmt"
	"strconv"
	"time"

	"github.com/iopig/feed-scale/back/common"
	fspub "github.com/iopig/feed-scale/back/pigpub"
	"github.com/iopig/feed-scale/interface/grpc/go_out/fsapi"
)

func GetPigstysHistoryBySty(farmid int) (pPigstyInfoList *[]*fsapi.PigstyInfo, err error) {

	pigstyInfoList := make([]*fsapi.PigstyInfo, 0, 20)
	rows, err := common.MysqlDbHandle.Query(
		"SELECT sty_id,feed_time,feed_count,   feed_weight, start_time"+
			" FROM feed_res "+
			"where  own_id=?   "+
			"order by sty_id",
		farmid)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	var pigstyId uint32
	var pigNum uint32
	var averageWeight uint32
	var daysage uint32
	var speciesId uint32
	var speciesName string
	var styName string

	for rows.Next() {
		if err := rows.Scan(&pigstyId, &pigNum, &averageWeight,
			&daysage, &speciesId, &styName, &speciesName); err != nil {
			fmt.Println("getPigstys error:")
			fmt.Println(err)
			return nil, err
		}
		pigstyInfoList = append(pigstyInfoList, &fsapi.PigstyInfo{
			PigstyId:      strconv.FormatInt(int64(pigstyId), 10),
			PigNum:        pigNum,
			AverageWeight: averageWeight,
			StyName:       styName,
			//TODO ...
		})
	}
	return
}

func GetPigstysBySty(farmid, styid int) (pi fspub.PigstyInfo, err error) {

	rows, err := common.MysqlDbHandle.Query(
		"SELECT sty_id,live_count,avg_weight,days_age,species_id,sty_name,pigspecies.name as cc"+
			" FROM pigsty_current, pigspecies "+
			"where  sty_id=? and own_id=? and pigspecies.id = pigsty_current.species_id   ",
		styid, farmid)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	var pigstyId uint32
	var pigNum uint32
	var averageWeight uint32
	var daysage uint32
	var speciesId uint32
	var speciesName string
	var styName string

	for rows.Next() {
		if err := rows.Scan(&pigstyId, &pigNum, &averageWeight,
			&daysage, &speciesId, &styName, &speciesName); err != nil {
			fmt.Println("getPigstys error:")
			fmt.Println(err)
			return pi, err
		}
		pi = fspub.PigstyInfo{
			PigstyId:      strconv.FormatInt(int64(pigstyId), 10),
			PigNum:        pigNum,
			AverageWeight: averageWeight,
			DaysAge:       daysage,
			SpeciesID:     speciesId,
			SpeciesName:   speciesName,
			StyName:       styName,
		}
	}
	return
}

func saveResult(pl *pistyLog) (err error) {

	//TODO 增加合法性校验，设备是否合法，是否已经开始了一次的喂料过程，距离上次喂料时间有多长。
	//当前重量是否合理，时间是否合理。

	stm, err := common.MysqlDbHandle.Prepare(
		"replace  into feed_res(sty_id,feed_time,feed_count,   feed_weight, start_time,own_id) " +
			"values(?,from_unixtime(?),?,   ?,from_unixtime(?),?  )")
	if err != nil {
		fmt.Println(err)

		return
	}
	defer stm.Close()

	res, err := stm.Exec(pl.PigstyId, pl.StartTime, pl.FeedCount,
		pl.FedWeight, pl.LastTime, pl.FarmId)
	if err != nil {
		fmt.Println(err)

		return
	}
	_, err = res.LastInsertId()
	if err != nil {
		fmt.Println("insert LastInsertId:", err)

		return
	}

	return
}

func CalcRawInfo(in *fsapi.ScaleDevRawData, sp *ScaleProcess) (err error) {

	//上料
	if in.FeedType == fsapi.FeedType_LOAD {

		//更新总共放在称上的重量
		sp.FedWeight = sp.FedWeight + in.CurrentWeight - sp.CurrentWeight
		//更新当前称重量
		sp.CurrentWeight = in.CurrentWeight
		//间隔时间大于1天，那么为新的喂料
		if common.GetDayInterval(sp.LastTime, time.Now().Unix()) != 0 {
			sp.StartTime = time.Now().Unix()
			sp.LastTime = time.Now().Unix()
			sp.FeedCount = 1
		} else if sp.LastTime-time.Now().Unix() > 3600 { //间隔时间大于1个小时，为下次喂料
			sp.StartTime = time.Now().Unix()
			sp.LastTime = time.Now().Unix()
			sp.FeedCount++
		}

		return
	}
	//喂饲料	"strconv"
	styid, err := strconv.Atoi(in.PigstyId)
	if pd, ok := sp.PigstyData[styid]; ok {

		pd.FedWeight = pd.FedWeight + int64((in.CurrentWeight-sp.CurrentWeight)*1000)

		pd.LastTime = time.Now().Unix()

		//

	} else {
		//第一次喂食
		pi, _ := GetPigstysBySty(int(sp.FarmId), styid)

		sp.PigstyData[styid] = &pistyLog{
			PigstyId:       styid,
			FedWeight:      0,
			StyfirstWeight: int64(in.CurrentWeight * 1000),
			LastTime:       time.Now().Unix(),
			StartTime:      time.Now().Unix(),
			FeedCount:      sp.FeedCount,
			PigNum:         pi.PigNum,
			AverageWeight:  pi.AverageWeight,
			AdviseWeight:   pi.AdviseWeight,
			StyName:        pi.StyName,
			PigSpecies:     pi.SpeciesName,
			FarmId:         int(sp.FarmId),
		}

	}
	//save current data to db
	saveResult(sp.PigstyData[styid])
	return
}
