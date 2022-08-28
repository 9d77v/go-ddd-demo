package main

import (
	"flag"

	"github.com/9d77v/go-ddd-demo/internal/web"
	"github.com/9d77v/go-ddd-demo/pkg/base"
	"github.com/9d77v/go-pkg/env"
	"github.com/alibaba/ioc-golang"
	"github.com/alibaba/ioc-golang/config"
)

var namespaceId = env.String("NamespaceId", "pdc-dev")
var nacosAddr = env.String("NacosAddr", "127.0.0.1")
var nacosPort = env.Int("NacosPort", 8848)
var serviceName = env.String("ServiceName", "web-service")
var serverPort = env.Int("ServerPort", 7100)

func main() {
	var configPath = flag.String("conf", "conf", "请输入配置文件地址")
	flag.Parse()
	if err := ioc.Load(
		config.WithSearchPath(*configPath)); err != nil {
		panic(err)
	}
	app, err := web.GetAppSingleton(&web.Param{
		BaseParam: base.BaseParam{
			NamespaceId: namespaceId,
			NacosAddr:   nacosAddr,
			NacosPort:   nacosPort,
			ServiceName: serviceName,
			ServerPort:  uint64(serverPort),
		},
	})
	if err != nil {
		panic(err)
	}
	app.Run()
}
