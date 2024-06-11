package svc

import (
	"context"

	display "github.com/channel-io/ch-app-store/internal/appdisplay/model"
	"github.com/channel-io/ch-app-store/internal/approle/model"
	"github.com/channel-io/ch-app-store/internal/approle/svc"

	"github.com/channel-io/go-lib/pkg/errors/apierr"

	"github.com/pkg/errors"
)

type AccountAuthPermissionSvc struct {
	appAccountRepo AppAccountRepo
	displaySvc     AccountDisplayPermissionSvc
	delegate       *svc.AppRoleSvc
}

func NewAccountAuthPermissionSvc(
	appAccountRepo AppAccountRepo,
	displaySvc AccountDisplayPermissionSvc,
	delegate *svc.AppRoleSvc,
) *AccountAuthPermissionSvc {
	return &AccountAuthPermissionSvc{appAccountRepo: appAccountRepo, delegate: delegate, displaySvc: displaySvc}
}

func (s *AccountAuthPermissionSvc) UpdateRole(ctx context.Context, appID string, roleType model.RoleType, claims *svc.ClaimsDTO, accountID string) error {
	if _, err := s.appAccountRepo.Fetch(ctx, appID, accountID); err != nil {
		return err
	}

	callables, err := s.displaySvc.GetCallableDisplays(ctx, accountID)
	if err != nil {
		return err
	}

	for _, appClaim := range claims.AppClaims {
		if !s.isCallable(callables, appClaim.Service) {
			return apierr.Unauthorized(errors.New("app is not callable"))
		}
	}

	return s.delegate.UpdateRole(ctx, appID, roleType, claims)
}

func (s *AccountAuthPermissionSvc) isCallable(callables []*display.AppDisplay, targetAppID string) bool {
	for _, callable := range callables {
		if callable.AppID == targetAppID {
			return true
		}
	}
	return false
}

func (s *AccountAuthPermissionSvc) FetchRole(ctx context.Context, appID string, roleType model.RoleType, accountID string) (*svc.ClaimsDTO, error) {
	if _, err := s.appAccountRepo.Fetch(ctx, appID, accountID); err != nil {
		return nil, err
	}

	return s.delegate.FetchRole(ctx, appID, roleType)
}

func (s *AccountAuthPermissionSvc) GetAvailableNativeClaims(ctx context.Context, roleType model.RoleType) ([]*model.Claim, error) {
	return s.delegate.GetAvailableNativeClaims(ctx, roleType)
}

func (s *AccountAuthPermissionSvc) HasTokenIssuedBefore(ctx context.Context, appID string, accountID string) (bool, error) {
	if _, err := s.appAccountRepo.Fetch(ctx, appID, accountID); err != nil {
		return false, err
	}
	return s.delegate.HasIssuedBefore(ctx, appID)
}

func (s *AccountAuthPermissionSvc) RefreshToken(ctx context.Context, appID string, accountID string) (string, error) {
	if _, err := s.appAccountRepo.Fetch(ctx, appID, accountID); err != nil {
		return "", err
	}
	return s.delegate.RefreshAppSecret(ctx, appID)
}
