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
	serverSettingRepo function.AppServerSettingRepository
	roleSvc           *role.AppRoleSvc
	manager           app.AppCrudSvc
}

func NewAppDevSvcImpl(
	serverSettingRepo function.AppServerSettingRepository,
	roleSvc *role.AppRoleSvc,
	manager app.AppCrudSvc,
) *AppDevSvcImpl {
	return &AppDevSvcImpl{serverSettingRepo: serverSettingRepo, roleSvc: roleSvc, manager: manager}
}

func (s *AppDevSvcImpl) FetchApp(ctx context.Context, appID string) (AppResponse, error) {
	return tx.DoReturn(ctx, func(ctx context.Context) (AppResponse, error) {
		serverSetting, err := s.serverSettingRepo.Fetch(ctx, appID)
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
				App:           found,
				ServerSetting: serverSetting,
			},
		}, nil
	}, tx.ReadOnly())
}

func (s *AppDevSvcImpl) CreateApp(ctx context.Context, req AppRequest) (AppResponse, error) {
	created, err := tx.DoReturn(ctx, func(ctx context.Context) (*RemoteApp, error) {
		created, err := s.manager.Create(ctx, req.App)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		if err = s.serverSettingRepo.Save(ctx, req.App.ID, req.ServerSetting); err != nil {
			return nil, errors.WithStack(err)
		}
		return &RemoteApp{App: created, ServerSetting: req.ServerSetting}, nil
	})
	if err != nil {
		return AppResponse{}, err
	}

	roles, err := s.roleSvc.CreateRoles(ctx, req.ID, req.Roles)
	if err != nil {
		return AppResponse{}, err
	}

	return AppResponse{Roles: roles, RemoteApp: created}, nil

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
	functionmodel.ServerSetting
}
