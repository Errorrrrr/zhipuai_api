package message

import (
	"context"
	"fmt"

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

func (l *DingMessageLogic) DingMessage(req *types.DingMessageRequest) (resp string, err error) {
	fmt.Println(req)
	return
}
