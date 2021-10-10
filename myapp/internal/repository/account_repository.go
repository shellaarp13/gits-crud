package repository

import (
	"context"
	"myapp/entity"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// AccountRepository connects entity.Account with database.
type AccountRepository struct {
	db *gorm.DB
}

// NewAccountRepository creates an instance of RoleRepository.
func NewAccountRepository(db *gorm.DB) *AccountRepository {
	return &AccountRepository{
		db: db,
	}
}

// Insert inserts account data to database.
func (repo *AccountRepository) Insert(ctx context.Context, ent *entity.Account) error {
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Account{}).
		Create(ent).
		Error; err != nil {
		return errors.Wrap(err, "[AccountRepository-Insert]")
	}
	return nil
}

func (repo *AccountRepository) GetListAccount(ctx context.Context, limit, offset string) ([]*entity.Account, error) {
	var models []*entity.Account
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Account{}).
		Find(&models).
		Error; err != nil {
		return nil, errors.Wrap(err, "[AccountRepository-FindAll]")
	}
	return models, nil
}

func (repo *AccountRepository) GetDetailAccount(ctx context.Context, Username string) (*entity.Account, error) {
	var models *entity.Account
	if err := repo.db.
		WithContext(ctx).
		Where("username = ?", Username).
		Find(&models).
		Error; err != nil {
		return nil, errors.Wrap(err, "[AccountRepository-FindById]")
	}
	return models, nil
}

func (repo *AccountRepository) DeleteAccount(ctx context.Context, Username string) error {
	if err := repo.db.
		WithContext(ctx).
		Delete(&entity.Account{Username: Username}).Error; err != nil {
		return errors.Wrap(err, "[AccountRepository-Delete]")
	}
	return nil
}

func (repo *AccountRepository) UpdateAccount(ctx context.Context, Username string, ent *entity.Account) error {
	var models *entity.Account
	if err := repo.db.
		WithContext(ctx).
		Where("username = ?", Username).
		Find(&models).
		Select("username", "password").
		Updates(ent).Error; err != nil {
		return errors.Wrap(err, "[AccountRepository-Update]")
	}
	return nil
}
