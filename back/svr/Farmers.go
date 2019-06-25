package fssvr

import (
	"fmt"
	"strconv"

	"github.com/iopig/feed-scale/back/common"
	fspub "github.com/iopig/feed-scale/back/pigpub"
)

type Farmers struct {
}

func (fms *Farmers) getFarmerByDev(devId string, farmer *fspub.Farmer) (err error) {

	rows, err := common.MysqlDbHandle.Query(
		"SELECT farmer.id,farmer.address,farmer.name,farmer.ver FROM device,farmer where device.id=? and device.owner_id=farmer.id  ", devId)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&farmer.Id, &farmer.Address, &farmer.Name, &farmer.Version); err != nil {
			fmt.Println("getFarmerByDev error:")
			fmt.Println(err)
			return err
		}

	}
	return
}

func (fms *Farmers) getPighouses(farmer *fspub.Farmer) (err error) {

	rows, err := common.MysqlDbHandle.Query(
		"SELECT id,detail, unix_timestamp(pig_birthday) FROM pighouse where  owner_id=?   ", farmer.Id)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	var houseid uint32
	var detail string
	var birthday int32

	for rows.Next() {
		if err := rows.Scan(&houseid, &detail, &birthday); err != nil {
			fmt.Println("getPighouses error:")
			fmt.Println(err)
			return err
		}
		farmer.PigHouseList = append(farmer.PigHouseList, fspub.PigHouseInfo{
			HouseId:    strconv.FormatInt(int64(houseid), 10),
			DetailInfo: detail,
			PigAge:     birthday,
		})
	}
	return
}

func (fms *Farmers) getPigstys(houseInfo *fspub.PigHouseInfo) (err error) {

	rows, err := common.MysqlDbHandle.Query(
		"SELECT sty_id,live_count,avg_weight,days_age,species_id,sty_name,pigspecies.name as cc"+
			" FROM pigsty_current, pigspecies "+
			"where  house_id=? and pigspecies.id = pigsty_current.species_id   "+
			"order by sty_id",
		houseInfo.HouseId)
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
			return err
		}
		houseInfo.PigstyInfoList = append(houseInfo.PigstyInfoList, fspub.PigstyInfo{
			PigstyId:      strconv.FormatInt(int64(pigstyId), 10),
			PigNum:        pigNum,
			AverageWeight: averageWeight,
			DaysAge:       daysage,
			SpeciesID:     speciesId,
			SpeciesName:   speciesName,
			StyName:       styName,
		})
	}
	return
}

func (fms *Farmers) GetFarmerInfoByDev(devId string) (farmerInfo *fspub.Farmer, err error) {
	var farmer fspub.Farmer

	fmt.Println("DEVID :", devId)
	// 获取猪场主信息
	if err = fms.getFarmerByDev(devId, &farmer); err != nil {
		fmt.Println("getFarmerByDev error:", err)
		return nil, err
	}
	// 获取猪舍信息
	if err = fms.getPighouses(&farmer); err != nil {
		fmt.Println("getFarmerByDev error:", err)
		return nil, err
	}
	// 获取猪圈信息
	for i, _ := range farmer.PigHouseList {
		if err = fms.getPigstys(&farmer.PigHouseList[i]); err != nil {
			fmt.Println("getFarmerByDev error:", err)
			return nil, err
		}
	}
	fmt.Println("farmer :", farmerInfo)
	return &farmer, nil
}
