package svc

import (
	"context"

	"github.com/pkg/errors"

	appmodel "github.com/channel-io/ch-app-store/internal/app/model"
	app "github.com/channel-io/ch-app-store/internal/app/svc"
	functionmodel "github.com/channel-io/ch-app-store/internal/apphttp/model"
	function "github.com/channel-io/ch-app-store/internal/apphttp/svc"
	role "github.com/channel-io/ch-app-store/internal/approle/svc"
	"github.com/channel-io/ch-app-store/lib/db/tx"
)

type AppDevSvc interface {
	CreateApp(ctx context.Context, req AppRequest) (AppResponse, error)
	FetchApp(ctx context.Context, appID string) (AppResponse, error)
	DeleteApp(ctx context.Context, appID string) error
}

type AppDevSvcImpl struct {
	urlRepo function.AppUrlRepository
	roleSvc *role.AppRoleSvc
	manager app.AppCrudSvc
}

func NewAppDevSvcImpl(
	urlRepo function.AppUrlRepository,
	roleSvc *role.AppRoleSvc,
	manager app.AppCrudSvc,
) *AppDevSvcImpl {
	return &AppDevSvcImpl{urlRepo: urlRepo, roleSvc: roleSvc, manager: manager}
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

		roles, err := s.roleSvc.FetchRoles(ctx, appID)
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

func (s *AppDevSvcImpl) CreateApp(ctx context.Context, req AppRequest) (AppResponse, error) {
	roles, err := s.roleSvc.CreateRoles(ctx, req.ID, req.Roles)
	if err != nil {
		return AppResponse{}, err
	}

	return tx.DoReturn(ctx, func(ctx context.Context) (AppResponse, error) {
		created, err := s.manager.Create(ctx, req.App)
		if err != nil {
			return AppResponse{}, errors.WithStack(err)
		}

		if err = s.urlRepo.Save(ctx, req.App.ID, req.Urls); err != nil {
			return AppResponse{}, errors.WithStack(err)
		}
		return AppResponse{Roles: roles, RemoteApp: &RemoteApp{App: created, Urls: req.Urls}}, nil
	})
}

func (s *AppDevSvcImpl) DeleteApp(ctx context.Context, appID string) error {
	return s.manager.Delete(ctx, appID)
}

type AppRequest struct {
	Roles []*role.RoleWithType `json:"roles"`
	*RemoteApp
}

type AppResponse struct {
	*RemoteApp
	Roles []*role.RoleWithCredential `json:"roles,omitempty"`
}

type RemoteApp struct {
	*appmodel.App
	functionmodel.Urls
}
