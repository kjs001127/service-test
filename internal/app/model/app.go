package model

type App struct {
	ID    string   `json:"id"`
	State AppState `json:"state"`

	Title       string  `json:"title"`
	AvatarURL   *string `json:"avatarUrl,omitempty"`
	Description *string `json:"description,omitempty"`

	I18nMap map[string]I18nFields `json:"i18NMap,omitempty"`

	IsBuiltIn bool `json:"isBuiltIn"`
}

type I18nFields struct {
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
}

type AppState string

const (
	AppStateEnabled  = AppState("enabled")
	AppStateDisabled = AppState("disabled")
)

type AppType string
