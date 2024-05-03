package svc

import (
	app "github.com/channel-io/ch-app-store/internal/app/svc"
	"github.com/channel-io/ch-app-store/internal/approle/model"
	protomodel "github.com/channel-io/ch-proto/auth/v1/go/model"
)

type ClaimManager struct {
	typeRule map[model.RoleType]TypeRule
	appSvc   app.AppCrudSvc
}

type TypeRule struct {
	availableClaims []*protomodel.Claim
	grantTypes      []protomodel.GrantType
	principalTypes  []string
}

func (r *TypeRule) defaultClaimsOf(appID string) []*protomodel.Claim {
	return []*protomodel.Claim{
		{
			Service: appID,
			Action:  "*",
			Scope:   []string{"channel-{id}"},
		},
	}
}

func (m *ClaimManager) typeRuleOf(t model.RoleType) TypeRule {
	return m.typeRule[t]
}
