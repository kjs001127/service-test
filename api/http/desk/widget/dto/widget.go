package dto

import (
	"github.com/channel-io/ch-app-store/api/http/desk/dto"
	"github.com/channel-io/ch-app-store/internal/widget/model"
)

type AppsWithWidgetsView struct {
	Apps       []*dto.SimpleAppView `json:"apps"`
	AppWidgets []*AppWidgetView     `json:"appWidgets"`
}

type AppWidgetView struct {
	ID    string      `json:"id"`
	AppID string      `json:"appId"`
	Scope model.Scope `json:"scope"`

	Name            string              `json:"name"`
	Description     *string             `json:"description,omitempty"`
	NameDescI18nMap map[string]*I18nMap `json:"nameDescI18nMap,omitempty"`

	DefaultName            *string             `json:"defaultName,omitempty"`
	DefaultDescription     *string             `json:"defaultDescription,omitempty"`
	DefaultNameDescI18nMap map[string]*I18nMap `json:"defaultNameDescI18nMap,omitempty"`
}

type I18nMap struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func NewAppWidgetViews(origins []*model.AppWidget) []*AppWidgetView {
	ret := make([]*AppWidgetView, 0, len(origins))
	for _, origin := range origins {
		ret = append(ret, &AppWidgetView{
			ID:                     origin.ID,
			Name:                   origin.Name,
			Scope:                  origin.Scope,
			AppID:                  origin.AppID,
			Description:            origin.Description,
			NameDescI18nMap:        nameDescI18nMap(origin.NameDescI18nMap),
			DefaultDescription:     origin.DefaultDescription,
			DefaultName:            origin.DefaultName,
			DefaultNameDescI18nMap: nameDescI18nMap(origin.DefaultNameDescI18nMap),
		})
	}
	return ret
}

func nameDescI18nMap(origin map[string]*model.I18nMap) map[string]*I18nMap {
	ret := make(map[string]*I18nMap)
	for key, val := range origin {
		ret[key] = &I18nMap{
			Name:        val.Name,
			Description: val.Description,
		}
	}
	return ret
}
