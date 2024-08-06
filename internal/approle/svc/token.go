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

type TokenSvc interface {
	IssueManagerToken(ctx context.Context, appID string, manager account.ManagerPrincipal) (general.IssueResponse, error)
	IssueUserToken(ctx context.Context, appID string, user session.UserPrincipal) (general.IssueResponse, error)
	IssueChannelToken(ctx context.Context, channelID string, appToken string) (general.IssueResponse, error)
	IssueAppToken(ctx context.Context, appToken string) (general.IssueResponse, error)
	RefreshToken(ctx context.Context, refreshToken string) (general.IssueResponse, error)
}

type TokenSvcImpl struct {
	rbacExchanger   *general.RBACExchanger
	installQuerySvc *app.InstalledAppQuerySvc
	tokenRepo       AppSecretRepository
	roleRepo        AppRoleRepository
}

func NewTokenSvc(
	rbacExchanger *general.RBACExchanger,
	installQuerySvc *app.InstalledAppQuerySvc,
	tokenRepo AppSecretRepository,
	roleRepo AppRoleRepository,
) *TokenSvcImpl {
	return &TokenSvcImpl{
		rbacExchanger:   rbacExchanger,
		installQuerySvc: installQuerySvc,
		tokenRepo:       tokenRepo,
		roleRepo:        roleRepo,
	}
}

func (s *TokenSvcImpl) IssueManagerToken(ctx context.Context, appID string, manager account.ManagerPrincipal) (general.IssueResponse, error) {
	appRole, err := s.roleRepo.FetchRoleByAppIDAndType(ctx, appID, model.RoleTypeManager)
	if err != nil {
		return general.IssueResponse{}, err
	}

	scopes := general.Scopes{"channel": {manager.ChannelID}, "manager": {manager.ID}}
	return s.rbacExchanger.ExchangeWithPrincipal(ctx, manager.Token, scopes, appRole.Credentials.ClientID)
}

func (s *TokenSvcImpl) IssueUserToken(ctx context.Context, appID string, user session.UserPrincipal) (general.IssueResponse, error) {
	appRole, err := s.roleRepo.FetchRoleByAppIDAndType(ctx, appID, model.RoleTypeUser)
	if err != nil {
		return general.IssueResponse{}, err
	}

	scopes := general.Scopes{"channel": {user.ChannelID}, "user": {user.ID}}
	return s.rbacExchanger.ExchangeWithPrincipal(ctx, user.Token, scopes, appRole.Credentials.ClientID)
}

func (s *TokenSvcImpl) IssueChannelToken(ctx context.Context, channelID string, appToken string) (general.IssueResponse, error) {
	token, err := s.tokenRepo.FetchBySecret(ctx, appToken)
	if err != nil {
		return general.IssueResponse{}, err
	}

	installed, err := s.installQuerySvc.CheckInstall(ctx, appmodel.InstallationID{
		AppID:     token.AppID,
		ChannelID: channelID,
	})
	if err != nil {
		return general.IssueResponse{}, err
	}
	if !installed {
		return general.IssueResponse{}, apierr.Unauthorized(err)
	}

	appRole, err := s.roleRepo.FetchRoleByAppIDAndType(ctx, token.AppID, model.RoleTypeChannel)
	if err != nil {
		return general.IssueResponse{}, err
	}

	scopes := general.Scopes{"channel": {channelID}, "app": {token.AppID}}
	return s.rbacExchanger.ExchangeWithClientSecret(ctx, appRole.Credentials.ClientID, appRole.Credentials.ClientSecret, scopes)
}

func (s *TokenSvcImpl) IssueAppToken(ctx context.Context, appToken string) (general.IssueResponse, error) {
	token, err := s.tokenRepo.FetchBySecret(ctx, appToken)
	if err != nil {
		return general.IssueResponse{}, err
	}
	appRole, err := s.roleRepo.FetchRoleByAppIDAndType(ctx, token.AppID, model.RoleTypeApp)
	if err != nil {
		return general.IssueResponse{}, err
	}
	scopes := general.Scopes{"app": {token.AppID}}
	return s.rbacExchanger.ExchangeWithClientSecret(ctx, appRole.Credentials.ClientID, appRole.Credentials.ClientSecret, scopes)
}

func (s *TokenSvcImpl) RefreshToken(ctx context.Context, refreshToken string) (general.IssueResponse, error) {
	return s.rbacExchanger.Refresh(ctx, refreshToken)
}
