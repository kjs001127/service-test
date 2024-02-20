package domain

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/channel-io/go-lib/pkg/errors/apierr"

	app "github.com/channel-io/ch-app-store/internal/app/domain"
	"github.com/channel-io/ch-app-store/lib/db/tx"
	"github.com/channel-io/ch-proto/auth/v1/go/model"
	"github.com/channel-io/ch-proto/auth/v1/go/service"
)

type AppRole struct {
	*model.RoleCredentials
	RoleID string
	Type   RoleType
	AppID  string
}

type RoleWithType struct {
	*model.Role
	Type RoleType `json:"type"`
}

type RoleWithCredential struct {
	*RoleWithType
	*model.RoleCredentials
}

type RoleClient interface {
	GetRole(ctx context.Context, roleID string) (*service.GetRoleResult, error)
	CreateRole(ctx context.Context, request *service.CreateRoleRequest) (*service.CreateRoleResult, error)
	UpdateRole(ctx context.Context, roleID string, request *service.ReplaceRoleClaimsRequest) (*service.ReplaceRoleClaimsResult, error)
	DeleteRole(ctx context.Context, roleID string) (*service.DeleteRoleResult, error)
}

type AppRoleRepository interface {
	Save(ctx context.Context, role *AppRole) error
	FetchByAppID(ctx context.Context, appID string) ([]*AppRole, error)
	FetchByRoleID(ctx context.Context, roleID string) (*AppRole, error)
	DeleteByAppID(ctx context.Context, appID string) error
}

type RoleType string

type AppRequest struct {
	Roles []*RoleWithType `json:"roles"`
	*RemoteApp
}

type AppResponse struct {
	*RemoteApp
	Roles []*RoleWithCredential
}

type TypeRule struct {
	GrantTypes            []model.GrantType
	GrantedPrincipalTypes []string
	GrantedScopes         []string
}

type AppDevSvc interface {
	CreateApp(ctx context.Context, req AppRequest) (AppResponse, error)
	FetchApp(ctx context.Context, appID string) (AppResponse, error)
	DeleteApp(ctx context.Context, appID string) error
	FetchAppByRoleID(ctx context.Context, clientID string) (*app.App, error)
}

type AppDevSvcImpl struct {
	roleCli   RoleClient
	roleRepo  AppRoleRepository
	urlRepo   AppUrlRepository
	manager   app.AppManager
	typeRules map[RoleType]TypeRule
}

func NewAppDevSvcImpl(
	roleCli RoleClient,
	roleRepo AppRoleRepository,
	urlRepo AppUrlRepository,
	manager app.AppManager,
	typeRules map[RoleType]TypeRule,
) *AppDevSvcImpl {
	return &AppDevSvcImpl{roleCli: roleCli, roleRepo: roleRepo, urlRepo: urlRepo, manager: manager, typeRules: typeRules}
}

func (s *AppDevSvcImpl) CreateApp(ctx context.Context, req AppRequest) (AppResponse, error) {
	remoteApp, err := s.createRemoteApp(ctx, req)
	if err != nil {
		return AppResponse{}, err
	}

	roles, err := s.createRoles(ctx, req)
	if err != nil {
		return AppResponse{}, err
	}

	return AppResponse{
		Roles:     roles,
		RemoteApp: remoteApp,
	}, nil
}

func (s *AppDevSvcImpl) createRemoteApp(ctx context.Context, req AppRequest) (*RemoteApp, error) {
	return tx.RunWith(ctx, func(ctx context.Context) (*RemoteApp, error) {
		created, err := s.manager.Create(ctx, req.App)
		if err != nil {
			return nil, err
		}

		if err := s.urlRepo.Save(ctx, req.App.ID, req.Urls); err != nil {
			return nil, err
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

		if err := checkScopes(r.Role, rules); err != nil {
			return nil, err
		}

		r.Claims = append(r.Claims, &model.Claim{
			Scope:   []string{"channel"},
			Service: req.ID,
			Action:  "*",
		})

		res, err := s.roleCli.CreateRole(ctx, &service.CreateRoleRequest{
			Claims:                r.Claims,
			AllowedPrincipalTypes: rules.GrantedPrincipalTypes,
			AllowedGrantTypes:     rules.GrantTypes,
		})
		if err != nil {
			return nil, err
		}

		if err := s.roleRepo.Save(ctx, &AppRole{
			AppID:           req.ID,
			RoleID:          res.Role.Id,
			RoleCredentials: res.Credentials,
			Type:            r.Type,
		}); err != nil {
			return nil, err
		}

		roles = append(roles, &RoleWithCredential{
			RoleWithType:    &RoleWithType{Role: res.Role, Type: r.Type},
			RoleCredentials: res.Credentials,
		})
	}
	return roles, nil
}

func checkScopes(role *model.Role, rule TypeRule) error {
	for _, c := range role.Claims {
		for _, s := range c.Scope {
			scope, _, _ := strings.Cut(s, "-")

			var checked bool
			for _, grantedScope := range rule.GrantedScopes {
				if grantedScope == scope {
					checked = true
				}
			}

			if !checked {
				return errors.New("scope is not granted for type")
			}
		}
	}
	return nil
}

func (s *AppDevSvcImpl) FetchApp(ctx context.Context, appID string) (AppResponse, error) {
	return tx.RunWith(ctx, func(ctx context.Context) (AppResponse, error) {
		urls, err := s.urlRepo.Fetch(ctx, appID)
		if err != nil {
			return AppResponse{}, err
		}

		found, err := s.manager.Fetch(ctx, appID)
		if err != nil {
			return AppResponse{}, err
		}

		appRoles, err := s.roleRepo.FetchByAppID(ctx, appID)
		if err != nil {
			return AppResponse{}, err
		}

		roles, err := s.fetchRoles(ctx, appRoles)
		if err != nil {
			return AppResponse{}, err
		}

		return AppResponse{
			Roles: roles,
			RemoteApp: &RemoteApp{
				App:  found,
				Urls: urls,
			},
		}, nil
	}, tx.WithReadOnly())
}

func (s *AppDevSvcImpl) FetchAppByRoleID(ctx context.Context, clientID string) (*app.App, error) {
	appRole, err := s.roleRepo.FetchByRoleID(ctx, clientID)

	found, err := s.manager.Fetch(ctx, appRole.AppID)
	if err != nil {
		return nil, err
	}

	return found, nil
}

func (s *AppDevSvcImpl) fetchRoles(ctx context.Context, appRoles []*AppRole) ([]*RoleWithCredential, error) {
	var roles []*RoleWithCredential
	for _, c := range appRoles {
		role, err := s.roleCli.GetRole(ctx, c.RoleID)
		if err != nil {
			return nil, err
		}
		roles = append(roles, &RoleWithCredential{
			RoleWithType:    &RoleWithType{Role: role.Role, Type: c.Type},
			RoleCredentials: c.RoleCredentials,
		})
	}
	return roles, nil
}

func (s *AppDevSvcImpl) DeleteApp(ctx context.Context, appID string) error {
	appRoles, err := s.roleRepo.FetchByAppID(ctx, appID)
	if err != nil {
		return err
	}
	for _, c := range appRoles {
		if _, err := s.roleCli.DeleteRole(ctx, c.RoleID); err != nil {
			return err
		}
	}

	return tx.Run(ctx, func(ctx context.Context) error {
		if err := s.roleRepo.DeleteByAppID(ctx, appID); err != nil {
			return err
		}
		if err := s.urlRepo.Delete(ctx, appID); err != nil {
			return err
		}

		if err := s.manager.Delete(ctx, appID); err != nil {
			return err
		}
		return nil
	})
}
