package services

import (
	"context"

	"gitlab.com/trivery-id/skadi/internal/product/domain"
	"gitlab.com/trivery-id/skadi/internal/product/repositories"
	"gitlab.com/trivery-id/skadi/utils/errors"
	"gitlab.com/trivery-id/skadi/utils/metadata"
)

func (svc *ProductService) CreateNewOrder(ctx context.Context, in CreateNewOrderInput) (*domain.Order, error) {
	user := metadata.GetUserFromContext(ctx)

	if err := in.Validate(); err != nil {
		return nil, err
	}

	order := &domain.Order{
		RequesterID: user.ID,
		ProductID:   in.ProductID,
		PriceID:     in.PriceID,

		Quantity: in.Quantity,
		Unit:     in.Unit,
		Notes:    in.Notes,

		Deal: domain.OrderDeal{
			Location:   in.Deal.Location,
			Time:       in.Deal.Time,
			Method:     in.Deal.Method,
			IncludeBox: in.Deal.IncludeBox,
		},
		State: domain.NewOrderState(),
	}
	if err := svc.OrderRepository.Add(ctx, order); err != nil {
		return nil, err
	}

	return order, nil
}

func (svc *ProductService) TakeOrder(ctx context.Context, in TakeOrderInput) error {
	user := metadata.GetUserFromContext(ctx)

	if err := in.Validate(); err != nil {
		return err
	}

	order, err := svc.OrderRepository.Find(ctx, in.OrderID)
	if err != nil {
		return err
	}

	// TODO: implement mutex
	if err := order.AcceptedBy(user.ID); err != nil {
		return err
	}

	return svc.OrderRepository.Update(ctx, order)
}

func (svc *ProductService) DropOrder(ctx context.Context, in DropOrderInput) error {
	user := metadata.GetUserFromContext(ctx)

	if err := in.Validate(); err != nil {
		return err
	}

	order, err := svc.OrderRepository.Find(ctx, in.OrderID)
	if err != nil {
		return err
	}

	if err := order.Drop(user.ID, in.Reason); err != nil {
		return err
	}

	return svc.OrderRepository.Update(ctx, order)
}

func (svc *ProductService) GetOrder(ctx context.Context, in GetOrderInput) (*domain.Order, error) {
	user := metadata.GetUserFromContext(ctx)

	if err := in.Validate(); err != nil {
		return nil, err
	}

	order, err := svc.OrderRepository.Find(ctx, in.OrderID)
	if err != nil {
		return nil, err
	}

	if order.RequesterID != user.ID || order.ShopperID != user.ID {
		return nil, errors.NewForbiddenError("not your order")
	}

	return order, nil
}

func (svc *ProductService) GetAllOrders(ctx context.Context, in GetAllOrdersInput) ([]domain.Order, error) {
	user := metadata.GetUserFromContext(ctx)

	if err := in.Validate(); err != nil {
		return nil, err
	}

	return svc.OrderRepository.FindAllByUserID(ctx, user.ID, repositories.FindAllInput{
		Limit:  in.Limit,
		Offset: in.Offset,
	})
}

func (svc *ProductService) DeleteOrder(ctx context.Context, in DeleteOrderInput) error {
	user := metadata.GetUserFromContext(ctx)

	if err := in.Validate(); err != nil {
		return err
	}

	order, err := svc.OrderRepository.Find(ctx, in.OrderID)
	if err != nil {
		return err
	}

	if order.ShopperID != user.ID {
		return errors.NewForbiddenError("not your order")
	}
	if !order.IsDeletable() {
		return errors.NewForbiddenError("can't no longer delete order")
	}

	return svc.OrderRepository.Delete(ctx, order.ID)
}
