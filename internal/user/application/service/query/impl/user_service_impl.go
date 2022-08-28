package impl

import (
	"context"
	"errors"

	"github.com/9d77v/go-ddd-demo/api/proto/user/pb"
	"github.com/9d77v/go-ddd-demo/internal/user/application/assembler"
	"github.com/9d77v/go-ddd-demo/internal/user/application/service/query"
	"github.com/9d77v/go-ddd-demo/internal/user/domain/entity"
	"github.com/9d77v/go-ddd-demo/internal/user/infrastructure/repository"
	"gorm.io/gorm"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton
type UserServiceImpl struct {
	UserRepository repository.UserRepositoryIOCInterface `singleton:""`
}

// UserInfo implements query.UserService
func (s *UserServiceImpl) UserInfo(ctx context.Context, in *pb.UserInfoRequest) (*pb.UserInfoResponse, error) {
	user := &entity.User{}
	err := s.UserRepository.GetDB().Where("id=?", in.Id).First(user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &pb.UserInfoResponse{User: assembler.NewUserPBFromEntity(user)}, err
}

// UserPage implements query.UserService
func (s *UserServiceImpl) UserPage(ctx context.Context, in *pb.UserPageRequest) (*pb.UserPageResponse, error) {
	var totalCount int64 = 0
	users := []*entity.User{}
	err := s.UserRepository.GetDB().Model(&entity.User{}).Count(&totalCount).Error
	if err != nil || totalCount == 0 {
		return &pb.UserPageResponse{}, nil
	}
	err = s.UserRepository.GetDB().Offset(int(in.Page-1) * int(in.Size)).Limit(int(in.Size)).Find(&users).Error
	return &pb.UserPageResponse{TotalCount: totalCount, Edges: assembler.NewUserPBsFromEntitys(users)}, err
}

var _ query.UserService = &UserServiceImpl{}
