package svc

import (
	"context"

	appmodel "github.com/channel-io/ch-app-store/internal/app/model"
	appsvc "github.com/channel-io/ch-app-store/internal/app/svc"
	displaymodel "github.com/channel-io/ch-app-store/internal/appdisplay/model"
	"github.com/channel-io/ch-app-store/lib/db/tx"
)

type AppWithDisplay struct {
	ID    string   `json:"id"`
	State AppState `json:"state"`

	Title       string  `json:"title"`
	AvatarURL   *string `json:"avatarUrl,omitempty"`
	Description *string `json:"description,omitempty"`

	IsPrivate          bool             `json:"isPrivate"`
	ManualURL          *string          `json:"manualUrl,omitempty"`
	DetailDescriptions []map[string]any `json:"detailDescriptions,omitempty"`
	DetailImageURLs    []string         `json:"detailImageUrls,omitempty"`

	I18nMap map[string]I18nFields `json:"i18NMap,omitempty"`

	IsBuiltIn bool `json:"isBuiltIn"`
}

type AppState string

type I18nFields struct {
	Title              string           `json:"title"`
	DetailImageURLs    []string         `json:"detailImageUrls,omitempty"`
	DetailDescriptions []map[string]any `json:"detailDescriptions,omitempty"`
	Description        string           `json:"description,omitempty"`
	ManualURL          string           `json:"manualURL,omitempty"`
}

func (a *AppWithDisplay) ConvertToApp() *appmodel.App {
	return &appmodel.App{
		ID:          a.ID,
		Title:       a.Title,
		AvatarURL:   a.AvatarURL,
		Description: a.Description,
		I18nMap:     convertToAppI18nMap(a.I18nMap),
		IsBuiltIn:   a.IsBuiltIn,
	}
}

func (a *AppWithDisplay) ConvertToDisplay() *displaymodel.AppDisplay {
	return &displaymodel.AppDisplay{
		AppID:              a.ID,
		IsPrivate:          a.IsPrivate,
		ManualURL:          a.ManualURL,
		DetailDescriptions: a.DetailDescriptions,
		DetailImageURLs:    a.DetailImageURLs,
		I18nMap:            convertToDisplayI18nMap(a.I18nMap),
	}
}

