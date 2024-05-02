package account

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

var accountID = "123"
var jwtServiceKey = "serviceKey"
var tokenRawJson = json.RawMessage(fmt.Sprintf(`{ "key": "%s" }`, accountID))

func getToken(t *testing.T) string {
	m := make(map[string]any)
	if err := json.Unmarshal(tokenRawJson, &m); err != nil {
		t.Error(err)
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(m))
	tokenStr, err := token.SignedString([]byte(jwtServiceKey))
	if err != nil {
		t.Error(err)
	}
	return tokenStr
}

func TestParse(t *testing.T) {
	token := getToken(t)
	accountParser := NewAccountParser(jwtServiceKey)

	principal, err := accountParser.ParseAccount(context.Background(), token)

	assert.Nil(t, err)
	assert.Equal(t, principal.Token, Token(token))
	assert.Equal(t, principal.Account, Account(accountID))
}
