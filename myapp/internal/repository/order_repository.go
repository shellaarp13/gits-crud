package repository

import (
	"context"
	"myapp/entity"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// OrderRepository connects entity.Order with database.
type OrderRepository struct {
	db *gorm.DB
}

// NewOrderRepository creates an instance of RoleRepository.
func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

// Insert inserts order data to database.
func (repo *OrderRepository) Insert(ctx context.Context, ent *entity.Order) error {
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Order{}).
		Create(ent).
		Error; err != nil {
		return errors.Wrap(err, "[OrderRepository-Insert]")
	}
	return nil
}

func (repo *OrderRepository) GetListOrder(ctx context.Context, limit, offset string) ([]*entity.Order, error) {
	var models []*entity.Order
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Order{}).
		Find(&models).
		Error; err != nil {
		return nil, errors.Wrap(err, "[OrderRepository-FindAll]")
	}
	return models, nil
}

func (repo *OrderRepository) GetDetailOrder(ctx context.Context, ID uuid.UUID) (*entity.Order, error) {
	var models *entity.Order
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Order{}).
		Take(&models, ID).
		Error; err != nil {
		return nil, errors.Wrap(err, "[OrderRepository-FindById]")
	}
	return models, nil
}

func (repo *OrderRepository) DeleteOrder(ctx context.Context, Order_Number uuid.UUID) error {
	if err := repo.db.
		WithContext(ctx).
		Where("order_number = ?", Order_Number).
		Delete(&entity.Order{Order_Number: Order_Number}).Error; err != nil {
		return errors.Wrap(err, "[OrderRepository-Delete]")
	}
	return nil
}

func (repo *OrderRepository) UpdateOrder(ctx context.Context, ent *entity.Order) error {
	if err := repo.db.
		WithContext(ctx).
		Where("order_number = ?", ent.Order_Number).
		Model(&entity.Order{Order_Number: ent.Order_Number}).
		Select("customer_id", "customer_name", "to_street", "to_city", "to_zip", "ship_date").
		Updates(ent).Error; err != nil {
		return errors.Wrap(err, "[OrderRepository-Update]")
	}
	return nil
}
