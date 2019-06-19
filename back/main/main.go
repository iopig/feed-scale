package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"runtime"

	fssvr "github.com/iopig/feed-scale/back/svr"
	"github.com/iopig/feed-scale/interface/grpc/go_out/fsapi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	confFile string // 配置文件路径
)

//定义一个常量：gRPC的端口
const (
	port = ":50051"
)

func initArgs() {
	flag.StringVar(&confFile, "config", "./gateway.json", "where gateway.json is.")
	flag.Parse()
}

func initEnv() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	var (
		//	err error
		grpcServer *grpc.Server
		svrApi     fssvr.SvrApi
	)

	// 初始化环境
	initArgs()
	initEnv()

	go fssvr.StartUdpSvr(&svrApi)

	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatal("fail to listen")
		goto ERR
	}

	grpcServer = grpc.NewServer()

	fsapi.RegisterFsPadServer(grpcServer, &svrApi)

	svrApi.Init()

	reflection.Register(grpcServer)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("fail to server")
	}

	os.Exit(0)

ERR:
	fmt.Fprintln(os.Stderr, err)
	os.Exit(-1)
}
