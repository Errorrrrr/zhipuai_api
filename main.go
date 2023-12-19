package main

import (
	"flag"
	"fmt"
	"net/http"
	"zhipuai_api/pkg/resource"

	"zhipuai_api/internal/config"
	"zhipuai_api/internal/handler"
	"zhipuai_api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/zhipuaiapi-dev.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	if err := resource.Init(c); err != nil {
		panic(err)
	}
	defer resource.Destroy()
	config.Conf = &c
	opt := rest.WithNotFoundHandler(NotFoundHandler())
	server := rest.MustNewServer(c.RestConf, opt)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

func NotFoundHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		http.Error(writer, `the incorrect Route .`, http.StatusOK)
		return
	}
}
