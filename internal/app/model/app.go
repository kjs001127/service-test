package model

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
	AppStateStable   = AppState("stable")
	AppStateUnstable = AppState("unstable")
)

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
