package dto

import (
	appsvc "github.com/channel-io/ch-app-store/internal/app/svc"
)

type DetailedAppView struct {
	ID                 string                   `json:"id"`
	Title              string                   `json:"title"`
	AvatarURL          *string                  `json:"avatarUrl,omitempty"`
	Description        *string                  `json:"description,omitempty"`
	ManualURL          *string                  `json:"manualUrl,omitempty"`
	DetailDescriptions []map[string]any         `json:"detailDescriptions,omitempty"`
	DetailImageURLs    []string                 `json:"detailImageUrls,omitempty"`
	I18nMap            map[string]AppDetailI18n `json:"i18nMap"`
	IsBuiltIn          bool                     `json:"isBuiltIn"`
	IsPrivate          bool                     `json:"isPrivate"`
}

type AppDetailI18n struct {
	Title              string           `json:"title"`
	Description        string           `json:"description"`
	DetailDescriptions []map[string]any `json:"detailDescriptions,omitempty"`
	DetailImageURLs    []string         `json:"detailImageUrls,omitempty"`
	ManualURL          *string          `json:"manualUrl,omitempty"`
}

func convertI18n(app *appsvc.AppDetail) map[string]AppDetailI18n {
	ret := make(map[string]AppDetailI18n)
	for lang, i18n := range app.I18nMap {
		ret[lang] = AppDetailI18n{
			Title:              i18n.Title,
			Description:        i18n.Description,
			DetailDescriptions: i18n.DetailDescriptions,
			DetailImageURLs:    i18n.DetailImageURLs,
			ManualURL:          i18n.ManualURL,
		}
	}
	return ret
}

func FromDetail(model *appsvc.AppDetail) *DetailedAppView {
	return &DetailedAppView{
		ID:                 model.ID,
		Title:              model.Title,
		Description:        model.Description,
		I18nMap:            convertI18n(model),
		DetailDescriptions: model.DetailDescriptions,
		DetailImageURLs:    model.DetailImageURLs,
		ManualURL:          model.ManualURL,
		IsPrivate:          model.IsPrivate,
		AvatarURL:          model.AvatarURL,
		IsBuiltIn:          model.IsBuiltIn,
	}
}
