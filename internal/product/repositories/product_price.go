package repositories

import (
	"context"

	"gitlab.com/trivery-id/skadi/internal/product/domain"
	"gitlab.com/trivery-id/skadi/utils/errors"
	"gitlab.com/trivery-id/skadi/utils/logger"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ProductPriceRepository struct {
	DB *gorm.DB
}

func NewProductPriceRepository(db *gorm.DB) *ProductPriceRepository {
	return &ProductPriceRepository{
		DB: db,
	}
}

var (
	errAddProductPrice     = "failed to add new product price to database"
	errUpdateProductPrice  = "failed to update product price in database"
	errFindProductPrice    = "failed to find product price in database"
	errFindAllProductPrice = "failed to find all product prices in database"
	errDeleteProductPrice  = "failed to delete product price from database"
)

func (r *ProductPriceRepository) Add(ctx context.Context, price *domain.ProductPrice) error {
	log := logger.GetLoggerWithCtx(ctx).With(
		zap.String("function", "ProductPriceRepository.Add"),
	)

	if err := r.DB.Create(&price).Error; err != nil {
		log.Error(errAddProductPrice, zap.Error(err))
		return NewRepositoryError(errAddProductPrice, err)
	}

	return nil
}

func (r *ProductPriceRepository) Update(ctx context.Context, price *domain.ProductPrice) error {
	log := logger.GetLoggerWithCtx(ctx).With(
		zap.String("function", "ProductPriceRepository.Update"),
		zap.Uint64("priceID", price.ID),
	)

	if err := r.DB.Where("id = ?", price.ID).Save(&price).Error; err != nil {
		log.Error(errUpdateProductPrice, zap.Error(err))
		return NewRepositoryError(errUpdateProductPrice, err)
	}

	return nil
}

func (r *ProductPriceRepository) Find(ctx context.Context, priceID uint64) (*domain.ProductPrice, error) {
	log := logger.GetLoggerWithCtx(ctx).With(
		zap.String("function", "ProductPriceRepository.Find"),
		zap.Uint64("priceID", priceID),
	)

	price := domain.ProductPrice{}
	if err := r.DB.Where("id = ?", priceID).First(&price).Error; err != nil {
		err = NewRepositoryError(errFindProductPrice, err)
		if !errors.IsNotFoundError(err) {
			log.Error(errFindProductPrice, zap.Error(err))
		}

		return nil, err
	}

	return &price, nil
}

func (r *ProductPriceRepository) FindAll(ctx context.Context, in FindAllInput) ([]domain.ProductPrice, error) {
	log := logger.GetLoggerWithCtx(ctx).With(
		zap.String("function", "ProductPriceRepository.FindAll"),
		zap.Any("input", in),
	)

	prices := []domain.ProductPrice{}

	in.FillDefault()
	if err := r.DB.Where(in.Where()).Limit(*in.Limit).Offset(*in.Offset).Find(&prices).Error; err != nil {
		log.Error(errFindAllProductPrice, zap.Error(err))
		return nil, NewRepositoryError(errFindAllProductPrice, err)
	}

	return prices, nil
}

func (r *ProductPriceRepository) Delete(ctx context.Context, priceID uint64) error {
	log := logger.GetLoggerWithCtx(ctx).With(
		zap.String("function", "ProductPriceRepository.Delete"),
		zap.Uint64("priceID", priceID),
	)

	if err := r.DB.Where("id = ?", priceID).Delete(&domain.ProductPrice{}).Error; err != nil {
		log.Error(errDeleteProductPrice, zap.Error(err))
		return NewRepositoryError(errDeleteProductPrice, err)
	}

	return nil
}
