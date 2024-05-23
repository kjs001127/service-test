package dto

import (
	"time"

	"github.com/channel-io/ch-app-store/internal/app/model"
	cmd "github.com/channel-io/ch-app-store/internal/command/model"
)

type InstalledAppDetailView struct {
	App            *AppDetailView `json:"app"`
	Commands       []*CommandView `json:"commands"`
	CommandEnabled bool           `json:"commandEnabled"`
}

type AppStoreDetailView struct {
	App      *AppDetailView `json:"app"`
	Commands []*CommandView `json:"commands"`
}

type AppView struct {
	ID          string                 `json:"id"`
	Title       string                 `json:"title"`
	AvatarURL   *string                `json:"avatarUrl,omitempty"`
	Description *string                `json:"description,omitempty"`
	IsBuiltIn   bool                   `json:"isBuiltIn"`
	I18nMap     map[string]AppViewI18n `json:"i18nMap"`
	IsPrivate   bool                   `json:"isPrivate"`

	// legacy fields
	State string `json:"state"`
}

type AppViewI18n struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func convertAppViewI18n(app *model.App) map[string]AppViewI18n {
	ret := make(map[string]AppViewI18n)
	for lang, i18n := range app.I18nMap {
		ret[lang] = AppViewI18n{
			Title:       i18n.Title,
			Description: i18n.Description,
		}
	}
	return ret
}

func NewAppView(origin *model.App) *AppView {
	return &AppView{
		ID:          origin.ID,
		Title:       origin.Title,
		AvatarURL:   origin.AvatarURL,
		Description: origin.Description,
		IsBuiltIn:   origin.IsBuiltIn,
		I18nMap:     convertAppViewI18n(origin),
		IsPrivate:   origin.IsPrivate,

		// legacy fields
		State: "",
	}
}

func NewAppViews(origins []*model.App) []*AppView {
	ret := make([]*AppView, 0, len(origins))
	for _, origin := range origins {
		ret = append(ret, NewAppView(origin))
	}
	return ret
}

type WysiwygView struct {
	Apps     []*AppView     `json:"apps"`
	Commands []*CommandView `json:"commands"`
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

func NewAppDetailView(origin *model.App) *AppDetailView {
	return &AppDetailView{
		ID:                 origin.ID,
		Title:              origin.Title,
		AvatarURL:          origin.AvatarURL,
		Description:        origin.Description,
		IsBuiltIn:          origin.IsBuiltIn,
		ManualURL:          origin.ManualURL,
		DetailDescriptions: origin.DetailDescriptions,
		DetailImageURLs:    origin.DetailImageURLs,
		I18nMap:            convertAppDetailI18n(origin),
	}
}

type AppDetailI18n struct {
	Title              string           `json:"title"`
	Description        string           `json:"description"`
	DetailDescriptions []map[string]any `json:"detailDescriptions,omitempty"`
	DetailImageURLs    []string         `json:"detailImageUrls,omitempty"`
}

func convertAppDetailI18n(app *model.App) map[string]AppDetailI18n {
	ret := make(map[string]AppDetailI18n)
	for lang, i18n := range app.I18nMap {
		ret[lang] = AppDetailI18n{
			Title:              i18n.Title,
			Description:        i18n.Description,
			DetailDescriptions: i18n.DetailDescriptions,
			DetailImageURLs:    i18n.DetailImageURLs,
		}
	}
	return ret
}

func NewAppDetailViews(origins []*model.App) []*AppDetailView {
	ret := make([]*AppDetailView, 0, len(origins))
	for _, origin := range origins {
		ret = append(ret, NewAppDetailView(origin))
	}
	return ret
}

type CommandView struct {
	AppID string    `json:"appId,o"`
	Name  string    `json:"name"`
	Scope cmd.Scope `json:"scope"`

	Description     *string                `json:"description,omitempty"`
	NameDescI18nMap map[string]cmd.I18nMap `json:"nameDescI18nMap,omitempty"`

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
	Language       string `json:"language"`
	CommandEnabled bool   `json:"commandEnabled"`
}
