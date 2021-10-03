package services

import (
	"context"

	"gitlab.com/trivery-id/skadi/internal/product/domain"
	"gitlab.com/trivery-id/skadi/utils/metadata"
)

func (svc *ProductService) CreateNewProductPrice(ctx context.Context, in CreateNewProductPriceInput) (*domain.ProductPrice, error) {
	user := metadata.GetUserFromContext(ctx)

	if err := in.Validate(); err != nil {
		return nil, err
	}

	product, err := svc.ProductRepository.Find(ctx, in.ProductID)
	if err != nil {
		return nil, err
	}

	price := &domain.ProductPrice{
		UserID:           user.ID,
		ProductID:        product.ID,
		Currency:         in.Currency,
		Price:            in.Price,
		IsPriceEstimated: in.IsPriceEstimated,
	}
	if err := svc.ProductPriceRepository.Add(ctx, price); err != nil {
		return nil, err
	}

	return price, nil
}

func (svc *ProductService) UpdateProductPrice(ctx context.Context, in UpdateProductPriceInput) (*domain.ProductPrice, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}

	price, err := svc.ProductPriceRepository.Find(ctx, in.ID)
	if err != nil {
		return nil, err
	}

	if err := svc.CanUpdateProductPrice(ctx, price); err != nil {
		return nil, err
	}

	if in.Price != nil {
		price.Price = *in.Price
	}
	if in.IsPriceEstimated != nil {
		price.IsPriceEstimated = *in.IsPriceEstimated
	}
	if err := svc.ProductPriceRepository.Update(ctx, price); err != nil {
		return nil, err
	}

	return price, nil
}

func (svc *ProductService) DeleteProductPrice(ctx context.Context, priceID uint64) error {
	if err := svc.CanDeleteProductPriceByID(ctx, priceID); err != nil {
		return err
	}

	return svc.ProductPriceRepository.Delete(ctx, priceID)
}
