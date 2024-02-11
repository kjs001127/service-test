package domain

import (
	"context"

	app "github.com/channel-io/ch-app-store/internal/app/domain"
)

type ConfigValidator struct {
}

func NewConfigValidator() *ConfigValidator {
	return &ConfigValidator{}
}

func (c ConfigValidator) ValidateConfig(
	ctx context.Context,
	app *app.App,
	channelID string,
	input app.ConfigMap,
) error {
	return nil
}
