package common

import (
	"errors"
	"strconv"
	"strings"
)

//单位映射表
//“注意”：多字节的单位一定放在最前面，如kg放在k和g的前面。
type unitOfWeight struct {
	unitstr    string
	firstMul   int
	secondSize int
}

func getUnitSlice() []unitOfWeight {
	return []unitOfWeight{
		{unitstr: "kg", firstMul: 1000, secondSize: 3}, //
		{unitstr: "KG", firstMul: 1000, secondSize: 3},
		{unitstr: "k", firstMul: 1000, secondSize: 3},
		{unitstr: "K", firstMul: 1000, secondSize: 3},
		{unitstr: "g", firstMul: 1, secondSize: 0},
		{unitstr: "G", firstMul: 1, secondSize: 0},
	}
}

//重量转换为以克为单位的,只能转换常规表示方式如 10kg，10.11 kg，1g，1.1g等
//TODO test
func WeightToNumber(weightStr string) (int, error) {
	var zeroStr = "0000000000"
	if weightStr == "" {
		//重量为“”时，表示为0克
		return 0, nil
	} else {
		//之前有用户把之前的用户.后面字段 转换成数字并加1
		weightStrEx := strings.Split(weightStr, ".")
		if len(weightStrEx) > 2 {
			return 0, errors.New("weight data format error!")
		}

		UnitSlice := getUnitSlice()
		var unitStr string
		var intweight int

		if len(weightStrEx) != 2 {
			//content := weightStr[1:len(weightStr)]
			//return strconv.Atoi(content)
			unitStr = weightStr
		} else {
			unitStr = weightStrEx[1]
			intweight, _ = strconv.Atoi(weightStrEx[0])
		}

		for _, v := range UnitSlice {

			if strings.Contains(unitStr, v.unitstr) {

				if len(weightStrEx) != 2 {
					intweightFirstStr := weightStr[0 : len(weightStr)-len(v.unitstr)]
					intweight, _ = strconv.Atoi(intweightFirstStr)
				}
				//intweight, err := strconv.Atoi(weightStrEx[0])
				intweight = intweight * v.firstMul
				if len(weightStrEx) != 2 {
					return intweight, nil
				}
				intweightSecStr := weightStrEx[1][0 : len(weightStrEx[1])-len(v.unitstr)]
				if v.secondSize <= len(intweightSecStr) {
					intweightSecStr = intweightSecStr[0:v.secondSize]
				} else {
					intweightSecStr = intweightSecStr + zeroStr[0:v.secondSize-len(intweightSecStr)]
				}

				intweightSecond, err := strconv.Atoi(intweightSecStr)
				intweight = intweight + intweightSecond
				return intweight, err

			}
		}

		return 0, nil
	}
}
