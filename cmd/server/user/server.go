/*
 * Copyright (c) 2022, Alibaba Group;
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"flag"

	"github.com/alibaba/ioc-golang/config"

	"github.com/9d77v/go-ddd-demo/internal/user"
	"github.com/9d77v/go-ddd-demo/pkg/base"
	"github.com/9d77v/go-ddd-demo/pkg/util"
	"github.com/9d77v/go-pkg/env"
	"github.com/alibaba/ioc-golang"
)

const (
	appName     = "go-ddd-demo"
	serviceName = "user-service"
)

var etcdAddress = env.String("ETCD_ADDRESS", "http://localhost:2379")

func main() {
	var configPath = flag.String("conf", "conf", "请输入配置文件地址")
	flag.Parse()
	if err := ioc.Load(
		config.WithSearchPath(*configPath)); err != nil {
		panic(err)
	}
	app, err := user.GetAppSingleton(&user.Param{
		BaseParam: base.BaseParam{
			AppName:     appName,
			ServiceName: serviceName,
			ServerHost:  util.GetNetworkIp(),
			ServerPort:  util.GetRandomPort(),
			EtcdAddress: etcdAddress,
		},
	})
	if err != nil {
		panic(err)
	}
	app.Run()
}
