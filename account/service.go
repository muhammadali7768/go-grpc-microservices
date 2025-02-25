package account

import (
	"context"

	"github.com/segmentio/ksuid"
)

type Service interface {
	PostAccount(ctx context.Context, name string) (*Account, error)
	GetAccount(ctx context.Context, id string) (*Account, error)
	GetAccounts(ctx context.Context, skip uint64, take uint64) ([]Account, error)
}

type Account struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type accountService struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &accountService{repo}
}

func (s *accountService) PostAccount(ctx context.Context, name string) (*Account, error) {
	id := ksuid.New().String()
	account := &Account{ID: id, Name: name}
	err := s.repo.PutAccount(ctx, *account)
	if err != nil {
		return nil, err
	}
	return account, nil
}
func (s *accountService) GetAccount(ctx context.Context, id string) (*Account, error) {
	acount, err := s.repo.GetAccountByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return acount, nil
}
func (s *accountService) GetAccounts(ctx context.Context, skip uint64, take uint64) ([]Account, error) {
	if take > 100 || (skip == 0 && take == 0) {
		take = 100
	}
	accounts, err := s.repo.ListAccounts(ctx, skip, take)
	if err != nil {
		return nil, err
	}
	return accounts, nil
}
