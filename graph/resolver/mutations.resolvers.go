package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"gitlab.com/trivery-id/skadi/graph/generated"
	"gitlab.com/trivery-id/skadi/graph/model"
	productSvc "gitlab.com/trivery-id/skadi/internal/product/services"
	todoSvc "gitlab.com/trivery-id/skadi/internal/todo/services"
	userSvc "gitlab.com/trivery-id/skadi/internal/user/services"
)

func (r *mutationResolver) RegisterUser(ctx context.Context, input model.RegisterUser) (*model.User, error) {
	user, err := userService.RegisterUser(ctx, userSvc.RegisterUserInput{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	})
	if err != nil {
		return nil, err
	}

	return model.NewUser(user), nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUser) (*model.User, error) {
	user, err := userService.UpdateUser(ctx, userSvc.UpdateUserInput{
		ID:                input.ID,
		Name:              input.Name,
		PhoneNumber:       input.PhoneNumber,
		ProfilePictureURL: input.ProfilePictureURL,
		CurrencyMain:      input.CurrencyMain,
		CurrencySub:       input.CurrencySub,
	})
	if err != nil {
		return nil, err
	}

	return model.NewUser(user), nil
}

func (r *mutationResolver) UpdateUserPassword(ctx context.Context, input model.UpdateUserPassword) (bool, error) {
	if err := userService.UpdateUserPassword(ctx, userSvc.UpdateUserPasswordInput{
		ID:          input.ID,
		Password:    input.Password,
		NewPassword: input.NewPassword,
	}); err != nil {
		return false, err
	}

	return true, nil
}

func (r *mutationResolver) CreateProduct(ctx context.Context, input model.CreateProduct) (*model.Product, error) {
	product, err := productService.CreateNewProduct(ctx, productSvc.CreateNewProductInput{
		Name:        input.Name,
		Description: input.Description,
		ImagesURLs:  input.ImagesURLs,
		Weight:      input.Weight,
		Dimensions:  input.Dimensions,
		Categories:  input.Categories,
		Location: productSvc.CreateNewProductLocationInput{
			Text:      input.Location.Text,
			Country:   input.Location.Country,
			Province:  input.Location.Province,
			City:      input.Location.City,
			Area:      input.Location.Area,
			Street:    input.Location.Street,
			Building:  input.Location.Building,
			Store:     input.Location.Store,
			Longitude: input.Location.Longitude,
			Latitude:  input.Location.Latitude,
		},
		Price: productSvc.CreateNewProductPriceInput{
			Currency:         input.Price.Currency,
			Price:            input.Price.Price,
			IsPriceEstimated: input.Price.IsPriceEstimated,
		},
	})
	if err != nil {
		return nil, err
	}

	return model.NewProduct(product), nil
}

func (r *mutationResolver) CreateProductLocation(ctx context.Context, input model.CreateProductLocation) (*model.ProductLocation, error) {
	location, err := productService.CreateNewProductLocation(ctx, productSvc.CreateNewProductLocationInput{
		ProductID: *input.ProductID,
		Text:      input.Text,
		Country:   input.Country,
		Province:  input.Province,
		City:      input.City,
		Area:      input.Area,
		Street:    input.Street,
		Building:  input.Building,
		Store:     input.Store,
		Longitude: input.Longitude,
		Latitude:  input.Latitude,
	})
	if err != nil {
		return nil, err
	}

	return model.NewProductLocation(location), nil
}

func (r *mutationResolver) UpdateProductLocation(ctx context.Context, input model.UpdateProductLocation) (*model.ProductLocation, error) {
	location, err := productService.UpdateProductLocation(ctx, productSvc.UpdateProductLocationInput{
		ID:        input.ID,
		Text:      input.Text,
		Province:  input.Province,
		City:      input.City,
		Area:      input.Area,
		Street:    input.Street,
		Building:  input.Building,
		Store:     input.Store,
		Longitude: input.Longitude,
		Latitude:  input.Latitude,
	})
	if err != nil {
		return nil, err
	}

	return model.NewProductLocation(location), nil
}

func (r *mutationResolver) DeleteProductLocation(ctx context.Context, input model.DeleteProductLocation) (bool, error) {
	if err := productService.DeleteProductLocation(ctx, input.ID); err != nil {
		return false, err
	}

	return true, nil
}

func (r *mutationResolver) CreateProductPrice(ctx context.Context, input model.CreateProductPrice) (*model.ProductPrice, error) {
	price, err := productService.CreateNewProductPrice(ctx, productSvc.CreateNewProductPriceInput{
		ProductID:        *input.ProductID,
		Currency:         input.Currency,
		Price:            input.Price,
		IsPriceEstimated: input.IsPriceEstimated,
	})
	if err != nil {
		return nil, err
	}

	return model.NewProductPrice(price), nil
}

func (r *mutationResolver) UpdateProductPrice(ctx context.Context, input model.UpdateProductPrice) (*model.ProductPrice, error) {
	price, err := productService.UpdateProductPrice(ctx, productSvc.UpdateProductPriceInput{
		ID:               input.ID,
		Price:            input.Price,
		IsPriceEstimated: input.IsPriceEstimated,
	})
	if err != nil {
		return nil, err
	}

	return model.NewProductPrice(price), nil
}

func (r *mutationResolver) DeleteProductPrice(ctx context.Context, input model.DeleteProductPrice) (bool, error) {
	if err := productService.DeleteProductPrice(ctx, input.ID); err != nil {
		return false, err
	}

	return true, nil
}

func (r *mutationResolver) CreateOrder(ctx context.Context, input model.CreateOrder) (*model.Order, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) TakeOrder(ctx context.Context, input model.TakeOrder) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DropOrder(ctx context.Context, input model.DropOrder) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteOrder(ctx context.Context, input model.DeleteOrder) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.CreateTodo) (*model.Todo, error) {
	todo, err := todoService.CreateNewTodo(ctx, todoSvc.CreateNewTodoInput{
		Text:        input.Text,
		Description: input.Description,
		UserID:      input.UserID,
	})
	if err != nil {
		return nil, err
	}

	return model.NewTodo(todo), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
