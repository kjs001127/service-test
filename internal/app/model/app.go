package model

type App struct {
	ID string `json:"id"`

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

type AppType string
