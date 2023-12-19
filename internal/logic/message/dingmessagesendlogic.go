package message

import (
	"context"
	"errors"
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
	err = logic.DingSend(req)
	return
}
