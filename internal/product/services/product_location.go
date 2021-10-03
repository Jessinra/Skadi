package services

import (
	"context"

	"gitlab.com/trivery-id/skadi/internal/product/domain"
	"gitlab.com/trivery-id/skadi/utils/metadata"
)

func (svc *ProductService) CreateNewProductLocation(ctx context.Context, in CreateNewProductLocationInput) (*domain.ProductLocation, error) {
	user := metadata.GetUserFromContext(ctx)

	in.FillDefault()
	if err := in.Validate(); err != nil {
		return nil, err
	}

	product, err := svc.ProductRepository.Find(ctx, in.ProductID)
	if err != nil {
		return nil, err
	}

	location := &domain.ProductLocation{
		UserID:    user.ID,
		ProductID: product.ID,
		Text:      in.Text,
		Country:   in.Country,
		Province:  in.Province,
		City:      in.City,
		Area:      *in.Area,
		Street:    *in.Street,
		Building:  *in.Building,
		Store:     *in.Store,
		Longitude: *in.Longitude,
		Latitude:  *in.Latitude,
	}
	if err := svc.ProductLocationRepository.Add(ctx, location); err != nil {
		return nil, err
	}

	return location, nil
}

func (svc *ProductService) UpdateProductLocation(ctx context.Context, in UpdateProductLocationInput) (*domain.ProductLocation, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}

	location, err := svc.ProductLocationRepository.Find(ctx, in.ID)
	if err != nil {
		return nil, err
	}

	if err := svc.CanUpdateProductLocation(ctx, location); err != nil {
		return nil, err
	}

	if in.Text != nil {
		location.Text = *in.Text
	}
	if in.Province != nil {
		location.Province = *in.Province
	}
	if in.City != nil {
		location.City = *in.City
	}
	if in.Area != nil {
		location.Area = *in.Area
	}
	if in.Street != nil {
		location.Street = *in.Street
	}
	if in.Building != nil {
		location.Building = *in.Building
	}
	if in.Store != nil {
		location.Store = *in.Store
	}
	if in.Longitude != nil {
		location.Longitude = *in.Longitude
	}
	if in.Latitude != nil {
		location.Latitude = *in.Latitude
	}
	if err := svc.ProductLocationRepository.Update(ctx, location); err != nil {
		return nil, err
	}

	return location, nil
}

func (svc *ProductService) DeleteProductLocation(ctx context.Context, locationID uint64) error {
	if err := svc.CanDeleteProductLocationByID(ctx, locationID); err != nil {
		return err
	}

	return svc.ProductLocationRepository.Delete(ctx, locationID)
}
