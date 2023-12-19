package message

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zhipuai_api/internal/logic/message"
	"zhipuai_api/internal/svc"
	"zhipuai_api/internal/types"
)

func DingMessageSendHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DingMessageSendRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := message.NewDingMessageSendLogic(r.Context(), svcCtx)
		err := l.DingMessageSend(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, nil)
		}
	}
}
