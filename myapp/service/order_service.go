package service

import (
	"context"
	"myapp/entity"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

var (
	// ErrNilOrder occurs when a nil order is passed.
	ErrNilOrder = errors.New("order is nil")
)

// OrderService responsible for any flow related to order.
// It also implements OrderService.
type OrderService struct {
	orderRepo OrderRepository
}

// NewOrderService creates an instance of OrderService.
func NewOrderService(orderRepo OrderRepository) *OrderService {
	return &OrderService{
		orderRepo: orderRepo,
	}
}

type OrderUseCase interface {
	Create(ctx context.Context, order *entity.Order) error
	GetListOrder(ctx context.Context, limit, offset string) ([]*entity.Order, error)
	GetDetailOrder(ctx context.Context, ID uuid.UUID) (*entity.Order, error)
	UpdateOrder(ctx context.Context, order *entity.Order) error
	DeleteOrder(ctx context.Context, ID uuid.UUID) error
}

type OrderRepository interface {
	Insert(ctx context.Context, order *entity.Order) error
	GetListOrder(ctx context.Context, limit, offset string) ([]*entity.Order, error)
	GetDetailOrder(ctx context.Context, ID uuid.UUID) (*entity.Order, error)
	UpdateOrder(ctx context.Context, order *entity.Order) error
	DeleteOrder(ctx context.Context, ID uuid.UUID) error
}

func (svc OrderService) Create(ctx context.Context, order *entity.Order) error {
	// Checking nil order
	if order == nil {
		return ErrNilOrder
	}

	// Generate id if nil
	if order.Order_Number == uuid.Nil {
		order.Order_Number = uuid.New()
	}

	if err := svc.orderRepo.Insert(ctx, order); err != nil {
		return errors.Wrap(err, "[OrderService-Create]")
	}
	return nil
}

func (svc OrderService) GetListOrder(ctx context.Context, limit, offset string) ([]*entity.Order, error) {
	order, err := svc.orderRepo.GetListOrder(ctx, limit, offset)
	if err != nil {
		return nil, errors.Wrap(err, "[OrderService-List]")
	}
	return order, nil
}

func (svc OrderService) GetDetailOrder(ctx context.Context, ID uuid.UUID) (*entity.Order, error) {
	order, err := svc.orderRepo.GetDetailOrder(ctx, ID)
	if err != nil {
		return nil, errors.Wrap(err, "[OrderService-Detail]")
	}
	return order, nil
}

func (svc OrderService) DeleteOrder(ctx context.Context, ID uuid.UUID) error {
	err := svc.orderRepo.DeleteOrder(ctx, ID)
	if err != nil {
		return errors.Wrap(err, "[OrderService-Delete]")
	}
	return nil
}

func (svc OrderService) UpdateOrder(ctx context.Context, order *entity.Order) error {
	// Checking nil order
	if order == nil {
		return ErrNilOrder
	}

	// Generate id if nil
	if order.Order_Number == uuid.Nil {
		order.Order_Number = uuid.New()
	}

	if err := svc.orderRepo.UpdateOrder(ctx, order); err != nil {
		return errors.Wrap(err, "[OrderService-Update]")
	}
	return nil
}
