//go:generate mockgen --build_flags=--mod=mod -destination=mocks/IProductRepository.go -package=mocks . IProductRepository
//go:generate mockgen --build_flags=--mod=mod -destination=mocks/IProductLocationRepository.go -package=mocks . IProductLocationRepository
//go:generate mockgen --build_flags=--mod=mod -destination=mocks/IProductPriceRepository.go -package=mocks . IProductPriceRepository

package services

import (
	"context"

	skadipsql "gitlab.com/trivery-id/skadi/datasources/postgres/skadi"
	"gitlab.com/trivery-id/skadi/internal/product/domain"
	"gitlab.com/trivery-id/skadi/internal/product/repositories"
	"gitlab.com/trivery-id/skadi/utils/errors"
)

type ProductService struct {
	ProductRepository         IProductRepository
	ProductPriceRepository    IProductPriceRepository
	ProductLocationRepository IProductLocationRepository
}

type IProductRepository interface {
	Add(ctx context.Context, product *domain.Product) error
	Find(ctx context.Context, productID uint64) (*domain.Product, error)
	FindAll(ctx context.Context, in repositories.FindAllInput) ([]domain.Product, error)
}

type IProductPriceRepository interface {
	Add(ctx context.Context, price *domain.ProductPrice) error
	Update(ctx context.Context, price *domain.ProductPrice) error
	Find(ctx context.Context, priceID uint64) (*domain.ProductPrice, error)
	Delete(ctx context.Context, priceID uint64) error
}

type IProductLocationRepository interface {
	Add(ctx context.Context, location *domain.ProductLocation) error
	Update(ctx context.Context, location *domain.ProductLocation) error
	Find(ctx context.Context, locationID uint64) (*domain.ProductLocation, error)
	Delete(ctx context.Context, locationID uint64) error
}

func NewProductService() (*ProductService, error) {
	return &ProductService{
		ProductRepository:         repositories.NewProductRepository(skadipsql.DB),
		ProductPriceRepository:    repositories.NewProductPriceRepository(skadipsql.DB),
		ProductLocationRepository: repositories.NewProductLocationRepository(skadipsql.DB),
	}, nil
}

func (ProductService) InitDependencies() error {
	return nil
}

func (svc *ProductService) Validate() error {
	if svc.ProductRepository == nil {
		return errors.NewUnprocessableEntityError("invalid user service, haven't set ProductRepository")
	}
	if svc.ProductPriceRepository == nil {
		return errors.NewUnprocessableEntityError("invalid user service, haven't set ProductPriceRepository")
	}
	if svc.ProductLocationRepository == nil {
		return errors.NewUnprocessableEntityError("invalid user service, haven't set ProductLocationRepository")
	}

	return nil
}
