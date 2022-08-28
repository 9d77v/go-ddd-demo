package command

import (
	"context"

	"github.com/9d77v/go-ddd-demo/api/proto/user/pb"
)

type UserService interface {
	CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserResponse, error)
}
