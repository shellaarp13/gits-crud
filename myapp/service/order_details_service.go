package service

import (
	"context"
	"myapp/entity"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

var (
	// ErrNilOrder_Details occurs when a nil order_details is passed.
	ErrNilOrder_Details = errors.New("order_details is nil")
)

// Order_DetailsService responsible for any flow related to order_details.
// It also implements Order_DetailsService.
type Order_DetailsService struct {
	order_detailsRepo Order_DetailsRepository
}

// NewOrder_DetailsService creates an instance of Order_DetailsService.
func NewOrder_DetailsService(order_detailsRepo Order_DetailsRepository) *Order_DetailsService {
	return &Order_DetailsService{
		order_detailsRepo: order_detailsRepo,
	}
}

type Order_DetailsUseCase interface {
	Create(ctx context.Context, order_details *entity.Order_Details) error
	GetListOrder_Details(ctx context.Context, limit, offset string) ([]*entity.Order_Details, error)
	GetDetailOrder_Details(ctx context.Context, ID uuid.UUID) (*entity.Order_Details, error)
	UpdateOrder_Details(ctx context.Context, order_details *entity.Order_Details) error
	DeleteOrder_Details(ctx context.Context, ID uuid.UUID) error
}

type Order_DetailsRepository interface {
	Insert(ctx context.Context, order_details *entity.Order_Details) error
	GetListOrder_Details(ctx context.Context, limit, offset string) ([]*entity.Order_Details, error)
	GetDetailOrder_Details(ctx context.Context, ID uuid.UUID) (*entity.Order_Details, error)
	UpdateOrder_Details(ctx context.Context, order_details *entity.Order_Details) error
	DeleteOrder_Details(ctx context.Context, ID uuid.UUID) error
}

func (svc Order_DetailsService) Create(ctx context.Context, order_details *entity.Order_Details) error {
	// Checking nil order_details
	if order_details == nil {
		return ErrNilOrder_Details
	}

	// Generate id if nil
	if order_details.Order_details_id == uuid.Nil {
		order_details.Order_details_id = uuid.New()
	}

	if err := svc.order_detailsRepo.Insert(ctx, order_details); err != nil {
		return errors.Wrap(err, "[Order_DetailsService-Create]")
	}
	return nil
}

func (svc Order_DetailsService) GetListOrder_Details(ctx context.Context, limit, offset string) ([]*entity.Order_Details, error) {
	order_details, err := svc.order_detailsRepo.GetListOrder_Details(ctx, limit, offset)
	if err != nil {
		return nil, errors.Wrap(err, "[Order_DetailsService-List]")
	}
	return order_details, nil
}

func (svc Order_DetailsService) GetDetailOrder_Details(ctx context.Context, ID uuid.UUID) (*entity.Order_Details, error) {
	order_details, err := svc.order_detailsRepo.GetDetailOrder_Details(ctx, ID)
	if err != nil {
		return nil, errors.Wrap(err, "[Order_DetailsService-Detail]")
	}
	return order_details, nil
}

func (svc Order_DetailsService) DeleteOrder_Details(ctx context.Context, ID uuid.UUID) error {
	err := svc.order_detailsRepo.DeleteOrder_Details(ctx, ID)
	if err != nil {
		return errors.Wrap(err, "[Order_DetailsService-Delete]")
	}
	return nil
}

func (svc Order_DetailsService) UpdateOrder_Details(ctx context.Context, order_details *entity.Order_Details) error {
	// Checking nil order_details
	if order_details == nil {
		return ErrNilOrder_Details
	}

	// Generate id if nil
	if order_details.Order_details_id == uuid.Nil {
		order_details.Order_details_id = uuid.New()
	}

	if err := svc.order_detailsRepo.UpdateOrder_Details(ctx, order_details); err != nil {
		return errors.Wrap(err, "[Order_DetailsService-Update]")
	}
	return nil
}
