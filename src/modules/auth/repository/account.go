package repository

import (
	"context"
	"database/sql"

	"anvarisy.tech/cleangolang/src/app"
	"anvarisy.tech/cleangolang/src/modules/auth/domain/entity"
	"anvarisy.tech/cleangolang/src/modules/auth/domain/usecase"
)

type Account struct {
	db *sql.DB
}

func Init(app *app.App) usecase.RepositoryInterface {
	var db = app.GetDb()
	return &Account{
		db: db,
	}
}

func (a *Account) GetAccountByEmail(context context.Context, email string) (entity.Account, error) {
	var account = entity.Account{}
	var err error
	return account, err
}

func (a *Account) CreateAccount(context context.Context, account entity.Account) error {
	return nil
}
