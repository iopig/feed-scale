package common

import (
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
