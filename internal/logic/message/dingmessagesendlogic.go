package message

import (
	"context"
	"encoding/json"
	"errors"
	"strings"
	"zhipuai_api/internal/common/dto"
	"zhipuai_api/internal/common/enum"
	logic "zhipuai_api/internal/logic/common"
	"zhipuai_api/internal/svc"
	"zhipuai_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DingMessageSendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDingMessageSendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DingMessageSendLogic {
	return &DingMessageSendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DingMessageSendLogic) DingMessageSend(req *types.DingMessageSendRequest) (err error) {
	if req.Text.Content == "" {
		return errors.New("content is empty")
	}
	prompt := []*dto.Prompt{
		{
			RoleName: enum.User,
			Content:  "使用Markdown格式进行回复消息：" + req.Text.Content,
		},
	}
	body := &dto.AiRequest{
		Model:       enum.ChatglmTurbo,
		Prompt:      prompt,
		TopP:        0.9,
		Temperature: 0.9,
	}
	ans, err := logic.Send(body, enum.ChatglmTurbo, enum.Invoke)
	if err != nil {
		return
	}
	aiResp := new(dto.AiResponse)
	err = json.Unmarshal([]byte(ans), aiResp)
	if err != nil {
		return
	}
	if aiResp.Code != 200 || len(aiResp.Data.Choices) <= 0 {
		return
	}

	req.Text.Content = aiResp.Data.Choices[0].Content
	if strings.HasSuffix(req.Text.Content, "\"") {
		req.Text.Content = strings.TrimSuffix(req.Text.Content, "\"")
	}
	if strings.HasPrefix(req.Text.Content, "\"") {
		req.Text.Content = strings.TrimPrefix(req.Text.Content, "\"")
	}
	req.Text.Content = strings.TrimSpace(req.Text.Content)
	err = logic.DingSend(req)
	return
}
