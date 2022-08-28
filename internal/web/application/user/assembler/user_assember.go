package assembler

import (
	"github.com/9d77v/go-ddd-demo/api/proto/user/pb"
	"github.com/9d77v/go-ddd-demo/internal/web/model"
	"github.com/9d77v/go-ddd-demo/pkg/datetime"
)

func NewUsersFromPBs(users []*pb.User) []*model.User {
	userPbs := make([]*model.User, len(users))
	for i, v := range users {
		userPbs[i] = NewUserFromPB(v)
	}
	return userPbs
}

func NewUserFromPB(user *pb.User) *model.User {
	if user == nil {
		return nil
	}
	return &model.User{
		ID:        user.Id,
		Phone:     user.Phone,
		Gender:    int(user.Gender),
		Nickname:  user.Nickname,
		CreatedAt: datetime.FormatTimeStamp(user.CreatedAt),
		UpdatedAt: datetime.FormatTimeStamp(user.UpdatedAt),
	}
}

func NewCreateUserPbFromUser(in model.NewUser) *pb.CreateUserRequest {
	return &pb.CreateUserRequest{
		Phone:    in.Phone,
		Password: in.Password,
		Nickname: in.Nickname,
		Gender:   int32(in.Gender),
	}
}
