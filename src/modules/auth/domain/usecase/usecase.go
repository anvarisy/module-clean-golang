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
	Login(context context.Context, account entity.Account) (entity.Account, error)
	Register(context context.Context, account entity.Account) error
}

type Usecase struct {
}

func Init() BusinessInterface {
	return &Usecase{}
}

func (uc *Usecase) Login(context context.Context, account entity.Account) (acc entity.Account, err error) {
	//TO DO
	/*
		@ Rule of main business
	*/
	return
}
func (uc *Usecase) Register(context context.Context, account entity.Account) (err error) {
	//TO DO
	/*
		@ Rule of main business
	*/
	return
}
