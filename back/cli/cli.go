package main

import (
	"context"
	"fmt"

	"github.com/iopig/feed-scale/interface/grpc/go_out/fsapi"
	"google.golang.org/grpc"
)

func PadLogin() (err error) {

	cli, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())

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
	PadLogin()
}
