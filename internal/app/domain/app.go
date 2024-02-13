package domain

import (
	"context"
)

type AppRepository interface {
	Save(ctx context.Context, app *App) (*App, error)
	FindApps(ctx context.Context, appIDs []string) ([]*App, error)
	FindApp(ctx context.Context, appID string) (*App, error)
	Index(ctx context.Context, since string, limit int) ([]*App, error)
	Delete(ctx context.Context, appID string) error
}

type ConfigSchemas []ConfigSchema

type ConfigSchema struct {
	Name       string
	Type       string
	Key        string
	Default    *string
	Help       *string
	Attributes map[string]any
}

func (s ConfigSchemas) DefaultConfig() ConfigMap {
	ret := make(ConfigMap)
	for _, schema := range s {
		if schema.Default != nil {
			continue
		}
		ret[schema.Name] = *schema.Default
	}
	return ret
}

type AppState string

const (
	AppStateStable   = AppState("stable")
	AppStateUnStable = AppState("unstable")
)

type App struct {
	ID    string
	State AppState

	Title       string
	AvatarURL   *string
	Description *string

	IsPrivate         bool
	ManualURL         *string
	DetailDescription map[string]any
	DetailImageURLs   *string

	ConfigSchemas ConfigSchemas
}

type ConfigMap map[string]string

type AppChannel struct {
	AppID     string
	ChannelID string
	Configs   ConfigMap
}

type Install struct {
	AppID     string
	ChannelID string
}

type AppChannelRepository interface {
	Fetch(ctx context.Context, identifier Install) (*AppChannel, error)
	FindAllByChannel(ctx context.Context, channelID string) ([]*AppChannel, error)
	Save(ctx context.Context, appChannel *AppChannel) (*AppChannel, error)
	Delete(ctx context.Context, identifier Install) error
	DeleteByAppID(ctx context.Context, appID string) error
}
