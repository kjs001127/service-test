package svc

import (
	"context"
	"crypto/rand"
	"encoding/base64"

	"github.com/channel-io/go-lib/pkg/errors/apierr"

	appmodel "github.com/channel-io/ch-app-store/internal/app/model"
	app "github.com/channel-io/ch-app-store/internal/app/svc"
	"github.com/channel-io/ch-app-store/internal/approle/model"
	"github.com/channel-io/ch-app-store/internal/auth/general"
	"github.com/channel-io/ch-app-store/internal/auth/principal/account"
	"github.com/channel-io/ch-app-store/internal/auth/principal/session"
)

type TokenSvc struct {
	rbacExchanger   *general.RBACExchanger
	installQuerySvc *app.AppInstallQuerySvc
	tokenRepo       AppTokenRepository
	roleRepo        AppRoleRepository
}

func NewTokenSvc(
	rbacExchanger *general.RBACExchanger,
	installQuerySvc *app.AppInstallQuerySvc,
	tokenRepo AppTokenRepository,
	roleRepo AppRoleRepository,
) *TokenSvc {
	return &TokenSvc{
		rbacExchanger:   rbacExchanger,
		installQuerySvc: installQuerySvc,
		tokenRepo:       tokenRepo,
		roleRepo:        roleRepo,
	}
}

func (s *TokenSvc) DeleteAppToken(ctx context.Context, appID string) error {
	return s.tokenRepo.Delete(ctx, appID)
}

func (s *TokenSvc) RefreshAppToken(ctx context.Context, appID string) (string, error) {
	token, err := generateToken()
	if err != nil {
		return "", err
	}

	if err := s.tokenRepo.Save(ctx, &model.AppToken{
		AppID: appID,
		Token: token,
	}); err != nil {
		return "", err
	}

	return token, nil
}

func generateToken() (string, error) {
	randomBytes := make([]byte, 16)
	if _, err := rand.Read(randomBytes); err != nil {
		return "", err
	}
	secret := base64.URLEncoding.EncodeToString(randomBytes)
	return secret, nil
}

func (s *TokenSvc) HasIssuedBefore(ctx context.Context, appID string) (bool, error) {
	_, err := s.tokenRepo.FetchByAppID(ctx, appID)
	if apierr.IsNotFound(err) {
		return false, nil
	}
	return true, nil
}

func (s *TokenSvc) IssueManagerToken(ctx context.Context, appID string, manager account.ManagerPrincipal) (general.IssueResponse, error) {
	appRole, err := s.roleRepo.FetchRoleByAppIDAndType(ctx, appID, model.RoleTypeManager)
	if err != nil {
		return general.IssueResponse{}, err
	}

	scopes := general.Scopes{"channel": {manager.ChannelID}, "manager": {manager.ID}, "app": {appID}}
	return s.rbacExchanger.ExchangeWithPrincipal(ctx, manager.Token, scopes, appRole.Credentials.ClientID)
}

func (s *TokenSvc) IssueUserToken(ctx context.Context, appID string, user session.UserPrincipal) (general.IssueResponse, error) {
	appRole, err := s.roleRepo.FetchRoleByAppIDAndType(ctx, appID, model.RoleTypeUser)
	if err != nil {
		return general.IssueResponse{}, err
	}

	scopes := general.Scopes{"channel": {user.ChannelID}, "user": {user.ID}, "app": {appID}}
	return s.rbacExchanger.ExchangeWithPrincipal(ctx, user.Token, scopes, appRole.Credentials.ClientID)
}

func (s *TokenSvc) IssueChannelToken(ctx context.Context, installID appmodel.InstallationID) (general.IssueResponse, error) {
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

func (s *TokenSvc) IssueAppToken(ctx context.Context, appToken string) (general.IssueResponse, error) {
	token, err := s.tokenRepo.FetchByToken(ctx, appToken)
	appRole, err := s.roleRepo.FetchRoleByAppIDAndType(ctx, token.AppID, model.RoleTypeApp)
	if err != nil {
		return general.IssueResponse{}, err
	}
	scopes := general.Scopes{"app": {token.AppID}}
	return s.rbacExchanger.ExchangeWithClientSecret(ctx, appRole.Credentials.ClientID, appRole.Credentials.ClientSecret, scopes)
}
