package config

import (
	"encoding/json"
	"os"
)

// 服务端配置
type AppConfig struct {
	AppName    string   `json:"app_name"`
	Port       string   `json:"port"`
	StaticPath string   `json:"static_path"`
	Mode       string   `json:"mode"`
	DataBase   DataBase `json:"data_base"`
	Redis      Redis    `json:"redis"`
}

// mysql配置
type DataBase struct {
	Drive    string `json:"drive"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Pwd      string `json:"pwd"`
	Host     string `json:"host"`
	Database string `json:"database"`
}

// Redis配置
type Redis struct {
	NetWork  string `json:"net_work"`
	Addr     string `json:"addr"`
	Port     string `json:"port"`
	Password string `json:"password"`
	Prefix   string `json:"prefix"`
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
