package svc

import (
	"context"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
	"github.com/pkg/errors"

	app "github.com/channel-io/ch-app-store/internal/app/svc"
	"github.com/channel-io/ch-app-store/internal/approle/model"
	authgen "github.com/channel-io/ch-app-store/internal/auth/general"
	protomodel "github.com/channel-io/ch-proto/auth/v1/go/model"
	"github.com/channel-io/ch-proto/auth/v1/go/service"
)

type AppRoleSvc struct {
	roleCli  authgen.RoleFetcher
	roleRepo AppRoleRepository
	typeRule map[model.RoleType]TypeRule
	appSvc   app.AppCrudSvc
}

type RoleDTO struct {
	Credentials *model.Credentials
	Claims      []*model.Claim
}

func (s *AppRoleSvc) CreateRoles(ctx context.Context, appID string) error {
	for _, roleType := range model.AvailableRoleTypes {
		typeRule := s.typeRule[roleType]
		res, err := s.roleCli.CreateRole(ctx, &service.CreateRoleRequest{
			Claims:                typeRule.defaultClaimsOf(appID),
			AllowedPrincipalTypes: typeRule.principalTypes,
			AllowedGrantTypes:     typeRule.grantTypes,
		})
		if err != nil {
			return err
		}

		if err := s.roleRepo.Save(ctx, &model.AppRole{
			AppID:  appID,
			Type:   roleType,
			RoleID: res.Role.Id,
			Credentials: &model.Credentials{
				ClientSecret: res.Credentials.ClientSecret,
				ClientID:     res.Credentials.ClientId,
			},
		}); err != nil {
			return err
		}
	}
	return nil
}

func (s *AppRoleSvc) UpdateRole(ctx context.Context, appID string, roleType model.RoleType, claims []*model.Claim) (*RoleDTO, error) {

	appRole, err := s.roleRepo.FetchRoleByAppIDAndType(ctx, appID, roleType)
	if err != nil {
		return nil, err
	}

	protoClaims, err := s.toProtoClaims(ctx, roleType, claims)
	if err != nil {
		return nil, err
	}

	res, err := s.roleCli.UpdateRole(ctx, &service.ReplaceRoleClaimsRequest{
		RoleId: appRole.RoleID,
		Claims: protoClaims,
	})

	return &RoleDTO{
		Claims:      s.fromProtoClaims(res.Role.Claims),
		Credentials: appRole.Credentials,
	}, nil
}

func (s *AppRoleSvc) FetchRole(ctx context.Context, appID string, roleType model.RoleType) (*RoleDTO, error) {
	appRole, err := s.roleRepo.FetchRoleByAppIDAndType(ctx, appID, roleType)
	if err != nil {
		return nil, err
	}

	role, err := s.roleCli.GetRole(ctx, appRole.RoleID)
	if err != nil {
		return nil, err
	}

	return &RoleDTO{
		Claims:      s.fromProtoClaims(role.Role.Claims),
		Credentials: appRole.Credentials,
	}, nil
}

func (s *AppRoleSvc) DeleteRoles(ctx context.Context, appID string) error {
	roles, err := s.roleRepo.FetchByAppID(ctx, appID)
	if apierr.IsNotFound(err) {
		return nil
	} else if err != nil {
		return err
	}

	for _, role := range roles {
		if _, err := s.roleCli.DeleteRole(ctx, role.RoleID); err != nil {
			return err
		}
	}

	return s.roleRepo.DeleteByAppID(ctx, appID)
}

func (s *AppRoleSvc) fromProtoClaims(claims []*protomodel.Claim) []*model.Claim {
	ret := make([]*model.Claim, 0, len(claims))
	for _, claim := range claims {
		ret = append(ret, s.fromProtoClaim(claim))
	}
	return ret
}

func (s *AppRoleSvc) fromProtoClaim(claim *protomodel.Claim) *model.Claim {
	return &model.Claim{
		Service: claim.Service,
		Action:  claim.Action,
	}
}

func (s *AppRoleSvc) toProtoClaims(ctx context.Context, roleType model.RoleType, claims []*model.Claim) ([]*protomodel.Claim, error) {
	ret := make([]*protomodel.Claim, 0, len(claims))
	for _, claim := range claims {
		res, err := s.toProtoClaim(ctx, roleType, claim)
		if err != nil {
			return nil, err
		}
		ret = append(ret, res)
	}
	return ret, nil
}

func (s *AppRoleSvc) toProtoClaim(ctx context.Context, roleType model.RoleType, claim *model.Claim) (*protomodel.Claim, error) {
	protoClaim, exists := s.findMatchingClaim(roleType, claim)
	if exists {
		return protoClaim, nil
	}

	if roleType == model.RoleTypeChannel && !exists {
		appFound, err := s.appSvc.Read(ctx, claim.Service)
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

func (s *AppRoleSvc) findMatchingClaim(t model.RoleType, claim *model.Claim) (*protomodel.Claim, bool) {
	claims := s.typeRule[t].availableClaims
	for _, protoClaim := range claims {
		if protoClaim.Service == claim.Service && protoClaim.Action == claim.Action {
			return protoClaim, true
		}
	}
	return nil, false
}

func (s *AppRoleSvc) GetAvailableClaims(ctx context.Context, roleType model.RoleType) ([]*model.Claim, error) {
	claims := s.typeRule[roleType].availableClaims
	ret := make([]*model.Claim, 0, len(claims))
	for _, claim := range claims {
		ret = append(ret, &model.Claim{
			Service: claim.Service,
			Action:  claim.Action,
		})
	}
	return ret, nil
}
