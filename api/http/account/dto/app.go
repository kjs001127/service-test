package dto

import (
	appmodel "github.com/channel-io/ch-app-store/internal/app/model"
)

type AppView struct {
	Title       string                 `json:"title"`
	Description *string                `json:"description"`
	AvatarUrl   *string                `json:"avatarUrl"`
	I18nMap     map[string]AppViewI18n `json:"i18nMap"`
}

func AppViewFrom(app *appmodel.App) *AppView {
	return &AppView{
		Title:       app.Title,
		Description: app.Description,
		AvatarUrl:   app.AvatarURL,
		I18nMap:     convertAppViewI18n(app),
	}
}

func AppViewsFrom(apps []*appmodel.App) []*AppView {
	ret := make([]*AppView, 0, len(apps))
	for _, app := range apps {
		ret = append(ret, AppViewFrom(app))
	}
	return ret
}

func convertAppViewI18n(app *appmodel.App) map[string]AppViewI18n {
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
	Title       string `json:"title"`
	Description string `json:"description"`
}

type AppModifyRequest struct {
	Title              string                `json:"title"`
	Description        *string               `json:"description,omitempty"`
	DetailImageURLs    []string              `json:"detailImageUrls,omitempty"`
	DetailDescriptions []map[string]any      `json:"detailDescriptions,omitempty"`
	AvatarUrl          *string               `json:"avatarUrl,omitempty"`
	I18nMap            map[string]I18nFields `json:"i18nMap,omitempty"`
}

type I18nFields struct {
	Title             string           `json:"title"`
	DetailImageURLs   []string         `json:"detailImageUrls,omitempty"`
	DetailDescription []map[string]any `json:"detailDescription,omitempty"`
	Description       string           `json:"description,omitempty"`
}

func (r *AppModifyRequest) convertI18nMap() map[string]appmodel.I18nFields {
	ret := make(map[string]appmodel.I18nFields)
	for lang, i18n := range r.I18nMap {
		ret[lang] = appmodel.I18nFields{
			Title:             i18n.Title,
			DetailImageURLs:   i18n.DetailImageURLs,
			DetailDescription: i18n.DetailDescription,
			Description:       i18n.Description,
		}
	}
	return ret
}

func (r *AppModifyRequest) ConvertToApp(appID string) *appmodel.App {
	return &appmodel.App{
		ID:                 appID,
		Title:              r.Title,
		Description:        r.Description,
		DetailImageURLs:    r.DetailImageURLs,
		DetailDescriptions: r.DetailDescriptions,
		I18nMap:            r.convertI18nMap(),
		AvatarURL:          r.AvatarUrl,
	}
}

type AppResponse struct {
	ID                 string                `json:"id"`
	Title              string                `json:"title"`
	Description        *string               `json:"description,omitempty"`
	DetailDescriptions []map[string]any      `json:"detailDescriptions,omitempty"`
	I18nMap            map[string]I18nFields `json:"i18nMap,omitempty"`
	AvatarURL          *string               `json:"avatarUrl,omitempty"`
	IsPrivate          bool                  `json:"isPrivate"`
}

func FromApp(model *appmodel.App) *AppResponse {
	return &AppResponse{
		ID:                 model.ID,
		Title:              model.Title,
		Description:        model.Description,
		I18nMap:            convertI18nMap(model.I18nMap),
		DetailDescriptions: model.DetailDescriptions,
		IsPrivate:          model.IsPrivate,
		AvatarURL:          model.AvatarURL,
	}
}

func FromApps(models []*appmodel.App) []*AppResponse {
	ret := make([]*AppResponse, 0, len(models))
	for _, m := range models {
		ret = append(ret, FromApp(m))
	}
	return ret
}

func convertI18nMap(fields map[string]appmodel.I18nFields) map[string]I18nFields {
	ret := make(map[string]I18nFields)
	for lang, i18n := range fields {
		ret[lang] = I18nFields{
			Title:             i18n.Title,
			DetailImageURLs:   i18n.DetailImageURLs,
			DetailDescription: i18n.DetailDescription,
			Description:       i18n.Description,
		}
	}
	return ret
}
