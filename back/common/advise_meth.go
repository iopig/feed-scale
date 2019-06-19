package common

import (
	"fmt"
)

//TODO recommendation algorithm
type RecommendateAlgm struct {
	RecomdByDays     map[int]uint32
	RecomdByWeight   map[int]uint32
	RecomdDailyCount int
}

//GatAdviseWeight 返回以克为单位的重量
func (ra *RecommendateAlgm) GatAdviseWeight(pigNum int, pigAge int) uint32 {
	var AdviseNum uint32 = 2000
	if ra.RecomdByDays == nil || len(ra.RecomdByDays) == 0 {
		ra.InitDays()
		if len(ra.RecomdByDays) == 0 {
			fmt.Println("NOT HAVA RECOMMENDATE WEIGHT!")
			return uint32(pigNum) * AdviseNum / 3
		}
	}
	if v, ok := ra.RecomdByDays[pigAge]; ok {
		AdviseNum = v

	} else {
		AdviseNum = 3800
		fmt.Println("AdviseNum not get,set default :", AdviseNum)
	}

	return uint32(pigNum) * AdviseNum / 3
}

//TODO ,推荐的参考数据应该存放在redis里。
func (ra *RecommendateAlgm) InitDays() error {

	//TODO get data from db
	ra.RecomdDailyCount = 3

	rows, err := MysqlDbHandle.Query(
		"SELECT  days_age,avg_value" +
			" FROM advise_by_days ")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer rows.Close()

	var ages int
	var averageValues uint32
	if ra.RecomdByDays == nil {
		ra.RecomdByDays = make(map[int]uint32)
	}
	for rows.Next() {
		if err := rows.Scan(&ages, &averageValues); err != nil {
			fmt.Println("getPigstys error:")
			fmt.Println(err)
			return err
		}
		ra.RecomdByDays[ages] = averageValues
	}
	return nil
}
