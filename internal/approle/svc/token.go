package svc

import (
	"context"

	"github.com/channel-io/go-lib/pkg/errors/apierr"

	appmodel "github.com/channel-io/ch-app-store/internal/app/model"
	app "github.com/channel-io/ch-app-store/internal/app/svc"
	"github.com/channel-io/ch-app-store/internal/approle/model"
	"github.com/channel-io/ch-app-store/internal/auth/general"
	"github.com/channel-io/ch-app-store/internal/auth/principal/account"
	"github.com/channel-io/ch-app-store/internal/auth/principal/session"
)

type TokenIssueSvc struct {
	rbacExchanger   *general.RBACExchanger
	installQuerySvc *app.AppInstallQuerySvc
	roleRepo        AppRoleRepository
}

func (s *TokenIssueSvc) IssueManagerToken(ctx context.Context, appID string, manager account.ManagerPrincipal) (general.IssueResponse, error) {
	appRole, err := s.roleRepo.FetchRoleByAppIDAndType(ctx, appID, model.RoleTypeManager)
	if err != nil {
		return general.IssueResponse{}, err
	}

	scopes := general.Scopes{"channel": {manager.ChannelID}, "manager": {manager.ID}, "app": {appID}}
	return s.rbacExchanger.ExchangeWithPrincipal(ctx, manager.Token, scopes, appRole.Credentials.ClientID)
}

func (s *TokenIssueSvc) IssueUserToken(ctx context.Context, appID string, user session.UserPrincipal) (general.IssueResponse, error) {
	appRole, err := s.roleRepo.FetchRoleByAppIDAndType(ctx, appID, model.RoleTypeUser)
	if err != nil {
		return general.IssueResponse{}, err
	}

	scopes := general.Scopes{"channel": {user.ChannelID}, "user": {user.ID}, "app": {appID}}
	return s.rbacExchanger.ExchangeWithPrincipal(ctx, user.Token, scopes, appRole.Credentials.ClientID)
}

func (s *TokenIssueSvc) IssueChannelToken(ctx context.Context, installID appmodel.InstallationID) (general.IssueResponse, error) {
	installed, err := s.installQuerySvc.CheckInstall(ctx, installID)
	if err != nil {
		return general.IssueResponse{}, err
	}
	if !installed {
		return general.IssueResponse{}, apierr.Unauthorized(err)
	}

	appRole, err := s.roleRepo.FetchRoleByAppIDAndType(ctx, installID.AppID, model.RoleTypeChannel)
	if err != nil {
		return general.IssueResponse{}, err
	}

	scopes := general.Scopes{"channel": {installID.ChannelID}, "app": {installID.AppID}}
	return s.rbacExchanger.ExchangeWithClientSecret(ctx, appRole.Credentials.ClientID, appRole.Credentials.ClientSecret, scopes)
}

func (s *TokenIssueSvc) IssueAppToken(ctx context.Context, appID string) (general.IssueResponse, error) {
	appRole, err := s.roleRepo.FetchRoleByAppIDAndType(ctx, appID, model.RoleTypeApp)
	if err != nil {
		return general.IssueResponse{}, err
	}
	scopes := general.Scopes{"app": {appID}}
	return s.rbacExchanger.ExchangeWithClientSecret(ctx, appRole.Credentials.ClientID, appRole.Credentials.ClientSecret, scopes)

}
