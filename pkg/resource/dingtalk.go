package resource

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dingtalkrobot_1_0 "github.com/alibabacloud-go/dingtalk/robot_1_0"
	"github.com/alibabacloud-go/tea/tea"
)

/**
 * 使用 Token 初始化账号Client
 * @return Client
 * @throws Exception
 */
func createClient() (_result *dingtalkrobot_1_0.Client) {
	config := &openapi.Config{}
	config.Protocol = tea.String("https")
	config.RegionId = tea.String("central")
	_result = &dingtalkrobot_1_0.Client{}
	var err error
	_result, err = dingtalkrobot_1_0.NewClient(config)
	if err != nil {
		panic(err)
	}
	return
}
