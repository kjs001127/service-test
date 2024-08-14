package model

type AppWidget struct {
	ID    string `json:"id"`
	AppID string `json:"appId"`

	Name            string              `json:"name"`
	Description     *string             `json:"description,omitempty"`
	NameDescI18nMap map[string]*I18nMap `json:"nameDescI18nMap,omitempty"`

	DefaultName            *string             `json:"defaultName,omitempty"`
	DefaultDescription     *string             `json:"defaultDescription,omitempty"`
	DefaultNameDescI18nMap map[string]*I18nMap `json:"defaultNameDescI18nMap,omitempty"`

	ActionFunctionName string `json:"actionFunctionName"`
}

type I18nMap struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (a *AppWidget) Validate() error {
	return nil
}
