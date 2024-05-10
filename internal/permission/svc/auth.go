package svc

import (
	"context"

	"github.com/channel-io/ch-app-store/internal/approle/model"
	"github.com/channel-io/ch-app-store/internal/approle/svc"
)

type AccountAuthPermissionSvc struct {
	appAccountRepo AppAccountRepo
	delegate       *svc.AppRoleSvc
	tokenSvc       *svc.TokenSvc
}

func NewAccountAuthPermissionSvc(
	appAccountRepo AppAccountRepo,
	delegate *svc.AppRoleSvc,
	tokenSvc *svc.TokenSvc,
) *AccountAuthPermissionSvc {
	return &AccountAuthPermissionSvc{appAccountRepo: appAccountRepo, delegate: delegate, tokenSvc: tokenSvc}
}

func (s *AccountAuthPermissionSvc) UpdateRole(ctx context.Context, appID string, roleType model.RoleType, claims []*model.Claim, accountID string) ([]*model.Claim, error) {
	if _, err := s.appAccountRepo.Fetch(ctx, appID, accountID); err != nil {
		return nil, err
	}

	return s.delegate.UpdateRole(ctx, appID, roleType, claims)
}

func (s *AccountAuthPermissionSvc) FetchRole(ctx context.Context, appID string, roleType model.RoleType, accountID string) ([]*model.Claim, error) {
	if _, err := s.appAccountRepo.Fetch(ctx, appID, accountID); err != nil {
		return nil, err
	}

	return s.delegate.FetchRole(ctx, appID, roleType)
}

func (s *AccountAuthPermissionSvc) GetAvailableClaims(ctx context.Context, roleType model.RoleType) ([]*model.Claim, error) {
	return s.delegate.GetAvailableClaims(ctx, roleType)
}

func (s *AccountAuthPermissionSvc) HasTokenIssuedBefore(ctx context.Context, appID string, accountID string) (bool, error) {
	if _, err := s.appAccountRepo.Fetch(ctx, appID, accountID); err != nil {
		return false, err
	}
	return s.tokenSvc.HasIssuedBefore(ctx, appID)
}

func (s *AccountAuthPermissionSvc) RefreshToken(ctx context.Context, appID string, accountID string) (string, error) {
	if _, err := s.appAccountRepo.Fetch(ctx, appID, accountID); err != nil {
		return "", err
	}
	return s.tokenSvc.RefreshAppSecret(ctx, appID)
}
