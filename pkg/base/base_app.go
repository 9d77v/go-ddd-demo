package base

import (
	"context"
	"fmt"
	"log"
	"strings"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
)

const _ttl = 10

type BaseApp struct {
	AppName     string
	ServiceName string
	ServerHost  string
	ServerPort  uint64
	EtcdClient  *clientv3.Client
}

type BaseParam struct {
	AppName     string
	ServiceName string
	ServerHost  string
	ServerPort  uint64
	EtcdAddress string
}

func NewBaseApp(p *BaseParam) BaseApp {
	cli, err := clientv3.NewFromURL(p.EtcdAddress)
	if err != nil {
		panic(err)
	}
	return BaseApp{
		AppName:     p.AppName,
		ServiceName: p.ServiceName,
		ServerPort:  p.ServerPort,
		ServerHost:  p.ServerHost,
		EtcdClient:  cli,
	}
}

func (a *BaseApp) Register() {
	target := a.AppName + "/services/" + a.ServiceName
	em, _ := endpoints.NewManager(a.EtcdClient, target)
	addr := fmt.Sprintf("%s:%d", a.ServerHost, a.ServerPort)
	key := target + "/" + strings.ReplaceAll(addr, ".", "-")
	lease := clientv3.NewLease(a.EtcdClient)
	leaseResp, _ := lease.Grant(context.TODO(), _ttl)
	leaseRespChan, err := lease.KeepAlive(context.TODO(), leaseResp.ID)
	if err != nil {
		log.Panicf("续租失败:%s\n", err.Error())
	}
	err = em.AddEndpoint(context.TODO(), key, endpoints.Endpoint{Addr: addr}, clientv3.WithLease(leaseResp.ID))
	if err != nil {
		log.Panicln("etce add endpoint failed")
	}
	go func() {
		for {
			leaseKeepResp := <-leaseRespChan
			if leaseKeepResp == nil {
				fmt.Printf("已经关闭续租功能\n")
				return
			}
		}
	}()
}

func (a *BaseApp) Deregister() {
	target := a.AppName + "/services/" + a.ServiceName
	em, _ := endpoints.NewManager(a.EtcdClient, target)
	addr := fmt.Sprintf("%s:%d", a.ServerHost, a.ServerPort)
	key := target + "/" + strings.ReplaceAll(addr, ".", "-")
	em.DeleteEndpoint(context.TODO(), key)
}
