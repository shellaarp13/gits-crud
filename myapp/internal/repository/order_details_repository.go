package repository

import (
	"context"
	"myapp/entity"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// Order_DetailsRepository connects entity.Order_Details with database.
type Order_DetailsRepository struct {
	db *gorm.DB
}

// NewOrder_DetailsRepository creates an instance of RoleRepository.
func NewOrder_DetailsRepository(db *gorm.DB) *Order_DetailsRepository {
	return &Order_DetailsRepository{
		db: db,
	}
}

// Insert inserts customer data to database.
func (repo *Order_DetailsRepository) Insert(ctx context.Context, ent *entity.Order_Details) error {
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Order_Details{}).
		Create(ent).
		Error; err != nil {
		return errors.Wrap(err, "[Order_DetailsRepository-Insert]")
	}
	return nil
}

func (repo *Order_DetailsRepository) GetListOrder_Details(ctx context.Context, limit, offset string) ([]*entity.Order_Details, error) {
	var models []*entity.Order_Details
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Order_Details{}).
		Find(&models).
		Error; err != nil {
		return nil, errors.Wrap(err, "[Order_DetailsRepository-FindAll]")
	}
	return models, nil
}

func (repo *Order_DetailsRepository) GetDetailOrder_Details(ctx context.Context, ID uuid.UUID) (*entity.Order_Details, error) {
	var models *entity.Order_Details
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Order_Details{}).
		Take(&models, ID).
		Error; err != nil {
		return nil, errors.Wrap(err, "[Order_DetailsRepository-FindById]")
	}
	return models, nil
}

func (repo *Order_DetailsRepository) DeleteOrder_Details(ctx context.Context, ID uuid.UUID) error {
	if err := repo.db.
		WithContext(ctx).
		Delete(&entity.Order_Details{Order_details_id: ID}).Error; err != nil {
		return errors.Wrap(err, "[Order_DetailsRepository-Delete]")
	}
	return nil
}

func (repo *Order_DetailsRepository) UpdateOrder_Details(ctx context.Context, ent *entity.Order_Details) error {
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Order_Details{Order_details_id: ent.Order_details_id}).
		Select("order_number", "product_id", "quantity_product").
		Updates(ent).Error; err != nil {
		return errors.Wrap(err, "[Order_DetailsRepository-Update]")
	}
	return nil
}
