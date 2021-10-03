package repositories

import (
	"context"

	"gitlab.com/trivery-id/skadi/internal/product/domain"
	"gitlab.com/trivery-id/skadi/utils/errors"
	"gitlab.com/trivery-id/skadi/utils/logger"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ProductLocationRepository struct {
	DB *gorm.DB
}

func NewProductLocationRepository(db *gorm.DB) *ProductLocationRepository {
	return &ProductLocationRepository{
		DB: db,
	}
}

var (
	errAddProductLocation     = "failed to add new product location to database"
	errUpdateProductLocation  = "failed to update product location in database"
	errFindProductLocation    = "failed to find product location in database"
	errFindAllProductLocation = "failed to find all product locations in database"
	errDeleteProductLocation  = "failed to delete product location from database"
)

func (r *ProductLocationRepository) Add(ctx context.Context, location *domain.ProductLocation) error {
	log := logger.GetLoggerWithCtx(ctx).With(
		zap.String("function", "ProductLocationRepository.Add"),
	)

	if err := r.DB.Create(&location).Error; err != nil {
		log.Error(errAddProductLocation, zap.Error(err))
		return NewRepositoryError(errAddProductLocation, err)
	}

	return nil
}

func (r *ProductLocationRepository) Update(ctx context.Context, location *domain.ProductLocation) error {
	log := logger.GetLoggerWithCtx(ctx).With(
		zap.String("function", "ProductLocationRepository.Update"),
		zap.Uint64("locationID", location.ID),
	)

	if err := r.DB.Where("id = ?", location.ID).Save(&location).Error; err != nil {
		log.Error(errUpdateProductLocation, zap.Error(err))
		return NewRepositoryError(errUpdateProductLocation, err)
	}

	return nil
}

func (r *ProductLocationRepository) Find(ctx context.Context, locationID uint64) (*domain.ProductLocation, error) {
	log := logger.GetLoggerWithCtx(ctx).With(
		zap.String("function", "ProductLocationRepository.Find"),
		zap.Uint64("locationID", locationID),
	)

	location := domain.ProductLocation{}
	if err := r.DB.Where("id = ?", locationID).First(&location).Error; err != nil {
		err = NewRepositoryError(errFindProductLocation, err)
		if !errors.IsNotFoundError(err) {
			log.Error(errFindProductLocation, zap.Error(err))
		}

		return nil, err
	}

	return &location, nil
}

func (r *ProductLocationRepository) FindAll(ctx context.Context, in FindAllInput) ([]domain.ProductLocation, error) {
	log := logger.GetLoggerWithCtx(ctx).With(
		zap.String("function", "ProductLocationRepository.FindAll"),
		zap.Any("input", in),
	)

	locations := []domain.ProductLocation{}

	in.FillDefault()
	if err := r.DB.Where(in.Where()).Limit(*in.Limit).Offset(*in.Offset).Find(&locations).Error; err != nil {
		log.Error(errFindAllProductLocation, zap.Error(err))
		return nil, NewRepositoryError(errFindAllProductLocation, err)
	}

	return locations, nil
}

func (r *ProductLocationRepository) Delete(ctx context.Context, locationID uint64) error {
	log := logger.GetLoggerWithCtx(ctx).With(
		zap.String("function", "ProductLocationRepository.Delete"),
		zap.Uint64("locationID", locationID),
	)

	if err := r.DB.Where("id = ?", locationID).Delete(&domain.ProductLocation{}).Error; err != nil {
		log.Error(errDeleteProductLocation, zap.Error(err))
		return NewRepositoryError(errDeleteProductLocation, err)
	}

	return nil
}
