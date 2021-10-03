package services

import (
	"context"
	"strings"

	"gitlab.com/trivery-id/skadi/internal/product/domain"
	"gitlab.com/trivery-id/skadi/internal/product/repositories"
)

func (svc *ProductService) CreateNewProduct(ctx context.Context, in CreateNewProductInput) (*domain.Product, error) {
	in.FillDefault()
	if err := in.Validate(); err != nil {
		return nil, err
	}

	product := &domain.Product{
		Name:        in.Name,
		Description: in.Description,
		ImagesURLs:  strings.Join(in.ImagesURLs, ","),
		Weight:      *in.Weight,
		Dimensions:  *in.Dimensions,
		Categories:  strings.Join(in.Categories, ","),
		Locations:   []domain.ProductLocation{},
		Prices:      []domain.ProductPrice{},
	}
	if err := svc.ProductRepository.Add(ctx, product); err != nil {
		return nil, err
	}

	in.Location.ProductID = product.ID
	location, err := svc.CreateNewProductLocation(ctx, in.Location)
	if err != nil {
		return nil, err
	}

	in.Price.ProductID = product.ID
	price, err := svc.CreateNewProductPrice(ctx, in.Price)
	if err != nil {
		return nil, err
	}

	product.Locations = append(product.Locations, *location)
	product.Prices = append(product.Prices, *price)

	return product, nil
}

func (svc *ProductService) GetProduct(ctx context.Context, productID uint64) (*domain.Product, error) {
	return svc.ProductRepository.Find(ctx, productID)
}

func (svc *ProductService) GetAllProducts(ctx context.Context, in GetAllProductsInput) ([]domain.Product, error) {
	return svc.ProductRepository.FindAll(ctx, repositories.FindAllInput{
		Limit:  in.Limit,
		Offset: in.Offset,
	})
}
