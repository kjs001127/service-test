package svc

import (
	"context"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
	"github.com/pkg/errors"

	"github.com/channel-io/ch-app-store/internal/app/model"
	appsvc "github.com/channel-io/ch-app-store/internal/app/svc"
	role "github.com/channel-io/ch-app-store/internal/role/model"
	"github.com/channel-io/ch-app-store/internal/role/svc"
)

type AccountAuthPermissionSvc struct {
	appAccountRepo AppAccountRepo
	appQuerySvc    appsvc.AppQuerySvc
	roleSvc        *svc.AppRoleSvc
	secretSvc      *svc.AppSecretSvc
}

func NewAccountAuthPermissionSvc(
	appAccountRepo AppAccountRepo,
	roleSvc *svc.AppRoleSvc,
	secretSvc *svc.AppSecretSvc,
	appQuerySvc appsvc.AppQuerySvc,
) *AccountAuthPermissionSvc {
	return &AccountAuthPermissionSvc{
		appAccountRepo: appAccountRepo,
		roleSvc:        roleSvc,
		secretSvc:      secretSvc,
		appQuerySvc:    appQuerySvc,
	}
}

func (s *AccountAuthPermissionSvc) CreateRole(ctx context.Context, claims *svc.ClaimsRequest, accountID string) (*svc.ClaimsResponse, error) {
	if _, err := s.appAccountRepo.Fetch(ctx, claims.AppID, accountID); err != nil {
		return nil, err
	}

	if err := s.hasPermission(ctx, accountID, claims.AppClaims); err != nil {
		return nil, err
	}

	return s.roleSvc.CreateRole(ctx, claims)
}

func (s *AccountAuthPermissionSvc) FetchLatestRole(ctx context.Context, appID string, roleType role.RoleType, accountID string) (*svc.ClaimsResponse, error) {
	if _, err := s.appAccountRepo.Fetch(ctx, appID, accountID); err != nil {
		return nil, err
	}

	return s.roleSvc.FetchLatestRole(ctx, appID, roleType)
}

func (s *AccountAuthPermissionSvc) GetAvailableApps(ctx context.Context, accountID string) ([]*model.App, error) {
	ownerships, err := s.appAccountRepo.FetchAllByAccountID(ctx, accountID)
	if err != nil {
		return nil, err
	}

	var appIDs []string
	for _, ownership := range ownerships {
		appIDs = append(appIDs, ownership.AppID)
	}

	publics, err := s.appQuerySvc.ListPublicApps(ctx, "", 500)
	if err != nil {
		return nil, err
	}

	ownApps, err := s.appQuerySvc.ReadAllByAppIDs(ctx, appIDs)
	if err != nil {
		return nil, err
	}

	return append(ownApps, publics...), nil
}

func (s *AccountAuthPermissionSvc) GetAvailableNativeClaims(ctx context.Context, appID string, roleType role.RoleType) ([]*role.Claim, error) {
	return s.roleSvc.FetchAvailableNativeClaims(ctx, appID, roleType)
}

func (s *AccountAuthPermissionSvc) HasTokenIssuedBefore(ctx context.Context, appID string, accountID string) (bool, error) {
	if _, err := s.appAccountRepo.Fetch(ctx, appID, accountID); err != nil {
		return false, err
	}
	return s.secretSvc.HasIssuedBefore(ctx, appID)
}

func (s *AccountAuthPermissionSvc) RefreshToken(ctx context.Context, appID string, accountID string) (string, error) {
	if _, err := s.appAccountRepo.Fetch(ctx, appID, accountID); err != nil {
		return "", err
	}
	return s.secretSvc.RefreshAppSecret(ctx, appID)
}

func (s *AccountAuthPermissionSvc) hasPermission(ctx context.Context, accountID string, claims role.Claims) error {
	callables, err := s.GetAvailableApps(ctx, accountID)
	if err != nil {
		return err
	}

	for _, claim := range claims {
		if !contains(claim, callables) {
			return apierr.Unauthorized(errors.New("claim rejected"))
		}
	}
	return nil
}

func contains(claim *role.Claim, apps []*model.App) bool {
	for _, callable := range apps {
		if callable.ID == claim.Service {
			return true
		}
	}
	return false
}
