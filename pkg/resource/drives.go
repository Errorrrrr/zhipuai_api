package resource

import (
	dingtalkrobot_1_0 "github.com/alibabacloud-go/dingtalk/robot_1_0"
	"gorm.io/gorm"
	"zhipuai_api/internal/config"
)

var (
	DB         *gorm.DB
	DingClient *dingtalkrobot_1_0.Client
)

func Init(c config.Config) error {
	//DB = getMysqlDB(c.DBConfig)
	DingClient = createClient()
	return nil
}

func Destroy() {
	var err error
	if DB != nil {
		sqlDB, dErr := DB.DB()

		if dErr != nil {
			panic(dErr)
		}
		err = sqlDB.Close()
		if err != nil {
			panic(err)
		}
	}
}
