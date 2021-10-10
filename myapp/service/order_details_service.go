package service

import (
	"context"
	"myapp/entity"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

var (
	// ErrNilOrderDetails occurs when a nil order details is passed.
	ErrNilOrderDetails = errors.New("order details is nil")
)

// OrderDetailsService responsible for any flow related to OrderDetails.
// It also implements OrderDetailsService.
type OrderDetailsService struct {
	orderdetailsRepo OrderDetailsRepository
}

// NewOrderDetailsService creates an instance of OrderDetailsService.
func NewOrderDetailsService(orderdetailsRepo OrderDetailsRepository) *OrderDetailsService {
	return &OrderDetailsService{
		orderdetailsRepo: orderdetailsRepo,
	}
}

type OrderDetailsUseCase interface {
	Create(ctx context.Context, orderdetails *entity.OrderDetails) error
	GetListOrderDetails(ctx context.Context, limit, offset string) ([]*entity.OrderDetails, error)
	GetDetailOrderDetails(ctx context.Context, ID uuid.UUID) (*entity.OrderDetails, error)
	UpdateOrderDetails(ctx context.Context, orderdetails *entity.OrderDetails) error
	DeleteOrderDetails(ctx context.Context, ID uuid.UUID) error
}

type OrderDetailsRepository interface {
	Insert(ctx context.Context, orderdetails *entity.OrderDetails) error
	GetListOrderDetails(ctx context.Context, limit, offset string) ([]*entity.OrderDetails, error)
	GetDetailOrderDetails(ctx context.Context, ID uuid.UUID) (*entity.OrderDetails, error)
	UpdateOrderDetails(ctx context.Context, orderdetails *entity.OrderDetails) error
	DeleteOrderDetails(ctx context.Context, ID uuid.UUID) error
}

func (svc OrderDetailsService) Create(ctx context.Context, orderdetails *entity.OrderDetails) error {
	// Checking nil order details
	if orderdetails == nil {
		return ErrNilOrderDetails
	}

	// Generate id if nil
	if orderdetails.OrderDetails_ID == uuid.Nil {
		orderdetails.OrderDetails_ID = uuid.New()
	}

	if err := svc.orderdetailsRepo.Insert(ctx, orderdetails); err != nil {
		return errors.Wrap(err, "[OrderDetailsService-Create]")
	}
	return nil
}

func (svc OrderDetailsService) GetListOrderDetails(ctx context.Context, limit, offset string) ([]*entity.OrderDetails, error) {
	orderdetails, err := svc.orderdetailsRepo.GetListOrderDetails(ctx, limit, offset)
	if err != nil {
		return nil, errors.Wrap(err, "[OrderDetailsService-List]")
	}
	return orderdetails, nil
}

func (svc OrderDetailsService) GetDetailOrderDetails(ctx context.Context, ID uuid.UUID) (*entity.OrderDetails, error) {
	orderdetails, err := svc.orderdetailsRepo.GetDetailOrderDetails(ctx, ID)
	if err != nil {
		return nil, errors.Wrap(err, "[OrderDetailsService-Detail]")
	}
	return orderdetails, nil
}

func (svc OrderDetailsService) DeleteOrderDetails(ctx context.Context, ID uuid.UUID) error {
	err := svc.orderdetailsRepo.DeleteOrderDetails(ctx, ID)
	if err != nil {
		return errors.Wrap(err, "[OrderDetailsService-Delete]")
	}
	return nil
}

func (svc OrderDetailsService) UpdateOrderDetails(ctx context.Context, orderdetails *entity.OrderDetails) error {
	// Checking nil order details
	if orderdetails == nil {
		return ErrNilOrderDetails
	}

	// Generate id if nil
	if orderdetails.OrderDetails_ID == uuid.Nil {
		orderdetails.OrderDetails_ID = uuid.New()
	}

	if err := svc.orderdetailsRepo.UpdateOrderDetails(ctx, orderdetails); err != nil {
		return errors.Wrap(err, "[OrderDetailsService-Update]")
	}
	return nil
}
