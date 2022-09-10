package impl

import (
	"context"
	"fmt"

	"github.com/9d77v/go-ddd-demo/api/proto/user/pb"
	"github.com/9d77v/go-ddd-demo/internal/web/application/user/assembler"
	"github.com/9d77v/go-ddd-demo/internal/web/application/user/service/query"
	"github.com/9d77v/go-ddd-demo/internal/web/model"
	"github.com/9d77v/go-ddd-demo/pkg/grpc"
	autowire "github.com/alibaba/ioc-golang/autowire"
	googleGRPC "google.golang.org/grpc"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton
type UserServiceImpl struct {
	HelloServiceClient pb.UserQueryServiceClient `grpc:"user-service"`
}

// UserInfo implements query.UserService
func (s *UserServiceImpl) UserInfo(ctx context.Context, id string) (*model.User, error) {
	user, err := s.HelloServiceClient.UserInfo(ctx, &pb.UserInfoRequest{Id: id})
	if user == nil || err != nil {
		return &model.User{}, err
	}
	return assembler.NewUserFromPB(user.User), nil
}

// Users implements query.UserService
func (s *UserServiceImpl) UserPage(ctx context.Context, page, size int) (*model.UserConnection, error) {
	user, err := s.HelloServiceClient.UserPage(ctx, &pb.UserPageRequest{Page: int32(page), Size: int32(size)})
	if err != nil {
		return nil, err
	}
	userConn := new(model.UserConnection)
	userConn.TotalCount = int(user.TotalCount)
	userConn.Edges = assembler.NewUsersFromPBs(user.Edges)
	return userConn, nil
}

var _ query.UserService = &UserServiceImpl{}

func init() {
	// register grpc client
	grpc.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return new(pb.UserQueryServiceClient)
		},
		ParamFactory: func() interface{} {
			return &googleGRPC.ClientConn{}
		},
		ConstructFunc: func(impl interface{}, param interface{}) (interface{}, error) {
			conn := param.(*googleGRPC.ClientConn)
			fmt.Println("create conn target ", conn.Target())
			return pb.NewUserQueryServiceClient(conn), nil
		},
	})
}
