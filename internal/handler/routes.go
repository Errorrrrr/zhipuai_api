package handler

import (
	"net/http"

	message "zhipuai_api/internal/handler/message"
	"zhipuai_api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes([]rest.Route{
		{
			Method: http.MethodGet,
			Path:   "/",
			Handler: func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("hello world"))
			},
		},
	})
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/invoke/send",
				Handler: message.InvokeSendHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/dingtalk/message/send",
				Handler: message.DingMessageSendHandler(serverCtx),
			},
		},
	)
}
