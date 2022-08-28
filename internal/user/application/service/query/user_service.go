package query

import (
	"context"

	"github.com/9d77v/go-ddd-demo/api/proto/user/pb"
)

type UserService interface {
	UserPage(ctx context.Context, in *pb.UserPageRequest) (*pb.UserPageResponse, error)
	UserInfo(ctx context.Context, in *pb.UserInfoRequest) (*pb.UserInfoResponse, error)
}
