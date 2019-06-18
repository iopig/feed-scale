package common

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

const (
	TYPE_PAD_LOGIN_REQ   = 0x01
	TYPE_PAD_LOGIN_RES   = 0x02
	TYPE_RAW_DATA_UPLOAD = 0x03
	TYPE_RAW_DATA_RES    = 0x04
)

const (
	AFS_PROTO_HEADER_FIRST  = 0xFE
	AFS_PROTO_HEADER_SECOND = 0xEF
	AFS_PROTO_HEADER_THIRD  = 0x9F
	AFS_PROTO_HEADER_FOURTH = 0xF9
)

func getNextRawData(sep, srcbyte []byte, n int) (RawArr [][]byte) {
	RawArr = bytes.Split(srcbyte, sep)
	return
}

func EncodeAFSProto(rawDataBuff []byte, dataType int8) (resByte []byte, err error) {

	resByte = make([]byte, 0, 8+len(rawDataBuff))
	//同步标志
	resByte = append(resByte, AFS_PROTO_HEADER_FIRST,
		AFS_PROTO_HEADER_SECOND,
		AFS_PROTO_HEADER_THIRD,
		AFS_PROTO_HEADER_FOURTH)
	//报文是否结束 第4位，长度1
	resByte = append(resByte, 0x00)
	//包长度  第5-6位，长度2
	datalen := len(rawDataBuff)
	var buffer bytes.Buffer
	err = binary.Write(&buffer, binary.BigEndian, int16(datalen))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	resByte = append(resByte, buffer.Bytes()...)
	//包类型  第7位，长度1
	// err = binary.Write(&buffer, binary.BigEndian, int16(dataType))
	// if err != nil {
	// 	fmt.Println(err)
	// 	return nil, err
	// }
	//resByte = append(resByte, buffer.Bytes()...)
	resByte = append(resByte, byte(dataType))

	resByte = append(resByte, rawDataBuff...)
	return

}
