package fspub

type DeviceType int32

const (
	FEED_SCALE DeviceType = 0 //饲料称

)

type DeviceInfo struct {
	DevId   string
	DevType DeviceType
}

//猪圈信息
type PigstyInfo struct {
	PigstyId      string
	PigNum        uint32
	AverageWeight uint32
	AdviseWeight  uint32
	LastFed       uint32
	PigIdList     []string
	DaysAge       uint32
	SpeciesID     uint32
	SpeciesName   string
	StyName       string
}

//猪舍信息
type PigHouseInfo struct {
	HouseId        string
	DetailInfo     string
	PigAge         int32
	PigstyInfoList []PigstyInfo
}

//猪场信息
type Farmer struct {
	Id           string
	Devices      []DeviceInfo
	Address      string
	Name         string
	PigHouseList []PigHouseInfo
	Version      int32
}
