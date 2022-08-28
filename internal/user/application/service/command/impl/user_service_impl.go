package impl

import (
	"context"

	"github.com/9d77v/go-ddd-demo/api/proto/user/pb"
	"github.com/9d77v/go-ddd-demo/internal/user/application/service/command"
	"github.com/9d77v/go-ddd-demo/internal/user/domain/entity"
	"github.com/9d77v/go-ddd-demo/internal/user/domain/enum"
	"github.com/9d77v/go-ddd-demo/internal/user/infrastructure/repository"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton
type UserServiceImpl struct {
	UserRepository repository.UserRepositoryIOCInterface `singleton:""`
}

// CreateUser implements command.UserService
func (s *UserServiceImpl) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user := entity.NewUser(in.Phone, in.Password, in.Nickname, enum.GenderEnum(in.Gender))
	err := s.UserRepository.GetDB().Save(user).Error
	return &pb.CreateUserResponse{Id: user.ID}, err
}

var _ command.UserService = &UserServiceImpl{}
