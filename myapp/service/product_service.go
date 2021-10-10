package service

import (
	"context"
	"myapp/entity"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

var (
	// ErrNilProduct occurs when a nil Product is passed.
	ErrNilProduct = errors.New("Product is nil")
)

// ProductService responsible for any flow related to customer.
// It also implements ProductService.
type ProductService struct {
	productRepo ProductRepository
}

// NewProductService creates an instance of ProductService.
func NewProductService(productRepo ProductRepository) *ProductService {
	return &ProductService{
		productRepo: productRepo,
	}
}

type ProductUseCase interface {
	Create(ctx context.Context, product *entity.Product) error
	GetListProduct(ctx context.Context, limit, offset string) ([]*entity.Product, error)
	GetDetailProduct(ctx context.Context, ID uuid.UUID) (*entity.Product, error)
	UpdateProduct(ctx context.Context, product *entity.Product) error
	DeleteProduct(ctx context.Context, ID uuid.UUID) error
}

type ProductRepository interface {
	Insert(ctx context.Context, product *entity.Product) error
	GetListProduct(ctx context.Context, limit, offset string) ([]*entity.Product, error)
	GetDetailProduct(ctx context.Context, ID uuid.UUID) (*entity.Product, error)
	UpdateProduct(ctx context.Context, product *entity.Product) error
	DeleteProduct(ctx context.Context, ID uuid.UUID) error
}

func (svc ProductService) Create(ctx context.Context, product *entity.Product) error {
	// Checking nil product
	if product == nil {
		return ErrNilProduct
	}

	// Generate id if nil
	if product.Product_ID == uuid.Nil {
		product.Product_ID = uuid.New()
	}

	if err := svc.productRepo.Insert(ctx, product); err != nil {
		return errors.Wrap(err, "[ProductService-Create]")
	}
	return nil
}

func (svc ProductService) GetListProduct(ctx context.Context, limit, offset string) ([]*entity.Product, error) {
	product, err := svc.productRepo.GetListProduct(ctx, limit, offset)
	if err != nil {
		return nil, errors.Wrap(err, "[ProductService-List]")
	}
	return product, nil
}

func (svc ProductService) GetDetailProduct(ctx context.Context, ID uuid.UUID) (*entity.Product, error) {
	product, err := svc.productRepo.GetDetailProduct(ctx, ID)
	if err != nil {
		return nil, errors.Wrap(err, "[ProductService-Detail]")
	}
	return product, nil
}

func (svc ProductService) DeleteProduct(ctx context.Context, ID uuid.UUID) error {
	err := svc.productRepo.DeleteProduct(ctx, ID)
	if err != nil {
		return errors.Wrap(err, "[ProductService-Delete]")
	}
	return nil
}

func (svc ProductService) UpdateProduct(ctx context.Context, product *entity.Product) error {
	// Checking nil product
	if product == nil {
		return ErrNilProduct
	}

	// Generate id if nil
	if product.Product_ID == uuid.Nil {
		product.Product_ID = uuid.New()
	}

	if err := svc.productRepo.UpdateProduct(ctx, product); err != nil {
		return errors.Wrap(err, "[ProductService-Update]")
	}
	return nil
}
