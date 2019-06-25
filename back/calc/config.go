package main

import (
	"encoding/json"
	"io/ioutil"
)

// 程序配置
type Config struct {
	FeedScaleUploadUdpServerAddr string `json:"FeedScaleUploadUdpServerAddr"`
	FeedScaleUploadGrpcPort      string `json:"FeedScaleUploadGrpcPort"`
	DataBaseConnetUrl            string `json:"DataBaseConnetUrl"`
}

var (
	G_config *Config
)

func InitConfig(filename string) (err error) {
	var (
		content []byte
		conf    Config
	)

	if content, err = ioutil.ReadFile(filename); err != nil {
		return
	}

	if err = json.Unmarshal(content, &conf); err != nil {
		return
	}

	G_config = &conf
	return
}
