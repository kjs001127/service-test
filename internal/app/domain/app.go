package domain

import (
	"context"
)

type App struct {
	ID    string   `json:"id"`
	State AppState `json:"state"`

	Title       string  `json:"title"`
	AvatarURL   *string `json:"avatarUrl,omitempty"`
	Description *string `json:"description,omitempty"`

	IsPrivate          bool             `json:"isPrivate"`
	ManualURL          *string          `json:"manualUrl,omitempty"`
	DetailDescriptions []map[string]any `json:"detailDescriptions,omitempty"`
	DetailImageURLs    []string         `json:"detailImageUrls,omitempty"`

	ConfigSchemas ConfigSchemas `json:"configSchemas,omitempty"`
	IsBuiltIn     bool          `json:"isBuiltIn"`
	Type          AppType       `json:"-"`
}

type AppState string

const (
	AppStateEnabled  = AppState("enabled")
	AppStateDisabled = AppState("disabled")
)

type ConfigMap map[string]string
type ConfigSchemas []ConfigSchema
type ConfigSchema struct {
	Name       string         `json:"name"`
	Type       string         `json:"type"`
	Key        string         `json:"key"`
	Default    *string        `json:"default,omitempty"`
	Help       *string        `json:"help,omitempty"`
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

type AppType string
type Typed[T any] struct {
	Type    AppType
	Handler T
}

func NewTyped[T any](t AppType, handler T) Typed[T] {
	return Typed[T]{Type: t, Handler: handler}
}

type AppRepository interface {
	Save(ctx context.Context, app *App) (*App, error)
	FindApps(ctx context.Context, appIDs []string) ([]*App, error)
	FindApp(ctx context.Context, appID string) (*App, error)
	FindBuiltInApps(ctx context.Context) ([]*App, error)
	FindPublicApps(ctx context.Context, since string, limit int) ([]*App, error)
	Delete(ctx context.Context, appID string) error
}

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
