package impl

import (
	"context"
	"fmt"

	"github.com/9d77v/go-ddd-demo/api/proto/user/pb"
	"github.com/9d77v/go-ddd-demo/internal/web/application/user/assembler"
	"github.com/9d77v/go-ddd-demo/internal/web/application/user/service/command"
	"github.com/9d77v/go-ddd-demo/internal/web/model"
	"github.com/9d77v/go-ddd-demo/pkg/grpc"
	autowire "github.com/alibaba/ioc-golang/autowire"
	googleGRPC "google.golang.org/grpc"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton
type UserServiceImpl struct {
	HelloServiceClient pb.UserServiceClient `grpc:"user-service"`
}

// CreateUser implements service.UserApp
func (s *UserServiceImpl) CreateUser(ctx context.Context, in model.NewUser) (*model.User, error) {
	user, err := s.HelloServiceClient.CreateUser(ctx, assembler.NewCreateUserPbFromUser(in))
	return &model.User{ID: user.Id}, err
}

var _ command.UserService = &UserServiceImpl{}

func init() {
	// register grpc client
	grpc.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return new(pb.UserServiceClient)
		},
		ParamFactory: func() interface{} {
			return &googleGRPC.ClientConn{}
		},
		ConstructFunc: func(impl interface{}, param interface{}) (interface{}, error) {
			conn := param.(*googleGRPC.ClientConn)
			fmt.Println("create conn target ", conn.Target())
			return pb.NewUserServiceClient(conn), nil
		},
	})
}
