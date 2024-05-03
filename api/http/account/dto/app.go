package dto

type AppModifyRequest struct {
	AppID             string         `json:"appID"`
	AccountID         string         `json:"accountID"`
	Title             string         `json:"title"`
	ManualURL         string         `json:"manualUrl,omitempty"`
	Description       *string        `json:"description,omitempty"`
	DetailImageURLs   []string       `json:"detailImageUrls,omitempty"`
	DetailDescription string         `json:"detailDescription,omitempty"`
	I18nMap           map[string]any `json:"i18nMap,omitempty"`
}
