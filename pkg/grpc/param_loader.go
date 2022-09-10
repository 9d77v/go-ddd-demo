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

package grpc

import (
	"errors"
	"fmt"
	"log"
	"net/url"

	clientv3 "go.etcd.io/etcd/client/v3"
	resolver "go.etcd.io/etcd/client/v3/naming/resolver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/alibaba/ioc-golang/autowire"
	"github.com/alibaba/ioc-golang/config"
)

type Config struct {
	Address string
}

type paramLoader struct {
}

/*
	 Load support load grpc field:
	 ```go
	 ResourceServiceClient resources.ResourceServiceClient `grpc:"resource-service"`
	 ```go

	 from:

	 ```yaml
	 autowire:
	   grpc:
		 resource-service:
		   address: 127.0.0.1:8080
	 ```

	 Make Dial and generate *grpc.ClientConn as param
*/
func (p *paramLoader) Load(_ *autowire.StructDescriptor, fi *autowire.FieldInfo) (interface{}, error) {
	if fi == nil {
		return nil, errors.New("not supported")
	}
	grpcConfig := &Config{}
	if err := config.LoadConfigByPrefix(fmt.Sprintf("autowire%[1]sgrpc%[1]s%[2]s", config.YamlConfigSeparator, fi.TagValue), grpcConfig); err != nil {
		return nil, err
	}
	u, err := url.Parse(grpcConfig.Address)
	if err != nil {
		log.Panicln("地址解析出错", err)
	}
	cli, cerr := clientv3.NewFromURL(u.Host)
	if cerr != nil {
		log.Panicln("init etcd client failed", cerr)
	}
	etcdResolver, err := resolver.NewBuilder(cli)
	if err != nil {
		log.Panicln("init etcd resolver failed", cerr)
	}
	o, err := grpc.Dial(grpcConfig.Address,
		grpc.WithResolvers(etcdResolver),
		grpc.WithDefaultServiceConfig(`{"loadBalancingConfig": [ { "round_robin": {} } ]}`),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Panicln("dial failed", err)
	}
	return o, err
}
