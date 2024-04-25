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

	IsBuiltIn bool `json:"isBuiltIn"`
}

type AppState string

const (
	AppStateEnabled  = AppState("enabled")
	AppStateDisabled = AppState("disabled")
)

type AppType string
