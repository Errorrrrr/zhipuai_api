package config

type DBConfig struct {
	Mysql MysqlCfg `json:"Mysql"`
}
type MysqlCfg struct {
	Host               string `json:"host"`
	Port               int    `json:"port"`
	DBName             string `json:"db_name"`
	Driver             string `json:"driver"`
	UserName           string `json:"user_name"`
	Password           string `json:"password"`
	MaxIdleConnections int    `json:"max_idle_connections"`
	MaxOpenConnections int    `json:"max_open_connections"`
	MaxLifetime        int    `json:"max_lifetime"`
	DBDebug            bool   `json:"db_debug"`
}

type DingtalkConfig struct {
	Dingtalk DingtalkCfg `json:"Dingtalk"`
}
type DingtalkCfg struct {
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
}
