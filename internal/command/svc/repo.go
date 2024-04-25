package svc

import (
	"context"

	appmodel "github.com/channel-io/ch-app-store/internal/app/model"
	"github.com/channel-io/ch-app-store/internal/command/model"
)

type CommandRepository interface {
	FetchByAppIDsAndScope(ctx context.Context, appIDs []string, scope model.Scope) ([]*model.Command, error)
	Fetch(ctx context.Context, key model.CommandKey) (*model.Command, error)

	FetchAllByAppIDs(ctx context.Context, appIDs []string) ([]*model.Command, error)
	FetchAllByAppID(ctx context.Context, appID string) ([]*model.Command, error)

	Delete(ctx context.Context, key model.CommandKey) error
	DeleteAllByAppID(ctx context.Context, appID string) error
	Save(ctx context.Context, resource *model.Command) (*model.Command, error)
}

type ActivationRepository interface {
	Save(ctx context.Context, activation *model.Activation) error
	Fetch(ctx context.Context, key appmodel.InstallationID) (*model.Activation, error)
	FetchAllByAppIDs(ctx context.Context, channelID string, appIDs []string) ([]*model.Activation, error)
	Delete(ctx context.Context, key appmodel.InstallationID) error
}

type ActivationSettingRepository interface {
	Fetch(ctx context.Context, appID string) (*model.ActivationSetting, error)
	Delete(ctx context.Context, appID string) error
	Save(ctx context.Context, activation *model.ActivationSetting) error
}
