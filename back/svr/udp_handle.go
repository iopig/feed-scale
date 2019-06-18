package fssvr

import (
	"bytes"
	"errors"
	"fmt"
	"net"

	"github.com/golang/protobuf/proto"
	"github.com/iopig/feed-scale/back/common"
	"github.com/iopig/feed-scale/interface/grpc/go_out/fsapi"
	"golang.org/x/net/context"
)

type UdpDataHandle struct {
	srv *SvrApi
	sep []byte
}

func (Udh *UdpDataHandle) Init(srv *SvrApi) {
	Udh.srv = srv
	Udh.sep = []byte{common.AFS_PROTO_HEADER_FIRST,
		common.AFS_PROTO_HEADER_SECOND,
		common.AFS_PROTO_HEADER_THIRD,
		common.AFS_PROTO_HEADER_FOURTH}
}

//TODO 暂时不处理 传输过程造成的分包问题
func (Udh *UdpDataHandle) getNextRawData(sep, srcbyte []byte, n int) (RawArr [][]byte) {
	RawArr = bytes.Split(srcbyte, sep)
	return
}
func (Udh *UdpDataHandle) dealPadLog(padLogByte []byte, recvAddr *net.UDPAddr, conn *net.UDPConn) (err error) {
	var in fsapi.DevInfoReq
	err = proto.Unmarshal(padLogByte, &in)
	if err != nil {
		fmt.Println(err)

		return nil
	}
	fmt.Println("dealPadLogin format data:", in)
	var ctx context.Context
	PigstyInfoRes, err := Udh.srv.PadLogin(ctx, &in)
	if err != nil {
		fmt.Println(err)

		return nil
	}
	styInfoByte, err := proto.Marshal(PigstyInfoRes)
	if err != nil {
		fmt.Println(err)

		return nil
	}
	// 封装数据包
	resByte, err := common.EncodeAFSProto(styInfoByte, common.TYPE_PAD_LOGIN_RES)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Printf("senddate data:%X \n", resByte)
	//发送数据包
	datalen, err := conn.WriteToUDP(resByte, recvAddr)
	if err != nil {
		fmt.Println(err)
	}
	if datalen != len(resByte) {
		err = fmt.Errorf("Udp send data length is error,should %d ,real send %d ",
			len(resByte), datalen)
	}

	return
}
func (Udh *UdpDataHandle) dealRawData(padLogByte []byte, recvAddr *net.UDPAddr, conn *net.UDPConn) (err error) {
	var in fsapi.UploadDevDateReq

	err = proto.Unmarshal(padLogByte, &in)
	if err != nil {
		fmt.Printf("ERROR:dealRawData proto Unmarshal %v", err)
		return nil
	}
	fmt.Println("dealRawData format data:", in)
	var ctx context.Context
	ResHeader, err := Udh.srv.UploadRawInfo(ctx, &in)
	if err != nil {
		fmt.Printf("ERROR:dealRawData res: %v", err)
		return nil
	}
	fmt.Println("UploadRawInfo:", ResHeader)
	//发送数据包

	return
}
func (Udh *UdpDataHandle) HandleData(srcbyte []byte, n int, recvAddr *net.UDPAddr, conn *net.UDPConn) (err error) {
	//sep := []byte{0xFE, 0xEF, 0x9F, 0xF9}
	if len(Udh.sep) != 4 || Udh.srv == nil {
		err = errors.New("not init")
		fmt.Println(err)
	}
	dataArr := Udh.getNextRawData(Udh.sep, srcbyte, n)

	for i, v := range dataArr {

		if len(v) < 8 {
			continue
		}
		fmt.Printf("index %d data:%X \n", i, v)
		//split 后，v中数据是字段从包头的第5个字段开始,即，从
		secHeadByte := v[0:4]

		switch secHeadByte[3] {
		case common.TYPE_PAD_LOGIN_REQ:
			Udh.dealPadLog(v[4:], recvAddr, conn)
		case common.TYPE_RAW_DATA_UPLOAD:
			Udh.dealRawData(v[4:], recvAddr, conn)
		default:
			return errors.New("ERROR:Udp data can't find type !")

		}

	}

	return
}
