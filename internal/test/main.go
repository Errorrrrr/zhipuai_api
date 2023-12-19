package main

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dingtalkrobot_1_0 "github.com/alibabacloud-go/dingtalk/robot_1_0"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"os"
)

/**
 * 使用 Token 初始化账号Client
 * @return Client
 * @throws Exception
 */
func CreateClient() (_result *dingtalkrobot_1_0.Client, _err error) {
	config := &openapi.Config{}
	config.Protocol = tea.String("https")
	config.RegionId = tea.String("central")
	_result = &dingtalkrobot_1_0.Client{}
	_result, _err = dingtalkrobot_1_0.NewClient(config)
	return _result, _err
}

func _main(args []*string) (_err error) {
	client, _err := CreateClient()
	if _err != nil {
		return _err
	}

	privateChatSendHeaders := &dingtalkrobot_1_0.PrivateChatSendHeaders{}
	privateChatSendHeaders.XAcsDingtalkAccessToken = tea.String("<your access token>")
	privateChatSendRequest := &dingtalkrobot_1_0.PrivateChatSendRequest{
		MsgParam:           tea.String("{\"content\":\"钉钉，让进步发生\"}"),
		MsgKey:             tea.String("sampleText"),
		OpenConversationId: tea.String("cid6******=="),
		RobotCode:          tea.String("dinght0mds5sdpwztcxf"),
	}
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		_, _err = client.PrivateChatSendWithOptions(privateChatSendRequest, privateChatSendHeaders, &util.RuntimeOptions{})
		if _err != nil {
			return _err
		}

		return nil
	}()

	if tryErr != nil {
		var err = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			err = _t
		} else {
			err.Message = tea.String(tryErr.Error())
		}
		if !tea.BoolValue(util.Empty(err.Code)) && !tea.BoolValue(util.Empty(err.Message)) {
			// err 中含有 code 和 message 属性，可帮助开发定位问题
		}

	}
	return _err
}

func main() {
	err := _main(tea.StringSlice(os.Args[1:]))
	if err != nil {
		panic(err)
	}
}
