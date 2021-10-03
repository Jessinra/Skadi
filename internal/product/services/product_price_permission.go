package services

import (
	"context"

	"gitlab.com/trivery-id/skadi/internal/product/domain"
	"gitlab.com/trivery-id/skadi/utils/errors"
	"gitlab.com/trivery-id/skadi/utils/metadata"
)

func (svc *ProductService) CanUpdateProductPriceByID(ctx context.Context, priceID uint64) error {
	return svc.basicProductPriceAuthorPermission(ctx, priceID, nil)
}

func (svc *ProductService) CanDeleteProductPriceByID(ctx context.Context, priceID uint64) error {
	return svc.basicProductPriceAuthorPermission(ctx, priceID, nil)
}

func (svc *ProductService) CanUpdateProductPrice(ctx context.Context, price *domain.ProductPrice) error {
	return svc.basicProductPriceAuthorPermission(ctx, 0, price)
}

func (svc *ProductService) CanDeleteProductPrice(ctx context.Context, price *domain.ProductPrice) error {
	return svc.basicProductPriceAuthorPermission(ctx, 0, price)
}

func (svc *ProductService) basicProductPriceAuthorPermission(ctx context.Context, priceID uint64, price *domain.ProductPrice) (err error) {
	if price == nil {
		price, err = svc.ProductPriceRepository.Find(ctx, priceID)
		if err != nil {
			return errors.NewForbiddenError("forbidden: permission denied")
		}
	}

	user := metadata.GetUserFromContext(ctx)
	if price.UserID != user.ID {
		return errors.NewForbiddenError("forbidden: permission denied")
	}

	return nil
}
