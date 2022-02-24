package usecase

import (
	"context"

	"anvarisy.tech/cleangolang/src/modules/auth/domain/entity"
)

type RepositoryInterface interface {
	GetAccountByEmail(context context.Context, email string) (entity.Account, error)
	CreateAccount(context context.Context, account entity.Account) error
}

type BusinessInterface interface {
	Login(context context.Context, account entity.Account) error
	Register(context context.Context, account entity.Account) error
}
