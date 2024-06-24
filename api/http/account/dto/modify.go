package dto

import (
	appmodel "github.com/channel-io/ch-app-store/internal/app/model"
	displaymodel "github.com/channel-io/ch-app-store/internal/appdisplay/model"
	displaysvc "github.com/channel-io/ch-app-store/internal/appdisplay/svc"
	permissionsvc "github.com/channel-io/ch-app-store/internal/permission/svc"
)

type AppWithDisplayModifyRequest struct {
	Title              string                           `json:"title"`
	Description        *string                          `json:"description,omitempty"`
	DetailImageURLs    []string                         `json:"detailImageUrls,omitempty"`
	DetailDescriptions []map[string]any                 `json:"detailDescriptions,omitempty"`
	ManualURL          *string                          `json:"manualUrl,omitempty"`
	AvatarUrl          *string                          `json:"avatarUrl,omitempty"`
	I18nMap            map[string]displaysvc.I18nFields `json:"i18nMap,omitempty"`
}

func (a *AppWithDisplayModifyRequest) ToAppModifyRequest() *permissionsvc.AppModifyRequest {
	return &permissionsvc.AppModifyRequest{
		Title:       a.Title,
		Description: a.Description,
		AvatarUrl:   a.AvatarUrl,
		I18nMap:     convertToAppI18n(a.I18nMap),
	}
}

func (a *AppWithDisplayModifyRequest) ToDisplayModifyRequest() *permissionsvc.DisplayModifyRequest {
	return &permissionsvc.DisplayModifyRequest{
		DetailImageURLs:    a.DetailImageURLs,
		DetailDescriptions: a.DetailDescriptions,
		ManualURL:          a.ManualURL,
		I18nMap:            convertToDisplayI18n(a.I18nMap),
	}
}

func convertToDisplayI18n(i18nMap map[string]displaysvc.I18nFields) map[string]displaymodel.I18nFields {
	ret := make(map[string]displaymodel.I18nFields)
	for lang, i18n := range i18nMap {
		ret[lang] = displaymodel.I18nFields{
			DetailImageURLs:    i18n.DetailImageURLs,
			DetailDescriptions: i18n.DetailDescriptions,
			ManualURL:          i18n.ManualURL,
		}
	}
	return ret
}

func convertToAppI18n(i18nMap map[string]displaysvc.I18nFields) map[string]appmodel.I18nFields {
	ret := make(map[string]appmodel.I18nFields)
	for lang, i18n := range i18nMap {
		ret[lang] = appmodel.I18nFields{
			Title:       i18n.Title,
			Description: i18n.Description,
		}
	}
	return ret
}
