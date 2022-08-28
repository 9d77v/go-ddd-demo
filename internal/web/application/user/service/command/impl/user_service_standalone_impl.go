package impl

import (
	"context"

	"github.com/9d77v/go-ddd-demo/internal/user/application/service/command/impl"
	"github.com/9d77v/go-ddd-demo/internal/web/application/user/assembler"
	"github.com/9d77v/go-ddd-demo/internal/web/application/user/service/command"
	"github.com/9d77v/go-ddd-demo/internal/web/model"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton
type UserServiceStandaloneImpl struct {
	UserService impl.UserServiceImplIOCInterface `singleton:""`
}

// CreateUser implements service.UserApp
func (s *UserServiceStandaloneImpl) CreateUser(ctx context.Context, in model.NewUser) (*model.User, error) {
	user, err := s.UserService.CreateUser(ctx, assembler.NewCreateUserPbFromUser(in))
	return &model.User{ID: user.Id}, err
}

var _ command.UserService = &UserServiceStandaloneImpl{}
