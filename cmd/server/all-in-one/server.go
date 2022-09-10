package main

import (
	"flag"
	"fmt"

	"github.com/9d77v/go-ddd-demo/internal/web"
	"github.com/alibaba/ioc-golang"
	"github.com/alibaba/ioc-golang/config"
)

func main() {
	var configPath = flag.String("conf", "conf", "请输入配置文件地址")
	flag.Parse()
	fmt.Println(*configPath)
	if err := ioc.Load(
		config.WithSearchPath(*configPath)); err != nil {
		panic(err)
	}
	app, err := web.GetStandaloneAppSingleton()
	if err != nil {
		panic(err)
	}
	app.Run()
}
