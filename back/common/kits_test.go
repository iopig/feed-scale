package common

import (
	"encoding/hex"
	"fmt"
	"log"
	"reflect"
	"testing"
	"time"
)

func TestWalk(t *testing.T) {

	var x float64 = 3.4
	//fmt.Println("type:", reflect.TypeOf(x))
	t.Log("type:", reflect.TypeOf(x))
	t.Logf("type:%s ", reflect.TypeOf(x))
	var WeightStr = "10.11111kg"
	weightNum, _ := WeightToNumber(WeightStr)
	t.Logf(" %s to %d g", WeightStr, weightNum)

	WeightStr = "10.1KG"
	weightNum, _ = WeightToNumber(WeightStr)
	t.Logf(" %s to %d g", WeightStr, weightNum)

	WeightStr = "10KG"
	weightNum, _ = WeightToNumber(WeightStr)
	t.Logf(" %s to %d g", WeightStr, weightNum)

	WeightStr = "10.1G"
	weightNum, _ = WeightToNumber(WeightStr)
	t.Logf(" %s to %d g", WeightStr, weightNum)

	WeightStr = "10G"
	weightNum, _ = WeightToNumber(WeightStr)
	t.Logf(" %s to %d g", WeightStr, weightNum)

	currentTime := time.Now().Unix()
	t.Logf(" %s to %d g", WeightStr, currentTime)

	//	t.Errorf("wrong number of function calls, got %d want %d", len(got), 1)

}
func Compare(a, b []byte) int {
	for i := 0; i < len(a) && i < len(b); i++ {
		switch {
		case a[i] == b[i]:
			break
		case a[i] > b[i]:
			return 1
		case a[i] < b[i]:
			return -1
		}
	}
	switch {
	case len(a) > len(b):
		return 1
	case len(a) < len(b):
		return -1
	}
	return 0
}

func TestByteSplit(t *testing.T) {
	recvDataStr := "FEEF9FF908FFFFFFFFFFFFFFFFFF01115C8FC2F5285C0F40209DEFCEA7B52DFEEF9FF908FFFFFFFFFFFFFFFFFF01115C8FC2F5285C1140209082CFA7B52DFEEF9FF908FFFFFFFFFFFFFFFFFF01111F85EB51B81E124020AA84CFA7B52DFEEF9FF908FFFFFFFFFFFFFFFFFF011152B81E85EB511340208B87CFA7B52DFEEF9FF908FFFFFFFFFFFFFFFFFF01113D0AD7A3703D134020B889CFA7B52DFEEF9FF908FFFFFFFFFFFFFFFFFF011152B81E85EB51194020D1E3CEA7B52DFEEF9FF908FFFFFFFFFFFFFFFFFF0111C3F5285C8FC2164020F9E5CEA7B52DFEEF9FF908FFFFFFFFFFFFFFFFFF01119A9999999999134020A3E8CEA7B52DFEEF9FF908FFFFFFFFFFFFFFFFFF011148E17A14AE47104020C8EACEA7B52DFEEF9FF908FFFFFFFFFFFFFFFFFF0111713D0AD7A3700F4020F1ECCEA7B52D"
	recvData, err := hex.DecodeString(recvDataStr)
	if err != nil {
		log.Fatal("")
	}
	sep := []byte{0xFE, 0xEF, 0x9F, 0xF9}
	// sep := make([]byte, 4)
	// sep[0] = 0xFE
	// sep[1] = 0xEF
	// sep[2] = 0x9F
	// sep[3] = 0xF9

	dataArr := getNextRawData(sep, recvData, len(recvData))
	for i, v := range dataArr {
		t.Logf(" data:%X ", v)
		fmt.Printf("index %d data:%X \n", i, v)
	}
	//	t.Errorf("wrong number of function calls, got %d want %d", len(got), 1)

}

func assginParameter() (ret int) {
	i := 0
	ret = 0
	for i < 3 {
		go func() {
			time.Sleep(1)
			fmt.Println("get value ", ret)
		}()
		i++
	}
	return 1
}
func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func TestInterface2(t *testing.T) {

	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}
