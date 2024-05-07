package account

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/golang-jwt/jwt"
)

const (
	accountFetch = "/admin/account"
)

type Account struct {
	ID string `json:"id"`
}

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

func (c *ManagerFetcherImpl) ParseAccount(ctx context.Context, token string) (Principal, error) {
	req := c.cli.R()
	req.SetContext(ctx)
	req.Header.Set(XAccountHeader, token)
	resp, err := req.Get(c.authUrl + accountFetch)
	if err != nil {
		return Principal{}, err
	}

	if !isSuccess(resp.StatusCode()) {
		return Principal{}, errors.New("auth failed")
	}

	body := resp.Body()
	var acc Account
	if err := json.Unmarshal(body, &acc); err != nil {
		return Principal{}, err
	}

	return Principal{Account: acc, Token: Token(token)}, nil
}
