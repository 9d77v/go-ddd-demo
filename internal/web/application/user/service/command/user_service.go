package command

import (
	"context"

	"github.com/9d77v/go-ddd-demo/internal/web/model"
)

type UserService interface {
	CreateUser(ctx context.Context, input model.NewUser) (*model.User, error)
}
