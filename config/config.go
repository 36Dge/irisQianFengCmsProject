package config

import (
	"encoding/json"
	"os"
)

type AppConfig struct {
	AppName    string `json:"app_name"`
	Port       string `json:"port"`
	StaticPath string `json:"static_path"`
	Mode       string `json:"mode"`
}

// 声明结构体对象
var SerConfig AppConfig

func InitConfig() *AppConfig {

	// 目前这个配置文件没有
	file, err := os.Open("/Users/dcw/Documents/go_project/src/irisDemo/QianFengCmsProject/config.json")
	if err != nil {
		panic(err.Error())
	}
	decoder := json.NewDecoder(file)
	conf := AppConfig{}
	err = decoder.Decode(&conf)
	if err != nil {
		panic(err.Error())
	}
	return &conf

}
