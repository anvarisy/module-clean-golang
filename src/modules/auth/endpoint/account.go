package endpoint

import (
	"context"

	helperjwt "anvarisy.tech/cleangolang/src/helper/jwt"
	"anvarisy.tech/cleangolang/src/modules/auth/domain/entity"
	"anvarisy.tech/cleangolang/src/modules/auth/domain/usecase"
	"github.com/spf13/viper"
)

type (
	Endpoint struct {
		business usecase.BusinessInterface
		repo     usecase.RepositoryInterface
	}
)

func Init(b usecase.BusinessInterface, r usecase.RepositoryInterface) Endpoint {
	return Endpoint{
		business: b,
		repo:     r,
	}
}

func (ep *Endpoint) Login(context context.Context, param LoginRequest) (response LoginResponse, err error) {
	var (
		account = entity.Account{}
	)
	account.Email = param.Email
	account.Password = param.Password
	account, err = ep.business.Login(context, account)
	if err != nil {
		return
	}
	var accessToken, expiredIn, _ = helperjwt.GenerateToken(helperjwt.Params{
		SecretKey:  []byte(viper.GetString("AUTH_ACCESS_TOKEN_SECRET")),
		Expiration: viper.GetInt64("AUTH_ACCESS_TOKEN_EXPIRATION"),
	}, account.Id, true)
	// generate refresh token
	var refreshToken, _, _ = helperjwt.GenerateToken(helperjwt.Params{
		SecretKey:  []byte(viper.GetString("AUTH_REFRESH_TOKEN_SECRET")),
		Expiration: viper.GetInt64("AUTH_REFRESH_TOKEN_EXPIRATION"),
	}, account.Id, true)
	response.AccessToken = accessToken
	response.RefreshToken = refreshToken
	response.ExpiredIn = expiredIn.(int64)
	return
}
