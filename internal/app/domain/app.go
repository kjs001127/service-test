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
	Name       string         `json:"name"`
	Type       string         `json:"type"`
	Key        string         `json:"key"`
	Default    *string        `json:"default"`
	Help       *string        `json:"help"`
	Attributes map[string]any `json:"attributes"`
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
	ID    string   `json:"id"`
	State AppState `json:"state"`

	Title       string  `json:"title"`
	AvatarURL   *string `json:"avatarUrl"`
	Description *string `json:"description"`

	IsPrivate         bool           `json:"isPrivate"`
	ManualURL         *string        `json:"manualUrl"`
	DetailDescription map[string]any `json:"detailDescription"`
	DetailImageURLs   *string        `json:"detailImageUrls"`

	ConfigSchemas ConfigSchemas `json:"configSchemas"`
}

type ConfigMap map[string]string

type AppChannel struct {
	AppID     string    `json:"appId"`
	ChannelID string    `json:"channelId"`
	Configs   ConfigMap `json:"configs"`
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
