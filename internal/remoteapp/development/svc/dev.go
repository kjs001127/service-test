package svc

import (
	"context"
	"fmt"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
	"github.com/pkg/errors"

	appmodel "github.com/channel-io/ch-app-store/internal/app/model"
	appsvc "github.com/channel-io/ch-app-store/internal/app/svc"
	"github.com/channel-io/ch-app-store/internal/remoteapp/development/model"
	interactionmodel "github.com/channel-io/ch-app-store/internal/remoteapp/interaction/model"
	"github.com/channel-io/ch-app-store/internal/remoteapp/interaction/svc"
	"github.com/channel-io/ch-app-store/lib/db/tx"
	protomodel "github.com/channel-io/ch-proto/auth/v1/go/model"
	"github.com/channel-io/ch-proto/auth/v1/go/service"
)

type AppDevSvc interface {
	CreateApp(ctx context.Context, req AppRequest) (AppResponse, error)
	FetchApp(ctx context.Context, appID string) (AppResponse, error)
	DeleteApp(ctx context.Context, appID string) error
	FetchAppByRoleID(ctx context.Context, clientID string) (*appmodel.App, error)
}

type AppDevSvcImpl struct {
	roleCli   RoleClient
	roleRepo  AppRoleRepository
	urlRepo   svc.AppUrlRepository
	manager   appsvc.AppCrudSvc
	typeRules map[model.RoleType]TypeRule
}

func NewAppDevSvcImpl(
	roleCli RoleClient,
	roleRepo AppRoleRepository,
	urlRepo svc.AppUrlRepository,
	manager appsvc.AppCrudSvc,
	typeRules map[model.RoleType]TypeRule,
) *AppDevSvcImpl {
	return &AppDevSvcImpl{roleCli: roleCli, roleRepo: roleRepo, urlRepo: urlRepo, manager: manager, typeRules: typeRules}
}

func (s *AppDevSvcImpl) CreateApp(ctx context.Context, req AppRequest) (AppResponse, error) {
	remoteApp, err := s.createRemoteApp(ctx, req)
	if err != nil {
		return AppResponse{}, errors.WithStack(err)
	}

	roles, err := s.createRoles(ctx, req)
	if err != nil {
		return AppResponse{}, errors.WithStack(err)
	}

	return AppResponse{
		Roles:     roles,
		RemoteApp: remoteApp,
	}, nil
}

func (s *AppDevSvcImpl) createRemoteApp(ctx context.Context, req AppRequest) (*RemoteApp, error) {
	return tx.DoReturn(ctx, func(ctx context.Context) (*RemoteApp, error) {
		created, err := s.manager.Create(ctx, req.App)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		if err = s.urlRepo.Save(ctx, req.App.ID, req.Urls); err != nil {
			return nil, errors.WithStack(err)
		}
		return &RemoteApp{App: created, Urls: req.Urls}, nil
	})
}

func (s *AppDevSvcImpl) createRoles(ctx context.Context, req AppRequest) ([]*RoleWithCredential, error) {
	var roles []*RoleWithCredential
	for _, r := range req.Roles {
		rules, ok := s.typeRules[r.Type]
		if !ok {
			return nil, apierr.NotFound(fmt.Errorf("no role type found %s", r.Type))
		}

		if err := checkScopes(r, rules); err != nil {
			return nil, err
		}

		r.Claims = append(r.Claims, defaultClaimsOf(req, rules)...)

		res, err := s.roleCli.CreateRole(ctx, &service.CreateRoleRequest{
			Claims:                r.Claims,
			AllowedPrincipalTypes: rules.GrantedPrincipalTypes,
			AllowedGrantTypes:     rules.GrantTypes,
		})
		if err != nil {
			return nil, errors.Wrap(err, "create role fail")
		}

		if err := s.roleRepo.Save(ctx, &model.AppRole{
			AppID:           req.ID,
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

func defaultClaimsOf(appReq AppRequest, rule TypeRule) []*protomodel.Claim {
	return []*protomodel.Claim{{
		Scope:   rule.GrantedScopes,
		Service: appReq.ID,
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

func (s *AppDevSvcImpl) FetchApp(ctx context.Context, appID string) (AppResponse, error) {
	return tx.DoReturn(ctx, func(ctx context.Context) (AppResponse, error) {
		urls, err := s.urlRepo.Fetch(ctx, appID)
		if err != nil {
			return AppResponse{}, errors.WithStack(err)
		}

		found, err := s.manager.Read(ctx, appID)
		if err != nil {
			return AppResponse{}, errors.WithStack(err)
		}

		appRoles, err := s.roleRepo.FetchByAppID(ctx, appID)
		if err != nil {
			return AppResponse{}, errors.WithStack(err)
		}

		roles, err := s.fetchRoles(ctx, appRoles)
		if err != nil {
			return AppResponse{}, errors.WithStack(err)
		}

		return AppResponse{
			Roles: roles,
			RemoteApp: &RemoteApp{
				App:  found,
				Urls: urls,
			},
		}, nil
	}, tx.ReadOnly())
}

func (s *AppDevSvcImpl) fetchRoles(ctx context.Context, appRoles []*model.AppRole) ([]*RoleWithCredential, error) {
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

func (s *AppDevSvcImpl) FetchAppByRoleID(ctx context.Context, clientID string) (*appmodel.App, error) {
	appRole, err := s.roleRepo.FetchByRoleID(ctx, clientID)
	if err != nil {
		return nil, err
	}

	found, err := s.manager.Read(ctx, appRole.AppID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return found, nil
}

func (s *AppDevSvcImpl) DeleteApp(ctx context.Context, appID string) error {
	appRoles, err := s.roleRepo.FetchByAppID(ctx, appID)
	if err != nil {
		return errors.WithStack(err)
	}
	for _, c := range appRoles {
		if _, err = s.roleCli.DeleteRole(ctx, c.RoleID); err != nil {
			return errors.Wrap(err, "error while deleting roles")
		}
	}

	return tx.Do(ctx, func(ctx context.Context) error {
		if err = s.roleRepo.DeleteByAppID(ctx, appID); err != nil {
			return errors.WithStack(err)
		}
		if err = s.urlRepo.Delete(ctx, appID); err != nil {
			return errors.WithStack(err)
		}
		if err = s.manager.Delete(ctx, appID); err != nil {
			return errors.WithStack(err)
		}
		return nil
	})
}

type TypeRule struct {
	GrantTypes            []protomodel.GrantType
	GrantedPrincipalTypes []string
	GrantedScopes         []string
}

type AppRequest struct {
	Roles []*RoleWithType `json:"roles"`
	*RemoteApp
}

type AppResponse struct {
	*RemoteApp
	Roles []*RoleWithCredential `json:"roles,omitempty"`
}

type RemoteApp struct {
	*appmodel.App
	interactionmodel.Urls
}

type RoleWithType struct {
	*protomodel.Role
	Type model.RoleType `json:"type"`
}

type RoleWithCredential struct {
	*RoleWithType
	*protomodel.RoleCredentials
}
