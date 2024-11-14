package svc

import (
	"context"
	"crypto/rand"
	"encoding/base64"

	"github.com/channel-io/go-lib/pkg/errors/apierr"

	"github.com/channel-io/ch-app-store/internal/role/model"
)

type AppSecretSvc struct {
	secretRepo AppSecretRepository
}

func NewAppSecretSvc(secretRepo AppSecretRepository) *AppSecretSvc {
	return &AppSecretSvc{secretRepo: secretRepo}
}

func (s *AppSecretSvc) DeleteAppSecret(ctx context.Context, appID string) error {
	return s.secretRepo.Delete(ctx, appID)
}

func (s *AppSecretSvc) RefreshAppSecret(ctx context.Context, appID string) (string, error) {
	token, err := generateSecret()
	if err != nil {
		return "", err
	}

	if err := s.secretRepo.Save(ctx, &model.AppSecret{
		AppID:  appID,
		Secret: token,
	}); err != nil {
		return "", err
	}

	return token, nil
}

func (s *AppSecretSvc) HasIssuedBefore(ctx context.Context, appID string) (bool, error) {
	_, err := s.secretRepo.FetchByAppID(ctx, appID)
	if apierr.IsNotFound(err) {
		return false, nil
	}
	return true, nil
}

func generateSecret() (string, error) {
	randomBytes := make([]byte, 16)
	if _, err := rand.Read(randomBytes); err != nil {
		return "", err
	}
	secret := base64.URLEncoding.EncodeToString(randomBytes)
	return secret, nil
}
