package config

import "github.com/zeromicro/go-zero/rest"

var Conf *Config

type Config struct {
	Env string `json:"Env"`
	rest.RestConf
	DBConfig
	DingtalkConfig `json:"Dingtalk"`
}

func IsTest() bool {
	return Conf.Env == "test"
}

func IsPre() bool {
	return Conf.Env == "pre"
}

func IsProd() bool {
	return Conf.Env == "prod"
}
func IsDev() bool {
	return Conf.Env == "dev"
}

func SetTimeOut() {
	Conf.RestConf.Timeout = Conf.Timeout
}

func GetConf() **Config {
	return &Conf
}
