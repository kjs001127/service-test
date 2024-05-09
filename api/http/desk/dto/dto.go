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
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	AvatarURL   *string `json:"avatarUrl,omitempty"`
	Description *string `json:"description,omitempty"`
	IsBuiltIn   bool    `json:"isBuiltIn"`

	// legacy fields
	State     string `json:"state"`
	IsPrivate bool   `json:"isPrivate"`
}

func NewAppView(origin *model.App) *AppView {
	return &AppView{
		ID:          origin.ID,
		Title:       origin.Title,
		AvatarURL:   origin.AvatarURL,
		Description: origin.Description,
		IsBuiltIn:   origin.IsBuiltIn,

		// legacy fields
		State:     "",
		IsPrivate: false,
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
	ID                 string           `json:"id"`
	Title              string           `json:"title"`
	AvatarURL          *string          `json:"avatarUrl,omitempty"`
	Description        *string          `json:"description,omitempty"`
	ManualURL          *string          `json:"manualUrl,omitempty"`
	DetailDescriptions []map[string]any `json:"detailDescriptions,omitempty"`
	DetailImageURLs    []string         `json:"detailImageUrls,omitempty"`
	IsBuiltIn          bool             `json:"isBuiltIn"`
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
	}
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

	Description     *string        `json:"description,omitempty"`
	NameDescI18nMap map[string]any `json:"nameDescI18nMap,omitempty"`

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
