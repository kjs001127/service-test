package domain

import (
	"context"

	"github.com/channel-io/go-lib/pkg/uid"

	app "github.com/channel-io/ch-app-store/internal/app/domain"
	"github.com/channel-io/ch-app-store/lib/db/tx"
)

type AppRole struct {
	RoleCredentials
	RoleID string
	AppID  string
}

type RoleCredentials struct {
	ClientID string
	Secret   string
}

type Role struct {
	ID     string
	Claims []Claim
	Type   RoleType
}

type RoleWithCredential struct {
	*Role
	RoleCredentials
}

type Claim struct {
	Service string
	Action  string
	Scopes  []string
}

type RoleClient interface {
	ReadRole(ctx context.Context, roleID string) (*Role, error)
	CreateRole(ctx context.Context, request *Role) (RoleWithCredential, error)
	UpdateRole(ctx context.Context, roleID string, claims []Claim) (*Role, error)
	DeleteRole(ctx context.Context, roleID string) error
}

type AppRoleRepository interface {
	Save(ctx context.Context, role *AppRole) error
	FetchByAppID(ctx context.Context, appID string) ([]*AppRole, error)
	FetchByRoleID(ctx context.Context, roleID string) (*AppRole, error)
	DeleteByAppID(ctx context.Context, appID string) error
}

type RoleType string

type AppRequest struct {
	Roles []*Role `json:"roles"`
	*RemoteApp
}

type AppResponse struct {
	*RemoteApp
	Roles []RoleWithCredential
}

type AppDevSvc interface {
	CreateApp(ctx context.Context, req AppRequest) (AppResponse, error)
	FetchApp(ctx context.Context, appID string) (AppResponse, error)
	DeleteApp(ctx context.Context, appID string) error
	FetchAppByRoleID(ctx context.Context, clientID string) (*app.App, error)
}

type AppDevSvcImpl struct {
	roleCli  RoleClient
	roleRepo AppRoleRepository
	urlRepo  AppUrlRepository
	manager  app.AppManager
}

func NewAppDevSvcImpl(
	roleCli RoleClient,
	roleRepo AppRoleRepository,
	urlRepo AppUrlRepository,
	manager app.AppManager,
) *AppDevSvcImpl {
	return &AppDevSvcImpl{
		roleCli:  roleCli,
		roleRepo: roleRepo,
		urlRepo:  urlRepo,
		manager:  manager,
	}
}

func (s *AppDevSvcImpl) CreateApp(ctx context.Context, req AppRequest) (AppResponse, error) {
	return tx.RunWith(ctx, func(ctx context.Context) (AppResponse, error) {
		req.ID = uid.New().Hex()

		created, err := s.manager.Create(ctx, req.App)
		if err != nil {
			return AppResponse{}, err
		}

		if err := s.urlRepo.Save(ctx, req.App.ID, req.Urls); err != nil {
			return AppResponse{}, err
		}

		roles, err := s.createRoles(ctx, req)
		if err != nil {
			return AppResponse{}, err
		}

		return AppResponse{
			Roles: roles,
			RemoteApp: &RemoteApp{
				App:  created,
				Urls: req.Urls,
			},
		}, nil

	})
}

func (s *AppDevSvcImpl) createRoles(ctx context.Context, req AppRequest) ([]RoleWithCredential, error) {
	var roles []RoleWithCredential
	for _, r := range req.Roles {
		roleWithCredential, err := s.roleCli.CreateRole(ctx, r)
		if err != nil {
			return nil, err
		}
		if err := s.roleRepo.Save(ctx, &AppRole{
			AppID:           req.ID,
			RoleID:          roleWithCredential.ID,
			RoleCredentials: roleWithCredential.RoleCredentials,
		}); err != nil {
			return nil, err
		}

		roles = append(roles, roleWithCredential)
	}
	return roles, nil
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

func (s *AppDevSvcImpl) fetchRoles(ctx context.Context, roleCredentials []*AppRole) ([]RoleWithCredential, error) {
	var roles []RoleWithCredential
	for _, c := range roleCredentials {
		role, err := s.roleCli.ReadRole(ctx, c.RoleID)
		if err != nil {
			return nil, err
		}
		roles = append(roles, RoleWithCredential{Role: role, RoleCredentials: c.RoleCredentials})
	}
	return roles, nil
}

func (s *AppDevSvcImpl) DeleteApp(ctx context.Context, appID string) error {
	return tx.Run(ctx, func(ctx context.Context) error {
		if err := s.urlRepo.Delete(ctx, appID); err != nil {
			return err
		}

		if err := s.manager.Delete(ctx, appID); err != nil {
			return err
		}

		appRoles, err := s.roleRepo.FetchByAppID(ctx, appID)
		if err != nil {
			return err
		}

		for _, c := range appRoles {
			if err := s.roleCli.DeleteRole(ctx, c.RoleID); err != nil {
				return err
			}
		}

		return s.roleRepo.DeleteByAppID(ctx, appID)
	})
}