func convertToDisplayI18nMap(i18nMap map[string]I18nFields) map[string]displaymodel.I18nFields {
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

func convertToAppI18nMap(i18nMap map[string]I18nFields) map[string]appmodel.I18nFields {
	ret := make(map[string]appmodel.I18nFields)
	for lang, i18n := range i18nMap {
		ret[lang] = appmodel.I18nFields{
			Title:       i18n.Title,
			Description: i18n.Description,
		}
	}
	return ret
}

type DisplayLifecycleSvc interface {
	Update(ctx context.Context, display *displaymodel.AppDisplay) (*displaymodel.AppDisplay, error)
}

type AppWithDisplayQuerySvc interface {
	Read(ctx context.Context, appID string) (*AppWithDisplay, error)
	ReadPublicAppsWithDisplay(ctx context.Context, since string, limit int) ([]*AppWithDisplay, error)
	ReadAllByAppIDs(ctx context.Context, appIDs []string) ([]*AppWithDisplay, error)
	AddDisplayToApps(ctx context.Context, apps []*appmodel.App) ([]*AppWithDisplay, error)
	AddDisplayToApp(ctx context.Context, apps *appmodel.App) (*AppWithDisplay, error)
	AddAppToDisplays(ctx context.Context, displays []*displaymodel.AppDisplay) ([]*AppWithDisplay, error)
}

type AppWithDisplayQuerySvcImpl struct {
	displayRepo AppDisplayRepository
	appRepo     appsvc.AppRepository
}

func NewAppWithDisplayQuerySvcImpl(displayRepo AppDisplayRepository, appRepo appsvc.AppRepository) *AppWithDisplayQuerySvcImpl {
	return &AppWithDisplayQuerySvcImpl{displayRepo: displayRepo, appRepo: appRepo}
}

func (a *AppWithDisplayQuerySvcImpl) AddAppToDisplays(ctx context.Context, displays []*displaymodel.AppDisplay) ([]*AppWithDisplay, error) {
	appIDs := make([]string, 0, len(displays))
	for _, display := range displays {
		appIDs = append(appIDs, display.AppID)
	}

	apps, err := a.appRepo.FindApps(ctx, appIDs)
	if err != nil {
		return nil, err
	}

	appMap := make(map[string]*appmodel.App)
	for _, app := range apps {
		appMap[app.ID] = app
	}

	ret := make([]*AppWithDisplay, 0, len(apps))
	for _, display := range displays {
		ret = append(ret, MergeAppAndDisplay(appMap[display.AppID], display))
	}

	return ret, nil
}

func (a *AppWithDisplayQuerySvcImpl) AddDisplayToApp(ctx context.Context, app *appmodel.App) (*AppWithDisplay, error) {
	display, err := a.displayRepo.FindDisplay(ctx, app.ID)
	if err != nil {
		return nil, err
	}

	return MergeAppAndDisplay(app, display), nil
}

func (a *AppWithDisplayQuerySvcImpl) AddDisplayToApps(ctx context.Context, apps []*appmodel.App) ([]*AppWithDisplay, error) {
	appIDs := make([]string, 0, len(apps))
	for _, app := range apps {
		appIDs = append(appIDs, app.ID)
	}

	displays, err := a.displayRepo.FindDisplays(ctx, appIDs)
	if err != nil {
		return nil, err
	}

	displayMap := make(map[string]*displaymodel.AppDisplay)
	for _, display := range displays {
		displayMap[display.AppID] = display
	}

	ret := make([]*AppWithDisplay, 0, len(apps))
	for _, app := range apps {
		ret = append(ret, MergeAppAndDisplay(app, displayMap[app.ID]))
	}

	return ret, nil
}

func (a *AppWithDisplayQuerySvcImpl) Read(ctx context.Context, appID string) (*AppWithDisplay, error) {
	return tx.DoReturn(ctx, func(ctx context.Context) (*AppWithDisplay, error) {
		display, err := a.displayRepo.FindDisplay(ctx, appID)
		if err != nil {
			return nil, err
		}

		app, err := a.appRepo.FindApp(ctx, appID)
		if err != nil {
			return nil, err
		}

		return MergeAppAndDisplay(app, display), nil
	}, tx.ReadOnly())
}

func (a *AppWithDisplayQuerySvcImpl) ReadPublicAppsWithDisplay(ctx context.Context, since string, limit int) ([]*AppWithDisplay, error) {
	return tx.DoReturn(ctx, func(ctx context.Context) ([]*AppWithDisplay, error) {
		publicDisplays, err := a.displayRepo.FindPublicDisplays(ctx, since, limit)
		if err != nil {
			return nil, err
		}

		appIDs := make([]string, 0, len(publicDisplays))
		displayMap := make(map[string]*displaymodel.AppDisplay)
		for _, publicDisplay := range publicDisplays {
			appIDs = append(appIDs, publicDisplay.AppID)
			displayMap[publicDisplay.AppID] = publicDisplay
		}

		publicApps, err := a.appRepo.FindApps(ctx, appIDs)
		if err != nil {
			return nil, err
		}

		ret := make([]*AppWithDisplay, 0, len(publicApps))
		for _, publicApp := range publicApps {
			ret = append(ret, MergeAppAndDisplay(publicApp, displayMap[publicApp.ID]))
		}

		return ret, nil
	}, tx.ReadOnly())
}

func (a *AppWithDisplayQuerySvcImpl) ReadAllByAppIDs(ctx context.Context, appIDs []string) ([]*AppWithDisplay, error) {
	return tx.DoReturn(ctx, func(ctx context.Context) ([]*AppWithDisplay, error) {
		displays, err := a.displayRepo.FindDisplays(ctx, appIDs)
		if err != nil {
			return nil, err
		}

		displayMap := make(map[string]*displaymodel.AppDisplay)
		for _, display := range displays {
			displayMap[display.AppID] = display
		}

		apps, err := a.appRepo.FindApps(ctx, appIDs)
		if err != nil {
			return nil, err
		}

		ret := make([]*AppWithDisplay, 0, len(apps))
		for _, app := range apps {
			ret = append(ret, MergeAppAndDisplay(app, displayMap[app.ID]))
		}

		return ret, nil
	}, tx.ReadOnly())
}

func MergeAppAndDisplay(app *appmodel.App, display *displaymodel.AppDisplay) *AppWithDisplay {
	return &AppWithDisplay{
		ID:                 app.ID,
		State:              "",
		Title:              app.Title,
		AvatarURL:          app.AvatarURL,
		Description:        app.Description,
		IsPrivate:          display.IsPrivate,
		ManualURL:          display.ManualURL,
		DetailDescriptions: display.DetailDescriptions,
		DetailImageURLs:    display.DetailImageURLs,
		I18nMap:            mergeI18nMaps(app.I18nMap, display.I18nMap),
		IsBuiltIn:          app.IsBuiltIn,
	}
}

func mergeI18nMaps(appI18nMap map[string]appmodel.I18nFields, displayI18nMap map[string]displaymodel.I18nFields) map[string]I18nFields {
	i18nMap := make(map[string]I18nFields)
	for lang, appFields := range appI18nMap {
		i18nMap[lang] = I18nFields{
			Title:              appFields.Title,
			DetailDescriptions: displayI18nMap[lang].DetailDescriptions,
			DetailImageURLs:    displayI18nMap[lang].DetailImageURLs,
			ManualURL:          displayI18nMap[lang].ManualURL,
			Description:        appFields.Description,
		}
	}
	for lang, displayFields := range displayI18nMap {
		_, ok := i18nMap[lang]
		if !ok {
			i18nMap[lang] = I18nFields{
				Title:              appI18nMap[lang].Title,
				DetailDescriptions: displayFields.DetailDescriptions,
				DetailImageURLs:    displayFields.DetailImageURLs,
				ManualURL:          displayFields.ManualURL,
				Description:        appI18nMap[lang].Description,
			}
		}
	}
	return i18nMap
}

type DisplayLifecycleSvcImpl struct {
	appDisplayRepo AppDisplayRepository
}

func NewDisplayLifecycleSvcImpl(
	appDisplayRepo AppDisplayRepository,
) *DisplayLifecycleSvcImpl {
	return &DisplayLifecycleSvcImpl{appDisplayRepo: appDisplayRepo}
}

func (a *DisplayLifecycleSvcImpl) Update(ctx context.Context, display *displaymodel.AppDisplay) (*displaymodel.AppDisplay, error) {
	return tx.DoReturn(ctx, func(ctx context.Context) (*displaymodel.AppDisplay, error) {
		_, err := a.appDisplayRepo.FindDisplay(ctx, display.AppID)
		if err != nil {
			return nil, err
		}

		ret, err := a.appDisplayRepo.Save(ctx, display)
		if err != nil {
			return nil, err
		}

		return ret, nil
	})
}
