package config

import (
	"encoding/json"
	"io/ioutil"
	"utils/logs"
)

var M = &configuration{}

func init() {
	// 指定配置文件
	b, err := ioutil.ReadFile("/root/go_learn/utils/config.json")
	if err != nil {
		logs.M.FatallnError(err)
	}
	// 配置文件加载到 configuration 中
	err = json.Unmarshal(b, M)
	if err != nil {
		logs.M.FatallnError(err)
	}
}

type configuration struct {
	Port       string  `json:"port"`
	ConsulAddr string  `json:"consuladdr"`
	Db         db      `json:"db"`
	Redis      redisDB `json:"redis"`
	Mail       mail    `json:"mail"`
}

type db struct {
	DBName   string `json:"dbname"`
	Addr     string `json:"addr"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type redisDB struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
	DBnum    int    `json:"db_num"`
}

type mail struct {
	ServerHost string `json:"serverhost"`
	ServerPort int    `json:"serverport"`
	FromEamil  string `json:"formeamil"`
	Password   string `json:"password"`
}
