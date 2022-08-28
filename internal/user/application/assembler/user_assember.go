package assembler

import (
	"github.com/9d77v/go-ddd-demo/api/proto/user/pb"
	"github.com/9d77v/go-ddd-demo/internal/user/domain/entity"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func NewUserPBsFromEntitys(users []*entity.User) []*pb.User {
	userPbs := make([]*pb.User, len(users))
	for i, v := range users {
		userPbs[i] = NewUserPBFromEntity(v)
	}
	return userPbs
}

func NewUserPBFromEntity(user *entity.User) *pb.User {
	if user == nil {
		return nil
	}
	return &pb.User{
		Id:        user.ID,
		Phone:     user.Phone,
		Gender:    int32(user.Gender),
		Nickname:  user.Nickname,
		CreatedAt: timestamppb.New(*user.CreatedAt),
		UpdatedAt: timestamppb.New(*user.UpdatedAt),
	}
}
