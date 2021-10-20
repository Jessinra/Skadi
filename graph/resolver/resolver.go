//go:generate go run github.com/99designs/gqlgen

package resolver

import (
	file "gitlab.com/trivery-id/skadi/internal/file/services"
	product "gitlab.com/trivery-id/skadi/internal/product/services"
	user "gitlab.com/trivery-id/skadi/internal/user/services"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{}

var (
	fileService    *file.FileService
	productService *product.ProductService
	userService    *user.UserService
)

func InitResolvers() {
	fileService = file.GetFileService()
	productService = product.GetProductService()
	userService = user.GetUserService()
}
