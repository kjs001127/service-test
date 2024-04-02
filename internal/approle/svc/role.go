package svc

import (
	"context"
	"fmt"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
	"github.com/pkg/errors"

	appsvc "github.com/channel-io/ch-app-store/internal/app/svc"
	"github.com/channel-io/ch-app-store/internal/approle/model"
	"github.com/channel-io/ch-app-store/internal/function/svc"
	protomodel "github.com/channel-io/ch-proto/auth/v1/go/model"
	"github.com/channel-io/ch-proto/auth/v1/go/service"
)

type AppRoleSvc struct {
	roleCli   RoleClient
	roleRepo  AppRoleRepository
	urlRepo   svc.AppUrlRepository
	manager   appsvc.AppCrudSvc
	typeRules map[model.RoleType]TypeRule
}

func NewAppRoleSvc(
	roleCli RoleClient,
	roleRepo AppRoleRepository,
	typeRules map[model.RoleType]TypeRule,
) *AppRoleSvc {
	return &AppRoleSvc{roleCli: roleCli, roleRepo: roleRepo, typeRules: typeRules}
}

func (s *AppRoleSvc) CreateRoles(ctx context.Context, appID string, req []*RoleWithType) ([]*RoleWithCredential, error) {
	var roles []*RoleWithCredential
	for _, r := range req {
		rules, ok := s.typeRules[r.Type]
		if !ok {
			return nil, apierr.NotFound(fmt.Errorf("no role type found %s", r.Type))
		}

		if err := checkScopes(r, rules); err != nil {
			return nil, err
		}

		r.Claims = append(r.Claims, defaultClaimsOf(appID, rules)...)

		res, err := s.roleCli.CreateRole(ctx, &service.CreateRoleRequest{
			Claims:                r.Claims,
			AllowedPrincipalTypes: rules.GrantedPrincipalTypes,
			AllowedGrantTypes:     rules.GrantTypes,
		})
		if err != nil {
			return nil, errors.Wrap(err, "create role fail")
		}

		if err := s.roleRepo.Save(ctx, &model.AppRole{
			AppID:           appID,
			RoleID:          res.Role.Id,
			RoleCredentials: res.Credentials,
			Type:            r.Type,
		}); err != nil {
			return nil, errors.WithStack(err)
		}

		roles = append(roles, &RoleWithCredential{
			RoleWithType:    &RoleWithType{Role: res.Role, Type: r.Type},
			RoleCredentials: res.Credentials,
		})
	}
	return roles, nil
}

const allActions = "*"

func defaultClaimsOf(appID string, rule TypeRule) []*protomodel.Claim {
	return []*protomodel.Claim{{
		Scope:   rule.GrantedScopes,
		Service: appID,
		Action:  allActions,
	}}
}

func checkScopes(role *RoleWithType, rule TypeRule) error {
	// mark granted scopes of input role type
	grantedScopes := make(map[string]bool)
	for _, grantedScope := range rule.GrantedScopes {
		grantedScopes[grantedScope] = true
	}

	for _, claim := range role.Claims {

		// mark input scopes of claim
		inputScopes := make(map[string]bool)
		for _, scope := range claim.Scope {
			inputScopes[scope] = true
		}

		// check if all granted scopes are in input scopes
		for _, grantedScope := range rule.GrantedScopes {
			if _, ok := inputScopes[grantedScope]; !ok {
				return fmt.Errorf("input scope must contain scope %s on role type %s", grantedScope, role.Type)
			}
		}

		// check if all input scopes are in granted scopes
		for _, inputScope := range claim.Scope {
			if _, ok := grantedScopes[inputScope]; !ok {
				return fmt.Errorf("input scope %s is not granted for role type %s", inputScope, role.Type)
			}
		}
	}
	return nil
}

func (s *AppRoleSvc) FetchRoles(ctx context.Context, appID string) ([]*RoleWithCredential, error) {
	appRoles, err := s.roleRepo.FetchByAppID(ctx, appID)
	if err != nil {
		return nil, err
	}

	var roles []*RoleWithCredential
	for _, c := range appRoles {
		role, err := s.roleCli.GetRole(ctx, c.RoleID)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		roles = append(roles, &RoleWithCredential{
			RoleWithType:    &RoleWithType{Role: role.Role, Type: c.Type},
			RoleCredentials: c.RoleCredentials,
		})
	}
	return roles, nil
}

func (s *AppRoleSvc) DeleteRoles(ctx context.Context, appID string) error {
	appRoles, err := s.roleRepo.FetchByAppID(ctx, appID)
	if err != nil {
		return errors.WithStack(err)
	}
	for _, c := range appRoles {
		if _, err = s.roleCli.DeleteRole(ctx, c.RoleID); err != nil {
			return errors.Wrap(err, "error while deleting roles")
		}
	}
	return nil
}

type TypeRule struct {
	GrantTypes            []protomodel.GrantType
	GrantedPrincipalTypes []string
	GrantedScopes         []string
}

type RoleWithType struct {
	*protomodel.Role
	Type model.RoleType `json:"type"`
}

type RoleWithCredential struct {
	*RoleWithType
	*protomodel.RoleCredentials
}
