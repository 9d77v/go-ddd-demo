package query

import (
	"context"

	"github.com/9d77v/go-ddd-demo/internal/web/model"
)

type UserService interface {
	UserPage(ctx context.Context, page, size int) (*model.UserConnection, error)
	UserInfo(ctx context.Context, id string) (*model.User, error)
}
