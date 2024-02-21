package general_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	mockgeneral "github.com/channel-io/ch-app-store/generated/mock/auth/general"
	"github.com/channel-io/ch-app-store/internal/auth/general"
	"github.com/channel-io/ch-proto/auth/v1/go/model"
	"github.com/channel-io/ch-proto/auth/v1/go/service"
)

var (
	jwtServiceKey = []byte("testkey")
	signingMethod = jwt.SigningMethodHS256
	userID        = "65d34af26b2fcb86812a"
	channelID     = "6170"
	tokenRawJson  = []byte(`
	{
  		"roleId": "3",
		"scope": [
			"channel-6170",
			"user-65d34af26b2fcb86812a"
  		],
 	 	"identity": "user-65d34af26b2fcb86812a"
	}`)
	testClaim = model.Claim{
		Action:  "writeUserChatMessage",
		Service: "api.channel.io",
		Scope:   []string{"channel-{id}", "user-{id}"},
	}
)

func testToken(t *testing.T) string {
	m := make(map[string]any)
	if err := json.Unmarshal(tokenRawJson, &m); err != nil {
		t.Error(err)
	}
	token := jwt.NewWithClaims(signingMethod, jwt.MapClaims(m))
	tokenStr, err := token.SignedString(jwtServiceKey)
	if err != nil {
		t.Error(err)
	}
	return tokenStr
}

func mockRoleFetcher(t *testing.T) general.RoleFetcher {
	mockFetcher := mockgeneral.NewRoleFetcher(t)
	mockFetcher.EXPECT().GetRole(mock.Anything, mock.Anything).Return(&service.GetRoleResult{
		Role: &model.Role{
			Claims: []*model.Claim{&testClaim},
		},
	}, nil)
	return mockFetcher
}

func testParser(t *testing.T) *general.ParserImpl {
	return general.NewParser(mockRoleFetcher(t), string(jwtServiceKey))
}

func TestCheckAction(t *testing.T) {
	parsed, err := testParser(t).Parse(context.Background(), testToken(t))
	if err != nil {
		t.Error(err)
	}

	assert.True(t, parsed.CheckAction(general.Service(testClaim.Service), general.Action(testClaim.Action)))
}

func TestCheckScope(t *testing.T) {
	parsed, err := testParser(t).Parse(context.Background(), testToken(t))
	if err != nil {
		panic(err)
	}

	assert.True(t, parsed.CheckScope("channel", channelID))
	assert.True(t, parsed.CheckScope("user", userID))
}

func TestIdentity(t *testing.T) {
	parsed, err := testParser(t).Parse(context.Background(), testToken(t))
	if err != nil {
		panic(err)
	}

	assert.Equal(t, "user", parsed.Type)
	assert.Equal(t, userID, parsed.ID)
}
