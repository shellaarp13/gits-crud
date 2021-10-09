package service

import (
	"context"
	"myapp/entity"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

var (
	// ErrNilCustomer occurs when a nil customer is passed.
	ErrNilCustomer = errors.New("customer is nil")
)

// CustomerService responsible for any flow related to customer.
// It also implements CustomerService.
type CustomerService struct {
	customerRepo CustomerRepository
}

// NewCustomerService creates an instance of CustomerService.
func NewCustomerService(customerRepo CustomerRepository) *CustomerService {
	return &CustomerService{
		customerRepo: customerRepo,
	}
}

type CustomerUseCase interface {
	Create(ctx context.Context, customer *entity.Customer) error
	GetListCustomer(ctx context.Context, limit, offset string) ([]*entity.Customer, error)
	GetDetailCustomer(ctx context.Context, ID uuid.UUID) (*entity.Customer, error)
	UpdateCustomer(ctx context.Context, customer *entity.Customer) error
	DeleteCustomer(ctx context.Context, ID uuid.UUID) error
}

type CustomerRepository interface {
	Insert(ctx context.Context, customer *entity.Customer) error
	GetListCustomer(ctx context.Context, limit, offset string) ([]*entity.Customer, error)
	GetDetailCustomer(ctx context.Context, ID uuid.UUID) (*entity.Customer, error)
	UpdateCustomer(ctx context.Context, customer *entity.Customer) error
	DeleteCustomer(ctx context.Context, ID uuid.UUID) error
}

func (svc CustomerService) Create(ctx context.Context, customer *entity.Customer) error {
	// Checking nil customer
	if customer == nil {
		return ErrNilCustomer
	}

	// Generate id if nil
	if customer.Customer_ID == uuid.Nil {
		customer.Customer_ID = uuid.New()
	}

	if err := svc.customerRepo.Insert(ctx, customer); err != nil {
		return errors.Wrap(err, "[CustomerService-Create]")
	}
	return nil
}

func (svc CustomerService) GetListCustomer(ctx context.Context, limit, offset string) ([]*entity.Customer, error) {
	customer, err := svc.customerRepo.GetListCustomer(ctx, limit, offset)
	if err != nil {
		return nil, errors.Wrap(err, "[CustomerService-Create]")
	}
	return customer, nil
}

func (svc CustomerService) GetDetailCustomer(ctx context.Context, ID uuid.UUID) (*entity.Customer, error) {
	customer, err := svc.customerRepo.GetDetailCustomer(ctx, ID)
	if err != nil {
		return nil, errors.Wrap(err, "[CustomerService-Create]")
	}
	return customer, nil
}

func (svc CustomerService) DeleteCustomer(ctx context.Context, ID uuid.UUID) error {
	err := svc.customerRepo.DeleteCustomer(ctx, ID)
	if err != nil {
		return errors.Wrap(err, "[CustomerService-Create]")
	}
	return nil
}

func (svc CustomerService) UpdateCustomer(ctx context.Context, customer *entity.Customer) error {
	// Checking nil customer
	if customer == nil {
		return ErrNilCustomer
	}

	// Generate id if nil
	if customer.Customer_ID == uuid.Nil {
		customer.Customer_ID = uuid.New()
	}

	if err := svc.customerRepo.UpdateCustomer(ctx, customer); err != nil {
		return errors.Wrap(err, "[CustomerService-Update]")
	}
	return nil
}
