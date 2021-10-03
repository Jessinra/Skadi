package services

import (
	"context"

	"gitlab.com/trivery-id/skadi/internal/product/domain"
	"gitlab.com/trivery-id/skadi/utils/errors"
	"gitlab.com/trivery-id/skadi/utils/metadata"
)

func (svc *ProductService) CanUpdateProductLocationByID(ctx context.Context, locationID uint64) error {
	return svc.basicProductLocationAuthorPermission(ctx, locationID, nil)
}

func (svc *ProductService) CanDeleteProductLocationByID(ctx context.Context, locationID uint64) error {
	return svc.basicProductLocationAuthorPermission(ctx, locationID, nil)
}

func (svc *ProductService) CanUpdateProductLocation(ctx context.Context, location *domain.ProductLocation) error {
	return svc.basicProductLocationAuthorPermission(ctx, 0, location)
}

func (svc *ProductService) CanDeleteProductLocation(ctx context.Context, location *domain.ProductLocation) error {
	return svc.basicProductLocationAuthorPermission(ctx, 0, location)
}

func (svc *ProductService) basicProductLocationAuthorPermission(ctx context.Context, locationID uint64, location *domain.ProductLocation) (err error) {
	if location == nil {
		location, err = svc.ProductLocationRepository.Find(ctx, locationID)
		if err != nil {
			return errors.NewForbiddenError("forbidden: permission denied")
		}
	}

	user := metadata.GetUserFromContext(ctx)
	if location.UserID != user.ID {
		return errors.NewForbiddenError("forbidden: permission denied")
	}

	return nil
}
