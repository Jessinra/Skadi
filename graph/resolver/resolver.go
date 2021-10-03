//go:generate go run github.com/99designs/gqlgen

package resolver

import (
	product "gitlab.com/trivery-id/skadi/internal/product/services"
	todo "gitlab.com/trivery-id/skadi/internal/todo/services"
	user "gitlab.com/trivery-id/skadi/internal/user/services"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{}

var (
	productService *product.ProductService
	userService    *user.UserService
	todoService    *todo.TodoService
)

func InitResolvers() {
	productService = product.GetProductService()
	userService = user.GetUserService()
	todoService = &todo.TodoService{}
}
