package repositories

import (
	"context"

	"gitlab.com/trivery-id/skadi/internal/product/domain"
	"gitlab.com/trivery-id/skadi/utils/errors"
	"gitlab.com/trivery-id/skadi/utils/logger"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type OrderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{
		DB: db,
	}
}

var (
	errAddOrder     = "failed to add new order to database"
	errUpdateOrder  = "failed to update order in database"
	errFindOrder    = "failed to find order in database"
	errFindAllOrder = "failed to find all orders in database"
	errDeleteOrder  = "failed to delete order from database"
)

func (r *OrderRepository) Add(ctx context.Context, order *domain.Order) error {
	log := logger.GetLoggerWithCtx(ctx).With(
		zap.String("function", "OrderRepository.Add"),
	)

	if err := r.DB.Omit("Cancellations").Create(&order).Error; err != nil {
		log.Error(errAddOrder, zap.Error(err))
		return NewRepositoryError(errAddOrder, err)
	}

	return nil
}

func (r *OrderRepository) Update(ctx context.Context, order *domain.Order) error {
	log := logger.GetLoggerWithCtx(ctx).With(
		zap.String("function", "OrderRepository.Update"),
		zap.Uint64("orderID", order.ID),
	)

	if err := r.DB.Where("id = ?", order.ID).Save(&order).Error; err != nil {
		log.Error(errUpdateOrder, zap.Error(err))
		return NewRepositoryError(errUpdateOrder, err)
	}

	return nil
}

func (r *OrderRepository) Find(ctx context.Context, orderID uint64) (*domain.Order, error) {
	log := logger.GetLoggerWithCtx(ctx).With(
		zap.String("function", "OrderRepository.Find"),
		zap.Uint64("orderID", orderID),
	)

	order := domain.Order{}
	tx := r.DB.
		Where("id = ?", orderID).
		Preload("Price").
		First(&order)

	if err := tx.Error; err != nil {
		err = NewRepositoryError(errFindOrder, err)
		if !errors.IsNotFoundError(err) {
			log.Error(errFindOrder, zap.Error(err))
		}

		return nil, err
	}

	return &order, nil
}

func (r *OrderRepository) FindAllByUserID(ctx context.Context, userID uint64, in FindAllInput) ([]domain.Order, error) {
	log := logger.GetLoggerWithCtx(ctx).With(
		zap.String("function", "OrderRepository.FindAll"),
		zap.Any("input", in),
	)

	in.FillDefault()

	orders := []domain.Order{}
	tx := r.DB.
		Where("requester_id = ? OR shopper_id = ?", userID, userID).
		Where(in.Where()).
		Limit(*in.Limit).Offset(*in.Offset).
		Preload("Price").
		Find(&orders)

	if err := tx.Error; err != nil {
		log.Error(errFindAllOrder, zap.Error(err))
		return nil, NewRepositoryError(errFindAllOrder, err)
	}

	return orders, nil
}

func (r *OrderRepository) Delete(ctx context.Context, orderID uint64) error {
	log := logger.GetLoggerWithCtx(ctx).With(
		zap.String("function", "OrderRepository.Delete"),
		zap.Uint64("orderID", orderID),
	)

	if err := r.DB.Where("id = ?", orderID).Delete(&domain.Order{}).Error; err != nil {
		log.Error(errDeleteOrder, zap.Error(err))
		return NewRepositoryError(errDeleteOrder, err)
	}

	return nil
}
