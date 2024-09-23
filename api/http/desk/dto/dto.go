package dto

import (
	"time"

	appmodel "github.com/channel-io/ch-app-store/internal/app/model"
	displaysvc "github.com/channel-io/ch-app-store/internal/appdisplay/svc"
	"github.com/channel-io/ch-app-store/internal/appwidget/model"
	cmd "github.com/channel-io/ch-app-store/internal/command/model"
	"github.com/channel-io/ch-app-store/internal/command/svc"
)

type InstalledAppDetailView struct {
	App      *AppDetailView          `json:"app"`
	Commands []*InstalledCommandView `json:"commands"`
}

type AppStoreDetailView struct {
	App      *AppDetailView `json:"app"`
	Commands []*CommandView `json:"commands"`
}

type AppStoreAppView struct {
	ID          string                 `json:"id"`
	Title       string                 `json:"title"`
	AvatarURL   *string                `json:"avatarUrl,omitempty"`
	Description *string                `json:"description,omitempty"`
	IsBuiltIn   bool                   `json:"isBuiltIn"`
	I18nMap     map[string]AppViewI18n `json:"i18nMap"`
	IsPrivate   bool                   `json:"isPrivate"`
}

type SimpleAppView struct {
	ID          string                 `json:"id"`
	Title       string                 `json:"title"`
	AvatarURL   *string                `json:"avatarUrl,omitempty"`
	Description *string                `json:"description,omitempty"`
	IsBuiltIn   bool                   `json:"isBuiltIn"`
	I18nMap     map[string]AppViewI18n `json:"i18nMap"`
}

type AppViewI18n struct {
	Title       string `json:"title"`
	Description string `json:"description"`
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

func NewAppWithDisplayView(origin *displaysvc.AppWithDisplay) *AppStoreAppView {
	return &AppStoreAppView{
		ID:          origin.ID,
		Title:       origin.Title,
		AvatarURL:   origin.AvatarURL,
		Description: origin.Description,
		IsBuiltIn:   origin.IsBuiltIn,
		I18nMap:     convertAppWithDisplayViewI18n(origin),
		IsPrivate:   origin.IsPrivate,
	}
}

func NewAppWithDisplayViews(origins []*displaysvc.AppWithDisplay) []*AppStoreAppView {
	ret := make([]*AppStoreAppView, 0, len(origins))
	for _, origin := range origins {
		ret = append(ret, NewAppWithDisplayView(origin))
	}
	return ret
}

func NewAppView(origin *appmodel.App) *SimpleAppView {
	return &SimpleAppView{
		ID:          origin.ID,
		Title:       origin.Title,
		AvatarURL:   origin.AvatarURL,
		Description: origin.Description,
		IsBuiltIn:   origin.IsBuiltIn,
		I18nMap:     convertAppViewI18n(origin),
	}
}

func NewAppViews(origins []*appmodel.App) []*SimpleAppView {
	ret := make([]*SimpleAppView, 0, len(origins))
	for _, origin := range origins {
		ret = append(ret, NewAppView(origin))
	}
	return ret
}

type WysiwygView struct {
	Apps     []*SimpleAppView `json:"apps"`
	Commands []*CommandView   `json:"commands"`
}

type AppDetailView struct {
	ID                 string                   `json:"id"`
	Title              string                   `json:"title"`
	AvatarURL          *string                  `json:"avatarUrl,omitempty"`
	Description        *string                  `json:"description,omitempty"`
	ManualURL          *string                  `json:"manualUrl,omitempty"`
	DetailDescriptions []map[string]any         `json:"detailDescriptions,omitempty"`
	DetailImageURLs    []string                 `json:"detailImageUrls,omitempty"`
	I18nMap            map[string]AppDetailI18n `json:"i18nMap"`
	IsBuiltIn          bool                     `json:"isBuiltIn"`
}

func NewAppWithDisplayDetailView(origin *displaysvc.AppWithDisplay) *AppDetailView {
	return &AppDetailView{
		ID:                 origin.ID,
		Title:              origin.Title,
		AvatarURL:          origin.AvatarURL,
		Description:        origin.Description,
		IsBuiltIn:          origin.IsBuiltIn,
		ManualURL:          origin.ManualURL,
		DetailDescriptions: origin.DetailDescriptions,
		DetailImageURLs:    origin.DetailImageURLs,
		I18nMap:            convertAppWithDisplayI18n(origin),
	}
}

type AppDetailI18n struct {
	Title              string           `json:"title"`
	Description        string           `json:"description"`
	DetailDescriptions []map[string]any `json:"detailDescriptions,omitempty"`
	DetailImageURLs    []string         `json:"detailImageUrls,omitempty"`
	ManualURL          string           `json:"manualUrl,omitempty"`
}

func convertAppWithDisplayI18n(app *displaysvc.AppWithDisplay) map[string]AppDetailI18n {
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

type InstalledCommandView struct {
	*CommandView
	Enabled bool `json:"enabled"`
}

func NewInstalledCommandView(cmd *svc.CommandWithActivation) *InstalledCommandView {
	return &InstalledCommandView{
		CommandView: NewCommandView(cmd.Command),
		Enabled:     cmd.Enabled,
	}
}

func NewInstalledCommandViews(origins []*svc.CommandWithActivation) []*InstalledCommandView {
	ret := make([]*InstalledCommandView, 0, len(origins))
	for _, origin := range origins {
		ret = append(ret, NewInstalledCommandView(origin))
	}
	return ret
}

type CommandView struct {
	AppID string    `json:"appId,o"`
	Name  string    `json:"name"`
	Scope cmd.Scope `json:"scope"`

	Description     *string                        `json:"description,omitempty"`
	NameDescI18nMap map[string]cmd.NameDescI18nMap `json:"nameDescI18nMap,omitempty"`

	ParamDefinitions cmd.ParamDefinitions `json:"paramDefinitions,omitempty"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewCommandView(origin *cmd.Command) *CommandView {
	return &CommandView{
		AppID:            origin.AppID,
		Name:             origin.Name,
		NameDescI18nMap:  origin.NameDescI18NMap,
		Scope:            origin.Scope,
		Description:      origin.Description,
		ParamDefinitions: origin.ParamDefinitions,
		CreatedAt:        origin.CreatedAt,
		UpdatedAt:        origin.UpdatedAt,
	}
}

func NewCommandViews(origins []*cmd.Command) []*CommandView {
	ret := make([]*CommandView, 0, len(origins))
	for _, origin := range origins {
		ret = append(ret, NewCommandView(origin))
	}
	return ret
}

type CommandToggleRequest struct {
	Scope    cmd.Scope `json:"scope"`
	Name     string    `json:"name"`
	Enabled  bool      `json:"enabled"`
	Language string    `json:"language"`
}

type AppsWithWidgetsView struct {
	Apps       []*SimpleAppView `json:"apps"`
	AppWidgets []*AppWidgetView `json:"appWidgets"`
}

type AppWidgetView struct {
	ID    string `json:"id"`
	AppID string `json:"appId"`

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
