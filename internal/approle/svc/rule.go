package svc

import (
	"context"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
	"github.com/pkg/errors"

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
	defaultClaims   []*protomodel.Claim
	grantTypes      []protomodel.GrantType
	principalTypes  []string
}

func (r *TypeRule) defaultClaimsOf(appID string) []*protomodel.Claim {

}

func (r TypeRule) grantTypesOf(appID string) []protomodel.GrantType {

}

func (r TypeRule) availableClaimsOf(appID string) []*protomodel.Claim {

}

func (r TypeRule) principalTypesOf(appID string) []string {

}

func (m *ClaimManager) typeRuleOf(t model.RoleType) TypeRule {
	return m.typeRule[t]
}

func (m *ClaimManager) convertAll(ctx context.Context, appID string, roleType model.RoleType, claims []*model.Claim) ([]*protomodel.Claim, error) {
	ret := make([]*protomodel.Claim, 0, len(claims))
	for _, claim := range claims {
		res, err := m.convert(ctx, appID, roleType, claim)
		if err != nil {
			return nil, err
		}
		ret = append(ret, res)
	}
	return ret, nil
}

func (m *ClaimManager) convert(ctx context.Context, appID string, roleType model.RoleType, claim *model.Claim) (*protomodel.Claim, error) {
	protoClaim, exists := m.findMatchingClaim(roleType, appID, claim)
	if exists {
		return protoClaim, nil
	}

	if roleType == model.RoleTypeChannel && !exists {
		appFound, err := m.appSvc.Read(ctx, claim.Service)
		if err != nil {
			return nil, apierr.NotFound(err)
		}

		return &protomodel.Claim{
			Service: appFound.ID,
			Action:  "*",
			Scope:   []string{"channel-{id}"},
		}, nil
	}

	return nil, apierr.NotFound(errors.New("cannot accept claim"))
}

func (m *ClaimManager) findMatchingClaim(t model.RoleType, appID string, claim *model.Claim) (*protomodel.Claim, bool) {
	claims := m.typeRule[t].availableClaimsOf(appID)
	for _, protoClaim := range claims {
		if protoClaim.Service == claim.Service && protoClaim.Action == claim.Action {
			return protoClaim, true
		}
	}
	return nil, false
}

func (m *ClaimManager) GetAvailableClaims(ctx context.Context, appID string, roleType model.RoleType) ([]*model.Claim, error) {
	claims := m.typeRule[roleType].availableClaimsOf(appID)
	ret := make([]*model.Claim, 0, len(claims))
	for _, claim := range claims {
		ret = append(ret, &model.Claim{
			Service: claim.Service,
			Action:  claim.Action,
		})
	}
	return ret, nil
}
