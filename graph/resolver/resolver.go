//go:generate go run github.com/99designs/gqlgen

package resolver

import (
	todo "gitlab.com/trivery-id/skadi/internal/todo/services"
	user "gitlab.com/trivery-id/skadi/internal/user/services"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{}

var (
	userService *user.UserService
	todoService *todo.TodoService
)

func InitResolvers() {
	userService = user.GetUserService()
	todoService = &todo.TodoService{}
}
