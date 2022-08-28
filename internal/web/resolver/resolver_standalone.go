package resolver

import (
	user "github.com/9d77v/go-ddd-demo/internal/web/application/user/service/command/impl"
	userQuery "github.com/9d77v/go-ddd-demo/internal/web/application/user/service/query/impl"
	"github.com/9d77v/go-ddd-demo/internal/web/generated"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

//
//go:generate go run github.com/99designs/gqlgen generate

// +ioc:autowire=true
// +ioc:autowire:type=singleton
type StandaloneResolver struct {
	UserService      user.UserServiceStandaloneImplIOCInterface      `singleton:""`
	UserQueryService userQuery.UserServiceStandaloneImplIOCInterface `singleton:""`
}

// Mutation returns generated.MutationResolver implementation.
func (r *StandaloneResolver) Mutation() generated.MutationResolver {
	return &standaloneMutationResolver{r}
}

// Query returns generated.QueryResolver implementation.
func (r *StandaloneResolver) Query() generated.QueryResolver {
	return &standaloneQueryResolver{r}
}

type standaloneMutationResolver struct {
	*StandaloneResolver
}
type standaloneQueryResolver struct {
	*StandaloneResolver
}
