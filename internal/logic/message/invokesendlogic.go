package message

import (
	"context"
	"encoding/json"
	"errors"
	"zhipuai_api/internal/common/dto"
	"zhipuai_api/internal/common/enum"
	"zhipuai_api/internal/common/utils"
	logic "zhipuai_api/internal/logic/common"
	"zhipuai_api/internal/model"
	"zhipuai_api/pkg/resource"

	"zhipuai_api/internal/svc"
	"zhipuai_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InvokeSendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInvokeSendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InvokeSendLogic {
	return &InvokeSendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InvokeSendLogic) InvokeSend(req *types.InvokeParams) ([]*dto.Choices, error) {

	sessionMgr := model.ZhipuaiSessionMgr(resource.DB)
	session, err := sessionMgr.GetByOption(sessionMgr.WithUUID(req.UUID))
	if err != nil {
		return nil, err
	}
	if session.ID <= 0 {
		session = model.ZhipuaiSession{
			UUID:  req.UUID,
			Title: utils.GetTitleFromContent(req.Content),
		}
		err = sessionMgr.Insert(&session)
		if err != nil {
			return nil, err
		}
	}
	messageMgr := model.ZhipuaiMessageMgr(resource.DB)
	err = messageMgr.Insert(&model.ZhipuaiMessage{
		RoleName: enum.User,
		Content:  req.Content,
		UUID:     req.UUID,
	})
	if err != nil {
		return nil, err
	}
	messages, err := messageMgr.GetBySize(10)
	if err != nil {
		return nil, err
	}
	prompt := make([]*dto.Prompt, 0)
	for _, message := range messages {
		prompt = append(prompt, &dto.Prompt{
			RoleName: message.RoleName,
			Content:  message.Content,
		})
	}

	body := &dto.AiRequest{
		Model:       enum.ChatglmTurbo,
		Prompt:      prompt,
		TopP:        0.9,
		Temperature: 0.9,
	}

	resp, err := logic.Send(body, enum.ChatglmTurbo, enum.Invoke)
	if err != nil {
		return nil, err
	}
	go func() {
		logMgr := model.ZhipuaiSendLogMgr(resource.DB)
		reqJson, _ := json.Marshal(body)
		_ = logMgr.Insert(&model.ZhipuaiSendLog{
			Req:  string(reqJson),
			Resp: resp,
			UUID: req.UUID,
		})
	}()
	aiResp := new(dto.AiResponse)
	err = json.Unmarshal([]byte(resp), aiResp)
	if err != nil {
		return nil, err
	}
	if aiResp.Code != 200 {
		return nil, errors.New(aiResp.Msg)
	}

	return aiResp.Data.Choices, nil
}
