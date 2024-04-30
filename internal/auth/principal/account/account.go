package account

import (
	"context"

	"github.com/friendsofgo/errors"
	"github.com/golang-jwt/jwt"
)

type Account string

type Principal struct {
	Account
	Token Token
}

type Jwt struct {
	jwt.StandardClaims
	Key string
}

type Parser interface {
	ParseAccount(ctx context.Context, token string) (Principal, error)
}

type ParserImpl struct {
	jwtServiceKey string
}

func NewAccountParser(jwtServiceKey string) *ParserImpl {
	return &ParserImpl{jwtServiceKey: jwtServiceKey}
}

func (f *ParserImpl) ParseAccount(ctx context.Context, token string) (Principal, error) {
	parsed, err := jwt.ParseWithClaims(token, &Jwt{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(f.jwtServiceKey), nil
		}
		return Principal{}, errors.New("signing method differs")
	})
	if err != nil {
		e, ok := err.(*jwt.ValidationError)
		if !ok || ok && e.Errors&jwt.ValidationErrorIssuedAt == 0 { // Don't report error that token used before issued.
			return Principal{}, err
		}
	}

	claims := parsed.Claims.(*Jwt)
	return Principal{
		Token:   Token(token),
		Account: Account(claims.Key),
	}, nil
}
