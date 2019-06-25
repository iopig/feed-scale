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

type UdpCache struct {
	AddrString  string
	ProcessStep int
	Cache       []byte
}
type UdpDataHandle struct {
	srv         *SvrApi
	sep         []byte
	cacheByAddr map[string]UdpCache
}

func (Udh *UdpDataHandle) Init(srv *SvrApi) {
	Udh.srv = srv
	Udh.sep = []byte{common.AFS_PROTO_HEADER_FIRST,
		common.AFS_PROTO_HEADER_SECOND,
		common.AFS_PROTO_HEADER_THIRD,
		common.AFS_PROTO_HEADER_FOURTH}
	Udh.cacheByAddr = make(map[string]UdpCache)
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

		return err
	}
	fmt.Println("dealPadLogin format data:", in)
	var ctx context.Context
	PigstyInfoRes, err := Udh.srv.PadLogin(ctx, &in)
	if err != nil || PigstyInfoRes.ReqHeader == nil {
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
		return err
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
	//分包现象，如果这次收上次的分包，那么删除上次分包
	if ac, ok := Udh.cacheByAddr[recvAddr.String()]; ok {
		ac.ProcessStep++
		Udh.cacheByAddr[recvAddr.String()] = ac
		defer delete(Udh.cacheByAddr, recvAddr.String())
	}

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
			if err = Udh.dealRawData(v[4:], recvAddr, conn); err != nil && i == len(dataArr)-1 {
				fmt.Printf("record cache %s \n", recvAddr.String())
				ac := UdpCache{
					AddrString:  recvAddr.String(),
					ProcessStep: 0,
				}
				ac.Cache = make([]byte, 0, 2*len(v))
				ac.Cache = append(ac.Cache, v...)
				Udh.cacheByAddr[recvAddr.String()] = ac
				fmt.Printf("record cache %s \n", ac.AddrString)
			}
			//return
		default:
			if ac, ok := Udh.cacheByAddr[recvAddr.String()]; ok && i == 0 {
				fmt.Printf("ac.ProcessStep:%X \n", ac.ProcessStep)
				if ac.ProcessStep == 1 {

					ac.Cache = append(ac.Cache, v...)
					fmt.Printf("ac.Cache:%X \n", ac.Cache)
					switch ac.Cache[3] {
					case common.TYPE_PAD_LOGIN_REQ:
						Udh.dealPadLog(ac.Cache[4:], recvAddr, conn)
					case common.TYPE_RAW_DATA_UPLOAD:
						Udh.dealRawData(ac.Cache[4:], recvAddr, conn)
					default:
						return errors.New("ERROR:Udp data can't find type !")
					}
				}
			} else {
				return errors.New("ERROR:Udp data can't find type !")
			}

		}

	}

	return
}
