package repository

import (
	"context"
	"myapp/entity"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// CustomerRepository connects entity.Customer with database.
type CustomerRepository struct {
	db *gorm.DB
}

// NewCustomerRepository creates an instance of RoleRepository.
func NewCustomerRepository(db *gorm.DB) *CustomerRepository {
	return &CustomerRepository{
		db: db,
	}
}

// Insert inserts customer data to database.
func (repo *CustomerRepository) Insert(ctx context.Context, ent *entity.Customer) error {
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Customer{}).
		Create(ent).
		Error; err != nil {
		return errors.Wrap(err, "[CustomerRepository-Insert]")
	}
	return nil
}

func (repo *CustomerRepository) GetListCustomer(ctx context.Context, limit, offset string) ([]*entity.Customer, error) {
	var models []*entity.Customer
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Customer{}).
		Find(&models).
		Error; err != nil {
		return nil, errors.Wrap(err, "[CustomerRepository-FindAll]")
	}
	return models, nil
}

func (repo *CustomerRepository) GetDetailCustomer(ctx context.Context, ID uuid.UUID) (*entity.Customer, error) {
	var models *entity.Customer
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Customer{}).
		Take(&models, ID).
		Error; err != nil {
		return nil, errors.Wrap(err, "[CustomerRepository-FindById]")
	}
	return models, nil
}

func (repo *CustomerRepository) DeleteCustomer(ctx context.Context, ID uuid.UUID) error {
	if err := repo.db.
		WithContext(ctx).
		Delete(&entity.Customer{Customer_ID: ID}).Error; err != nil {
		return errors.Wrap(err, "[CustomerRepository-Delete]")
	}
	return nil
}

func (repo *CustomerRepository) UpdateCustomer(ctx context.Context, ent *entity.Customer) error {
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Customer{Customer_ID: ent.Customer_ID}).
		Select("first_name", "last_name", "street", "zip", "phone").
		Updates(ent).Error; err != nil {
		return errors.Wrap(err, "[CustomerRepository-Update]")
	}
	return nil
}
