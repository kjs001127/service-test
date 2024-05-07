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

	I18nMap map[string]I18nFields `json:"i18NMap,omitempty"`

	IsBuiltIn bool `json:"isBuiltIn"`
}

type I18nFields struct {
	Title             string   `json:"title"`
	DetailImageURLs   []string `json:"detailImageUrls,omitempty"`
	DetailDescription string   `json:"detailDescription,omitempty"`
	Description       string   `json:"description,omitempty"`
}

type AppState string

const (
	AppStateEnabled  = AppState("enabled")
	AppStateDisabled = AppState("disabled")
)

type AppType string
