package logic

import (
	"encoding/json"
	"errors"
	"fmt"
	dingtalkrobot_1_0 "github.com/alibabacloud-go/dingtalk/robot_1_0"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"zhipuai_api/internal/common/http"
	"zhipuai_api/internal/config"
	"zhipuai_api/internal/types"
	"zhipuai_api/pkg/resource"
)

type dingAccessTokenBody struct {
	AppKey    string `json:"appKey"`
	AppSecret string `json:"appSecret"`
}

type dingAccessTokenResponse struct {
	AccessToken string `json:"accessToken"`
	ExpireIn    int64  `json:"expireIn"`
}

func GetDingAccessToken() string {
	body := &dingAccessTokenBody{
		AppKey:    config.Conf.Dingtalk.AppKey,
		AppSecret: config.Conf.Dingtalk.AppSecret,
	}
	bodyStr, _ := json.Marshal(body)
	headers := map[string]string{
		"Text-Type": "application/json",
	}
	resp, err := http.HTTPRequest("POST", "https://api.dingtalk.com/v1.0/oauth2/accessToken", headers, bodyStr)
	if err != nil {
		panic(err)
	}
	ans := new(dingAccessTokenResponse)
	err = json.Unmarshal(resp, ans)
	if err != nil {
		panic(err)
	}
	return ans.AccessToken
}

func DingSend(args *types.DingMessageSendRequest) (_err error) {
	token := GetDingAccessToken()
	if token == "" {
		return errors.New("获取钉钉token失败")
	}
	client := resource.DingClient
	privateChatSendHeaders := &dingtalkrobot_1_0.PrivateChatSendHeaders{}
	privateChatSendHeaders.XAcsDingtalkAccessToken = tea.String(token)
	content := struct {
		Content string `json:"content"`
	}{
		Content: args.Text,
	}
	msgParam, _ := json.Marshal(content)
	privateChatSendRequest := &dingtalkrobot_1_0.PrivateChatSendRequest{
		MsgParam:           tea.String(string(msgParam)),
		MsgKey:             tea.String("sampleMarkdown"),
		RobotCode:          tea.String(config.Conf.Dingtalk.AppKey),
		OpenConversationId: tea.String(args.ConversationId),
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
			fmt.Println("code:", tea.StringValue(err.Code), "message:", tea.StringValue(err.Message))
		}

	}
	return _err
}
