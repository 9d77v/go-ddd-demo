package resolver

import (
	"github.com/9d77v/go-ddd-demo/internal/web/generated"
	user "github.com/9d77v/go-ddd-demo/internal/web/application/user/service/command/impl"
	userQuery "github.com/9d77v/go-ddd-demo/internal/web/application/user/service/query/impl"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

//
//go:generate go run github.com/99designs/gqlgen generate

// +ioc:autowire=true
// +ioc:autowire:type=singleton
type Resolver struct {
	UserService      user.UserServiceImplIOCInterface      `singleton:""`
	UserQueryService userQuery.UserServiceImplIOCInterface `singleton:""`
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct {
	*Resolver
}
type queryResolver struct {
	*Resolver
}
