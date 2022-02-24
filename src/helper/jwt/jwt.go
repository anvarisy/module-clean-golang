package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Params struct {
	SecretKey  []byte
	Expiration int64
}

func GenerateToken(params Params, identifier interface{}, withExpired bool) (string, interface{}, error) {
	var (
		ttl          = time.Duration(params.Expiration) * time.Second
		expiredIn    interface{}
		jwtWithClaim *jwt.Token
	)

	expiredIn = nil

	switch withExpired {
	case true:
		expiredIn = time.Now().UTC().Add(ttl).Unix()
		jwtWithClaim = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"identifier": identifier,
			"exp":        expiredIn,
		})
	default:
		jwtWithClaim = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"identifier": identifier,
		})
	}

	token, err := jwtWithClaim.SignedString(params.SecretKey)
	return token, expiredIn, err
}

func ParsingToken(secretKey []byte, tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("signing method invalid")
		} else if method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("signing method invalid")
		}

		return secretKey, nil
	})
}
