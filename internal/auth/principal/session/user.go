package session

import (
	"context"
	"strings"

	"github.com/friendsofgo/errors"
	"github.com/golang-jwt/jwt"
)

const XSessionHeader = "x-session"

type Token string

func (t Token) Type() string {
	return XSessionHeader
}

func (t Token) Value() string {
	return string(t)
}

type User struct {
	ChannelID string `json:"channelId"`
	ID        string `json:"id"`
}

type UserPrincipal struct {
	User
	Token Token
}

type UserJwt struct {
	jwt.StandardClaims
	Key string
}

type UserFetcher interface {
	FetchUser(ctx context.Context, token string) (UserPrincipal, error)
}

type UserFetcherImpl struct {
	jwtServiceKey string
}

func NewUserFetcherImpl(jwtServiceKey string) *UserFetcherImpl {
	return &UserFetcherImpl{jwtServiceKey: jwtServiceKey}
}

func (f *UserFetcherImpl) FetchUser(ctx context.Context, token string) (UserPrincipal, error) {
	parsed, err := jwt.ParseWithClaims(token, &UserJwt{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return f.jwtServiceKey, nil
		}
		return UserPrincipal{}, errors.New("signing method differs")
	})
	if err != nil {
		return UserPrincipal{}, err
	}
	if !parsed.Valid {
		return UserPrincipal{}, errors.New("token invalid")
	}

	claims := parsed.Claims.(*UserJwt)

	if claims.Issuer != "ses" {
		return UserPrincipal{}, errors.New("invalid iss")
	}

	userID, channelID, ok := strings.Cut(claims.Key, "-")
	if !ok {
		return UserPrincipal{}, errors.New("invalid Key")
	}
	return UserPrincipal{User: User{ID: userID, ChannelID: channelID}, Token: Token(token)}, nil
}
