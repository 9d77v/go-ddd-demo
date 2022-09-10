package main

import (
	"flag"

	"github.com/9d77v/go-ddd-demo/internal/web"
	"github.com/9d77v/go-ddd-demo/pkg/base"
	"github.com/9d77v/go-pkg/env"
	"github.com/alibaba/ioc-golang"
	"github.com/alibaba/ioc-golang/config"
)

var serviceName = env.String("ServiceName", "web-service")
var serverPort = env.Int("ServerPort", 7100)
var etcdAddress = env.String("ETCD_ADDRESS", "http://localhost:2379")

func main() {
	var configPath = flag.String("conf", "conf", "请输入配置文件地址")
	flag.Parse()
	if err := ioc.Load(
		config.WithSearchPath(*configPath)); err != nil {
		panic(err)
	}
	app, err := web.GetAppSingleton(&web.Param{
		BaseParam: base.BaseParam{
			ServiceName: serviceName,
			ServerPort:  uint64(serverPort),
			EtcdAddress: etcdAddress,
		},
	})
	if err != nil {
		panic(err)
	}
	app.Run()
}
