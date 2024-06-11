package dto

import (
	displaysvc "github.com/channel-io/ch-app-store/internal/appdisplay/svc"
)

type AppView struct {
	ID          string                 `json:"id"`
	Title       string                 `json:"title"`
	Description *string                `json:"description,omitempty"`
	AvatarUrl   *string                `json:"avatarUrl,omitempty"`
	IsPrivate   bool                   `json:"isPrivate"`
	I18nMap     map[string]AppViewI18n `json:"i18nMap"`
}

func AppWithDisplayViewFrom(app *displaysvc.AppWithDisplay) *AppView {
	return &AppView{
		ID:          app.ID,
		Title:       app.Title,
		Description: app.Description,
		AvatarUrl:   app.AvatarURL,
		I18nMap:     convertAppWithDisplayViewI18n(app),
		IsPrivate:   app.IsPrivate,
	}
}

func AppWithDisplayViewsFrom(apps []*displaysvc.AppWithDisplay) []*AppView {
	ret := make([]*AppView, 0, len(apps))
	for _, app := range apps {
		ret = append(ret, AppWithDisplayViewFrom(app))
	}
	return ret
}

func convertAppWithDisplayViewI18n(app *displaysvc.AppWithDisplay) map[string]AppViewI18n {
	ret := make(map[string]AppViewI18n)
	for lang, i18n := range app.I18nMap {
		ret[lang] = AppViewI18n{
			Title:       i18n.Title,
			Description: i18n.Description,
		}
	}
	return ret
}

type AppViewI18n struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}

type AppGeneral struct {
	ID                 string                    `json:"id"`
	Title              string                    `json:"title"`
	Description        *string                   `json:"description,omitempty"`
	DetailDescriptions []map[string]any          `json:"detailDescriptions,omitempty"`
	DetailImageURLs    []string                  `json:"detailImageUrls,omitempty"`
	I18nMap            map[string]AppGeneralI18n `json:"i18nMap,omitempty"`
	AvatarURL          *string                   `json:"avatarUrl,omitempty"`
	ManualURL          *string                   `json:"manualUrl"`
	IsPrivate          bool                      `json:"isPrivate"`
}

type AppGeneralI18n struct {
	Title              string           `json:"title"`
	Description        string           `json:"description"`
	DetailDescriptions []map[string]any `json:"detailDescriptions,omitempty"`
	DetailImageURLs    []string         `json:"detailImageUrls,omitempty"`
	ManualURL          string           `json:"manualUrl,omitempty"`
}

func convertAppWithDisplayI18n(app *displaysvc.AppWithDisplay) map[string]AppGeneralI18n {
	ret := make(map[string]AppGeneralI18n)
	for lang, i18n := range app.I18nMap {
		ret[lang] = AppGeneralI18n{
			Title:              i18n.Title,
			Description:        i18n.Description,
			DetailDescriptions: i18n.DetailDescriptions,
			DetailImageURLs:    i18n.DetailImageURLs,
			ManualURL:          i18n.ManualURL,
		}
	}
	return ret
}

func FromAppWithDisplay(model *displaysvc.AppWithDisplay) *AppGeneral {
	return &AppGeneral{
		ID:                 model.ID,
		Title:              model.Title,
		Description:        model.Description,
		I18nMap:            convertAppWithDisplayI18n(model),
		DetailDescriptions: model.DetailDescriptions,
		DetailImageURLs:    model.DetailImageURLs,
		ManualURL:          model.ManualURL,
		IsPrivate:          model.IsPrivate,
		AvatarURL:          model.AvatarURL,
	}
}

func FromAppsWithDisplay(models []*displaysvc.AppWithDisplay) []*AppGeneral {
	ret := make([]*AppGeneral, 0, len(models))
	for _, m := range models {
		ret = append(ret, FromAppWithDisplay(m))
	}
	return ret
}
