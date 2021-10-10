package repository

import (
	"context"
	"myapp/entity"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// OrderDetailsRepository connects entity.OrderDetails with database.
type OrderDetailsRepository struct {
	db *gorm.DB
}

// NewOrderDetailsRepository creates an instance of RoleRepository.
func NewOrderDetailsRepository(db *gorm.DB) *OrderDetailsRepository {
	return &OrderDetailsRepository{
		db: db,
	}
}

// Insert inserts order_details data to database.
func (repo *OrderDetailsRepository) Insert(ctx context.Context, ent *entity.OrderDetails) error {
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.OrderDetails{}).
		Create(ent).
		Error; err != nil {
		return errors.Wrap(err, "[OrderDetailsRepository-Insert]")
	}
	return nil
}

func (repo *OrderDetailsRepository) GetListOrderDetails(ctx context.Context, limit, offset string) ([]*entity.OrderDetails, error) {
	var models []*entity.OrderDetails
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.OrderDetails{}).
		Find(&models).
		Error; err != nil {
		return nil, errors.Wrap(err, "[OrderDetailsRepository-FindAll]")
	}
	return models, nil
}

func (repo *OrderDetailsRepository) GetDetailOrderDetails(ctx context.Context, ID uuid.UUID) (*entity.OrderDetails, error) {
	var models *entity.OrderDetails
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.OrderDetails{}).
		Take(&models, ID).
		Error; err != nil {
		return nil, errors.Wrap(err, "[OrderDetailsRepository-FindById]")
	}
	return models, nil
}

func (repo *OrderDetailsRepository) DeleteOrderDetails(ctx context.Context, ID uuid.UUID) error {
	if err := repo.db.
		WithContext(ctx).
		Delete(&entity.OrderDetails{OrderDetails_ID: ID}).Error; err != nil {
		return errors.Wrap(err, "[OrderDetailsRepository-Delete]")
	}
	return nil
}

func (repo *OrderDetailsRepository) UpdateOrderDetails(ctx context.Context, ent *entity.OrderDetails) error {
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.OrderDetails{OrderDetails_ID: ent.OrderDetails_ID}).
		Select("order_number", "product_id", "quantity_product").
		Updates(ent).Error; err != nil {
		return errors.Wrap(err, "[OrderDetailsRepository-Update]")
	}
	return nil
}
