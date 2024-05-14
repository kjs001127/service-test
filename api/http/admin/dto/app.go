package dto

import (
	appmodel "github.com/channel-io/ch-app-store/internal/app/model"
)

type AdminAppModifyRequest struct {
	Title              string                     `json:"title"`
	Description        *string                    `json:"description,omitempty"`
	DetailImageURLs    []string                   `json:"detailImageUrls,omitempty"`
	DetailDescriptions []map[string]any           `json:"detailDescriptions,omitempty"`
	AvatarUrl          *string                    `json:"avatarUrl,omitempty"`
	I18nMap            map[string]AdminI18nFields `json:"i18nMap,omitempty"`
}

type AdminI18nFields struct {
	Title             string           `json:"title"`
	DetailImageURLs   []string         `json:"detailImageUrls,omitempty"`
	DetailDescription []map[string]any `json:"detailDescription,omitempty"`
	Description       string           `json:"description,omitempty"`
}

func (r *AdminAppModifyRequest) convertI18nMap() map[string]appmodel.I18nFields {
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

func (r *AdminAppModifyRequest) ConvertToApp(appID string) *appmodel.App {
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

type AdminAppResponse struct {
	ID                 string                     `json:"id"`
	Title              string                     `json:"title"`
	Description        *string                    `json:"description,omitempty"`
	DetailDescriptions []map[string]any           `json:"detailDescriptions,omitempty"`
	I18nMap            map[string]AdminI18nFields `json:"i18nMap,omitempty"`
	AvatarURL          *string                    `json:"avatarUrl,omitempty"`
	IsPrivate          bool                       `json:"isPrivate"`
}

func FromApp(model *appmodel.App) *AdminAppResponse {
	return &AdminAppResponse{
		ID:                 model.ID,
		Title:              model.Title,
		Description:        model.Description,
		I18nMap:            convertI18nMap(model.I18nMap),
		DetailDescriptions: model.DetailDescriptions,
		IsPrivate:          model.IsPrivate,
		AvatarURL:          model.AvatarURL,
	}
}

func convertI18nMap(fields map[string]appmodel.I18nFields) map[string]AdminI18nFields {
	ret := make(map[string]AdminI18nFields)
	for lang, i18n := range fields {
		ret[lang] = AdminI18nFields{
			Title:             i18n.Title,
			DetailImageURLs:   i18n.DetailImageURLs,
			DetailDescription: i18n.DetailDescription,
			Description:       i18n.Description,
		}
	}
	return ret
}
