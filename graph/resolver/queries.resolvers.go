package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"gitlab.com/trivery-id/skadi/graph/generated"
	"gitlab.com/trivery-id/skadi/graph/model"
	productSvc "gitlab.com/trivery-id/skadi/internal/product/services"
	errs "gitlab.com/trivery-id/skadi/utils/errors"
	"gitlab.com/trivery-id/skadi/utils/metadata"
)

func (r *queryResolver) User(ctx context.Context, id uint64) (*model.User, error) {
	if !metadata.IsAuthenticated(ctx) {
		return nil, errs.ErrInvalidCredentials
	}

	user, err := userService.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}

	return model.NewUser(user), nil
}

func (r *queryResolver) Product(ctx context.Context, id uint64) (*model.Product, error) {
	if !metadata.IsAuthenticated(ctx) {
		return nil, errs.ErrInvalidCredentials
	}

	product, err := productService.GetProduct(ctx, id)
	if err != nil {
		return nil, err
	}

	return model.NewProduct(product), nil
}

func (r *queryResolver) Products(ctx context.Context, limit *int, offset *int) ([]model.Product, error) {
	if !metadata.IsAuthenticated(ctx) {
		return nil, errs.ErrInvalidCredentials
	}

	products, err := productService.GetAllProducts(ctx, productSvc.GetAllProductsInput{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, err
	}

	return model.NewProducts(products), nil
}

func (r *queryResolver) Order(ctx context.Context, id uint64) (*model.Order, error) {
	if !metadata.IsAuthenticated(ctx) {
		return nil, errs.ErrInvalidCredentials
	}

	order, err := productService.GetOrder(ctx, productSvc.GetOrderInput{
		OrderID: id,
	})
	if err != nil {
		return nil, err
	}

	return model.NewOrder(order), nil
}

func (r *queryResolver) Orders(ctx context.Context, userID uint64, limit *int, offset *int) ([]model.Order, error) {
	if !metadata.IsAuthenticated(ctx) {
		return nil, errs.ErrInvalidCredentials
	}

	orders, err := productService.GetAllOrders(ctx, productSvc.GetAllOrdersInput{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, err
	}

	return model.NewOrders(orders), nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
