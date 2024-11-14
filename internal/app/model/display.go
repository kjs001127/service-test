package model

type AppDisplay struct {
	AppID string `json:"appId"`

	ManualURL          *string          `json:"manualUrl,omitempty"`
	DetailDescriptions []map[string]any `json:"detailDescriptions,omitempty"`
	DetailImageURLs    []string         `json:"detailImageUrls,omitempty"`

	I18nMap map[string]DisplayI18n `json:"i18NMap,omitempty"`
}

type DisplayI18n struct {
	DetailImageURLs    []string         `json:"detailImageUrls,omitempty"`
	DetailDescriptions []map[string]any `json:"detailDescriptions,omitempty"`
	ManualURL          *string          `json:"manualURL,omitempty"`
}
