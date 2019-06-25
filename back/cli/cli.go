package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"net"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/iopig/feed-scale/back/common"
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
			Ts:      uint64(time.Now().Unix()),
		},
	})

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("no error:", farmerInfo)
	//fmt.Println(res.ResHeader.ErrCode)
	return err
}

func UdpPadLogin() (err error) {

	conn, err := net.Dial("udp", "127.0.0.1:1234")

	if err != nil {
		fmt.Println(err.Error())
	}
	defer conn.Close()

	farmerInfoReq := &fsapi.DevInfoReq{
		ReqHeader: &fsapi.ReqHeader{
			Version: 1,
			DevId:   "1",
			Ts:      uint64(time.Now().Unix()),
		},
	}
	databuf, err := proto.Marshal(farmerInfoReq)

	// 封装数据包
	resByte, err := common.EncodeAFSProto(databuf, common.TYPE_PAD_LOGIN_REQ)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	resByte, err = hex.DecodeString("FEEF9FF9002300010A1912103337343661353035313532633630616118C9C0BFCFB72D")
	n, err := conn.Write(resByte)
	//n, err := conn.Write(resByte)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("no error,send len ", n)

	msg := make([]byte, 1024)
	n, err = conn.Read(msg)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("no error,recv :", msg)
	var pistyInfoRes fsapi.PigstyInfoRes
	err = proto.Unmarshal(msg[8:n], &pistyInfoRes)

	fmt.Println("no error,recv :", pistyInfoRes)

	//fmt.Println(res.ResHeader.ErrCode)
	return err
}

func UploadRawInfo() (err error) {
	cli, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())

	if err != nil {
		fmt.Println(err.Error())
	}

	ccCli := fsapi.NewFsPadClient(cli)

	Ssle := []*fsapi.ScaleDevRawData{
		&fsapi.ScaleDevRawData{
			PigstyId:      "1",
			CurrentWeight: 2000,
			FeedType:      3,
			Timestamp:     uint64(time.Now().Unix()),
		},
	}

	// sd := fsapi.ScaleDevRawData{
	// 	PigstyId:      "1",
	// 	CurrentWeight: "2",
	// 	FeedType:      3,
	// 	Timestamp:     1234,
	// }
	// Ssle := make([]*fsapi.ScaleDevRawData, 0, 10)
	// Ssle = append(Ssle, &sd)

	uploadraw, err := ccCli.UploadRawInfo(context.Background(), &fsapi.UploadDevDateReq{
		ReqHeader: &fsapi.ReqHeader{
			Version: 1,
			DevId:   "1",
			Ts:      uint64(time.Now().Unix()),
		},
		DevRawData: Ssle,
	})

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("no error:", uploadraw)

	return err
}

func UploadRawInfoUdp() (err error) {
	conn, err := net.Dial("udp", "127.0.0.1:1234")

	if err != nil {
		fmt.Println(err.Error())
	}
	defer conn.Close()

	Ssle := []*fsapi.ScaleDevRawData{
		&fsapi.ScaleDevRawData{
			PigstyId:      "103",
			CurrentWeight: 2000,
			FeedType:      3,
			Timestamp:     uint64(time.Now().Unix()),
		},
	}

	uploadrawReq := &fsapi.UploadDevDateReq{
		ReqHeader: &fsapi.ReqHeader{
			Version: 1,
			DevId:   "1",
			Ts:      uint64(time.Now().Unix()),
		},
		DevRawData: Ssle,
	}

	databuf, err := proto.Marshal(uploadrawReq)

	// 封装数据包
	resByte, err := common.EncodeAFSProto(databuf, common.TYPE_RAW_DATA_UPLOAD)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	resByte, err = hex.DecodeString("")

	n, err := conn.Write(resByte)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("no error,send len ", n)

	return err
}

func main() {

	/*	a := []int{}
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
		fmt.Println("b:", b)*/
	//	PadLogin()
	fmt.Println("time now day :", time.Now().Unix())
	fmt.Println("time yes day :", time.Unix(1555689619, 0).Day())
	interv := (time.Now().Unix() - 1555689619) / 86400
	fmt.Println("time:", interv)
	UdpPadLogin()
	//UploadRawInfoUdp()
}
