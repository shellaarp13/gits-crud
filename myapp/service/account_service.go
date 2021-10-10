package service

import (
	"context"
	"myapp/entity"

	"github.com/pkg/errors"
)

var (
	// ErrNilAccount occurs when a nil account is passed.
	ErrNilAccount = errors.New("account is nil")
)

// AccountService responsible for any flow related to account.
// It also implements AccountService.
type AccountService struct {
	accountRepo AccountRepository
}

// NewAccountService creates an instance of AccountService.
func NewAccountService(accountRepo AccountRepository) *AccountService {
	return &AccountService{
		accountRepo: accountRepo,
	}
}

type AccountUseCase interface {
	Create(ctx context.Context, account *entity.Account) error
	GetListAccount(ctx context.Context, limit, offset string) ([]*entity.Account, error)
	GetDetailAccount(ctx context.Context, Username string) (*entity.Account, error)
	UpdateAccount(ctx context.Context, Username string, account *entity.Account) error
	DeleteAccount(ctx context.Context, Username string) error
}

type AccountRepository interface {
	Insert(ctx context.Context, account *entity.Account) error
	GetListAccount(ctx context.Context, limit, offset string) ([]*entity.Account, error)
	GetDetailAccount(ctx context.Context, Username string) (*entity.Account, error)
	UpdateAccount(ctx context.Context, Username string, account *entity.Account) error
	DeleteAccount(ctx context.Context, Username string) error
}

func (svc AccountService) Create(ctx context.Context, account *entity.Account) error {
	// Checking nil account
	if account == nil {
		return ErrNilAccount
	}

	if err := svc.accountRepo.Insert(ctx, account); err != nil {
		return errors.Wrap(err, "[AccountService-Create]")
	}
	return nil
}

func (svc AccountService) GetListAccount(ctx context.Context, limit, offset string) ([]*entity.Account, error) {
	account, err := svc.accountRepo.GetListAccount(ctx, limit, offset)
	if err != nil {
		return nil, errors.Wrap(err, "[AccountService-List]")
	}
	return account, nil
}

func (svc AccountService) GetDetailAccount(ctx context.Context, Username string) (*entity.Account, error) {
	account, err := svc.accountRepo.GetDetailAccount(ctx, Username)
	if err != nil {
		return nil, errors.Wrap(err, "[AccountService-Detail]")
	}
	return account, nil
}

func (svc AccountService) DeleteAccount(ctx context.Context, Username string) error {
	err := svc.accountRepo.DeleteAccount(ctx, Username)
	if err != nil {
		return errors.Wrap(err, "[AccountService-Delete]")
	}
	return nil
}

func (svc AccountService) UpdateAccount(ctx context.Context, Username string, account *entity.Account) error {
	// Checking nil account
	if account == nil {
		return ErrNilAccount
	}

	if err := svc.accountRepo.UpdateAccount(ctx, Username, account); err != nil {
		return errors.Wrap(err, "[AccountService-Update]")
	}
	return nil
}
