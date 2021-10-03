package repositories

import (
	"context"

	"gitlab.com/trivery-id/skadi/internal/product/domain"
	"gitlab.com/trivery-id/skadi/utils/errors"
	"gitlab.com/trivery-id/skadi/utils/logger"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		DB: db,
	}
}

var (
	errAddProduct     = "failed to add new product to database"
	errUpdateProduct  = "failed to update product in database"
	errFindProduct    = "failed to find product in database"
	errFindAllProduct = "failed to find all products in database"
	errDeleteProduct  = "failed to delete product from database"
)

func (r *ProductRepository) Add(ctx context.Context, product *domain.Product) error {
	log := logger.GetLoggerWithCtx(ctx).With(
		zap.String("function", "ProductRepository.Add"),
	)

	if err := r.DB.Create(&product).Error; err != nil {
		log.Error(errAddProduct, zap.Error(err))
		return NewRepositoryError(errAddProduct, err)
	}

	return nil
}

func (r *ProductRepository) Update(ctx context.Context, product *domain.Product) error {
	log := logger.GetLoggerWithCtx(ctx).With(
		zap.String("function", "ProductRepository.Update"),
		zap.Uint64("productID", product.ID),
	)

	if err := r.DB.Where("id = ?", product.ID).Save(&product).Error; err != nil {
		log.Error(errUpdateProduct, zap.Error(err))
		return NewRepositoryError(errUpdateProduct, err)
	}

	return nil
}

func (r *ProductRepository) Find(ctx context.Context, productID uint64) (*domain.Product, error) {
	log := logger.GetLoggerWithCtx(ctx).With(
		zap.String("function", "ProductRepository.Find"),
		zap.Uint64("productID", productID),
	)

	product := domain.Product{}
	tx := r.DB.
		Where("id = ?", productID).
		Preload("Locations").Preload("Prices").
		First(&product)

	if err := tx.Error; err != nil {
		err = NewRepositoryError(errFindProduct, err)
		if !errors.IsNotFoundError(err) {
			log.Error(errFindProduct, zap.Error(err))
		}

		return nil, err
	}

	return &product, nil
}

func (r *ProductRepository) FindAll(ctx context.Context, in FindAllInput) ([]domain.Product, error) {
	log := logger.GetLoggerWithCtx(ctx).With(
		zap.String("function", "ProductRepository.FindAll"),
		zap.Any("input", in),
	)

	in.FillDefault()

	products := []domain.Product{}
	tx := r.DB.
		Where(in.Where()).
		Limit(*in.Limit).Offset(*in.Offset).
		Preload("Locations").Preload("Prices").
		Find(&products)

	if err := tx.Error; err != nil {
		log.Error(errFindAllProduct, zap.Error(err))
		return nil, NewRepositoryError(errFindAllProduct, err)
	}

	return products, nil
}

func (r *ProductRepository) Delete(ctx context.Context, productID uint64) error {
	log := logger.GetLoggerWithCtx(ctx).With(
		zap.String("function", "ProductRepository.Delete"),
		zap.Uint64("productID", productID),
	)

	if err := r.DB.Where("id = ?", productID).Delete(&domain.Product{}).Error; err != nil {
		log.Error(errDeleteProduct, zap.Error(err))
		return NewRepositoryError(errDeleteProduct, err)
	}

	return nil
}
