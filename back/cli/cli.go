package main

import (
	"context"
	"fmt"

	"github.com/iopig/feed-scale/interface/grpc/go_out/fsapi"
	"google.golang.org/grpc"
)

func PadLogin() (err error) {

	cli, err := grpc.Dial("117.139.13.149:50051", grpc.WithInsecure())

	if err != nil {
		fmt.Println(err.Error())
	}

	ccCli := fsapi.NewFsPadClient(cli)

	farmerInfo, err := ccCli.PadLogin(context.Background(), &fsapi.DevInfoReq{
		ReqHeader: &fsapi.ReqHeader{
			Version: 1,
			DevId:   "1",
			Ts:      uint32(123456789),
		},
	})

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("no error:", farmerInfo)
	//fmt.Println(res.ResHeader.ErrCode)
	return err

}
func LoadCmd() (err error) {

	cli, err := grpc.Dial("117.139.13.149:50051", grpc.WithInsecure())

	if err != nil {
		fmt.Println(err.Error())
	}

	ccCli := fsapi.NewFsPadClient(cli)

	resHeader, err := ccCli.LoadCmd(context.Background(), &fsapi.LoadReq{
		ReqHeader: &fsapi.ReqHeader{
			Version: 1,
			DevId:   "1",
			Ts:      uint32(1559786077),
		},
		CurrentWeight: "100.11kg",
	})

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("no error:", resHeader)
	//fmt.Println(res.ResHeader.ErrCode)
	return err

}
func main() {

	a := []int{}
	a = nil

	b := append([]int(nil), a...)
	if b == nil {
		fmt.Println("b is nil :", b)
	}
	fmt.Println("b:", b)
	b = append(a[:0:0], a...)
	if b == nil {
		fmt.Println("b is nil :", b)
	}
	fmt.Println("b:", b)
	//PadLogin()

	var weightstr = "1011"
	fmt.Println("weight:", weightstr[0:3])
	fmt.Println("weight:", weightstr[0:4])
	fmt.Println("weight:", weightstr[1:3])
	fmt.Println("weight:", weightstr[1:4])
	fmt.Println("weight:", weightstr[2:4])
	fmt.Println("weight:", weightstr[2:3])

	LoadCmd()
}
