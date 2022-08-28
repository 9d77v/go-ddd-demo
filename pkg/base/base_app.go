package base

import (
	"github.com/9d77v/go-ddd-demo/pkg/util"
	config "github.com/alibaba/ioc-golang/extension/config_center/nacos"
	"github.com/alibaba/ioc-golang/extension/registry/nacos"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

type BaseApp struct {
	NacosClient  nacos.NamingClientIOCInterface
	ConfigClient config.ConfigClientIOCInterface
	ServiceName  string
	ServerPort   uint64
}

type BaseParam struct {
	NamespaceId string
	NacosAddr   string
	NacosPort   int
	ServiceName string
	ServerPort  uint64
}

func NewBaseApp(p *BaseParam) BaseApp {
	nacosParam := vo.NacosClientParam{
		ServerConfigs: []constant.ServerConfig{
			{
				IpAddr: p.NacosAddr,
				Port:   uint64(p.NacosPort),
			},
		},
		ClientConfig: &constant.ClientConfig{
			NamespaceId: p.NamespaceId,
		},
	}
	nacosClient, err := nacos.GetNamingClientIOCInterface(&nacos.Param{NacosClientParam: nacosParam})
	if err != nil {
		panic(err)
	}
	configClient, err := config.GetConfigClientIOCInterface(&config.Param{NacosClientParam: nacosParam})
	if err != nil {
		panic(err)
	}
	return BaseApp{
		NacosClient:  nacosClient,
		ConfigClient: configClient,
		ServiceName:  p.ServiceName,
		ServerPort:   p.ServerPort,
	}
}

func (a *BaseApp) Register() {
	_, err := a.NacosClient.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          util.GetNetworkIp(),
		Port:        a.ServerPort,
		ServiceName: a.ServiceName,
	})
	if err != nil {
		panic(err)
	}
}

func (a *BaseApp) Deregister() {
	_, err := a.NacosClient.DeregisterInstance(vo.DeregisterInstanceParam{
		Ip:          util.GetNetworkIp(),
		Port:        a.ServerPort,
		ServiceName: a.ServiceName,
	})
	if err != nil {
		panic(err)
	}
}
