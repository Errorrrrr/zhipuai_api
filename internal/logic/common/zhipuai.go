package logic

import (
	"encoding/json"
	"fmt"
	"zhipuai_api/internal/common/dto"
	"zhipuai_api/internal/common/http"
	"zhipuai_api/internal/common/utils"
)

func Send(body *dto.AiRequest, model, invokeMethod string) (string, error) {

	token, err := utils.GenToken()
	if err != nil {
		return "", err
	}

	url := "https://open.bigmodel.cn/api/paas/v3/model-api/%s/%s"
	headers := map[string]string{
		"Text-Type":     "application/json",
		"Authorization": token,
	}
	b, _ := json.Marshal(body)
	resp, err := http.HTTPRequest("POST", fmt.Sprintf(url, model, invokeMethod), headers, b)
	if err != nil {
		return "", err
	}
	return string(resp), nil
}
