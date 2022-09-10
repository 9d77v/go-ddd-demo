package user

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/9d77v/go-ddd-demo/api/proto/user/pb"
	"github.com/9d77v/go-ddd-demo/internal/user/application/service/command/impl"
	queryImpl "github.com/9d77v/go-ddd-demo/internal/user/application/service/query/impl"
	"github.com/9d77v/go-ddd-demo/pkg/base"

	"google.golang.org/grpc"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton
// +ioc:autowire:paramType=Param
// +ioc:autowire:constructFunc=Init
type App struct {
	base.BaseApp
	UserService      impl.UserServiceImplIOCInterface      `singleton:""`
	UserQueryService queryImpl.UserServiceImplIOCInterface `singleton:""`
}

type Param struct {
	base.BaseParam
}

func (p *Param) Init(a *App) (*App, error) {
	a.BaseApp = base.NewBaseApp(&p.BaseParam)
	return a, nil
}

func (a *App) Run() {
	errc := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		errc <- fmt.Errorf("%s", <-c)
	}()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", a.ServerPort))
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()
	pb.RegisterUserServiceServer(srv, a.UserService)
	pb.RegisterUserQueryServiceServer(srv, a.UserQueryService)
	a.Register()
	go func() {
		errc <- srv.Serve(lis)
		log.Printf("connect to http://localhost:%d/ for GraphQL playground", a.ServerPort)
	}()
	log.Printf("exiting (%v)", <-errc)
	srv.GracefulStop()
	log.Println("exited")
	a.Deregister()
}
