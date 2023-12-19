package message

import (
	"context"
	"errors"
	logic "zhipuai_api/internal/logic/common"
	"zhipuai_api/internal/svc"
	"zhipuai_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DingMessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDingMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DingMessageLogic {
	return &DingMessageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DingMessageLogic) DingMessage(req *types.DingMessageSendRequest) (err error) {
	if req.Content == "" {
		return errors.New("content is empty")
	}
	err = logic.DingSend(req.Content)
	return
}
