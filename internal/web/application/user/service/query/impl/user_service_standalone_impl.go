package impl

import (
	"context"

	"github.com/9d77v/go-ddd-demo/api/proto/user/pb"
	"github.com/9d77v/go-ddd-demo/internal/user/application/service/query/impl"
	"github.com/9d77v/go-ddd-demo/internal/web/application/user/assembler"
	"github.com/9d77v/go-ddd-demo/internal/web/application/user/service/query"
	"github.com/9d77v/go-ddd-demo/internal/web/model"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton
type UserServiceStandaloneImpl struct {
	UserService impl.UserServiceImplIOCInterface `singleton:""`
}

// UserInfo implements query.UserService
func (s *UserServiceStandaloneImpl) UserInfo(ctx context.Context, id string) (*model.User, error) {
	user, err := s.UserService.UserInfo(ctx, &pb.UserInfoRequest{Id: id})
	if user == nil || err != nil {
		return &model.User{}, err
	}
	return assembler.NewUserFromPB(user.User), nil
}

// Users implements query.UserService
func (s *UserServiceStandaloneImpl) UserPage(ctx context.Context, page, size int) (*model.UserConnection, error) {
	user, err := s.UserService.UserPage(ctx, &pb.UserPageRequest{Page: int32(page), Size: int32(size)})
	if err != nil {
		return nil, err
	}
	userConn := new(model.UserConnection)
	userConn.TotalCount = int(user.TotalCount)
	userConn.Edges = assembler.NewUsersFromPBs(user.Edges)
	return userConn, nil
}

var _ query.UserService = &UserServiceStandaloneImpl{}
