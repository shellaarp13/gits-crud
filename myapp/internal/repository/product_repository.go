package repository

import (
	"context"
	"myapp/entity"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// ProductRepository connects entity.Product with database.
type ProductRepository struct {
	db *gorm.DB
}

// NewProductRepository creates an instance of RoleRepository.
func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

// Insert inserts product data to database.
func (repo *ProductRepository) Insert(ctx context.Context, ent *entity.Product) error {
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Product{}).
		Create(ent).
		Error; err != nil {
		return errors.Wrap(err, "[ProductRepository-Insert]")
	}
	return nil
}

func (repo *ProductRepository) GetListProduct(ctx context.Context, limit, offset string) ([]*entity.Product, error) {
	var models []*entity.Product
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Product{}).
		Find(&models).
		Error; err != nil {
		return nil, errors.Wrap(err, "[ProductRepository-FindAll]")
	}
	return models, nil
}

func (repo *ProductRepository) GetDetailProduct(ctx context.Context, ID uuid.UUID) (*entity.Product, error) {
	var models *entity.Product
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Product{}).
		Take(&models, ID).
		Error; err != nil {
		return nil, errors.Wrap(err, "[ProductRepository-FindById]")
	}
	return models, nil
}

func (repo *ProductRepository) DeleteProduct(ctx context.Context, ID uuid.UUID) error {
	if err := repo.db.
		WithContext(ctx).
		Delete(&entity.Product{Product_ID: ID}).Error; err != nil {
		return errors.Wrap(err, "[ProductRepository-Delete]")
	}
	return nil
}

func (repo *ProductRepository) UpdateProduct(ctx context.Context, ent *entity.Product) error {
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Product{Product_ID: ent.Product_ID}).
		Select("product_name", "stock_p", "product_type", "price").
		Updates(ent).Error; err != nil {
		return errors.Wrap(err, "[ProductRepository-Update]")
	}
	return nil
}
