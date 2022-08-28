package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/9d77v/go-ddd-demo/internal/web/model"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	return r.UserService.CreateUser(ctx, input)
}

// UserPage is the resolver for the userPage field.
func (r *queryResolver) UserPage(ctx context.Context, page int, size int) (*model.UserConnection, error) {
	return r.UserQueryService.UserPage(ctx, page, size)
}

// UserInfo is the resolver for the userInfo field.
func (r *queryResolver) UserInfo(ctx context.Context, id string) (*model.User, error) {
	return r.UserQueryService.UserInfo(ctx, id)
}
