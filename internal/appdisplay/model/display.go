package model

type AppDisplay struct {
	AppID string `json:"appId"`

	IsPrivate          bool             `json:"isPrivate"`
	ManualURL          *string          `json:"manualUrl,omitempty"`
	DetailDescriptions []map[string]any `json:"detailDescriptions,omitempty"`
	DetailImageURLs    []string         `json:"detailImageUrls,omitempty"`

	I18nMap map[string]I18nFields `json:"i18NMap,omitempty"`
}

type I18nFields struct {
	DetailImageURLs    []string         `json:"detailImageUrls,omitempty"`
	DetailDescriptions []map[string]any `json:"detailDescriptions,omitempty"`
	ManualURL          string           `json:"manualURL,omitempty"`
}
