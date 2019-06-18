package fssvr

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/iopig/feed-scale/interface/grpc/go_out/fsapi"
)

func StartUdpSvr(srv *SvrApi) {
	//TODO 考虑udp的并发处理数据
	addr, err := net.ResolveUDPAddr("udp", "0.0.0.0:1234")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer conn.Close()

	var udpHandle UdpDataHandle
	udpHandle.Init(srv)
	var in fsapi.ScaleDevRawData
	in.CurrentWeight = 3.99
	//in.FeedType = fsapi.FeedType_WEIGHT_INFO
	in.PigstyId = -1
	in.Timestamp = 1560484862010
	// 对数据进行序列化
	data, err := proto.Marshal(&in)
	if err != nil {
		str := string(data[:])
		fmt.Println(str)

	}
	RawData := make([]byte, 1024)
	for {
		// Here must use make and give the lenth of buffer

		n, rAddr, err := conn.ReadFromUDP(RawData)

		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("%s Received from: %v\n", time.Now().Format("2006-01-02 15:04:05"), rAddr)
		fmt.Printf("UDP Received data:%X \n", RawData[0:n])
		err = udpHandle.HandleData(RawData[0:n], n, rAddr, conn)
		//err = proto.Unmarshal(RawData[4:n], &in)
		if err != nil {
			fmt.Println(err)
			continue

		}

		//	fmt.Println("Received data:", in)

		// upper := strings.ToUpper(strData)
		// _, err = conn.WriteToUDP([]byte(upper), rAddr)
		// if err != nil {
		// 	fmt.Println(err)
		// 	continue
		// }

		// fmt.Println("Send:", upper)
	}
}
