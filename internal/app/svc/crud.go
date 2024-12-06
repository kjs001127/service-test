package svc

import (
	"context"

	"github.com/channel-io/go-lib/pkg/errors/apierr"

	"github.com/channel-io/ch-app-store/internal/app/model"
	"github.com/channel-io/ch-app-store/lib/db/tx"

	"github.com/channel-io/go-lib/pkg/uid"
	"github.com/pkg/errors"
)

type AppLifeCycleEventListener interface {
	OnAppCreate(ctx context.Context, app *model.App) error
	OnAppDelete(ctx context.Context, app *model.App) error
	OnAppModify(ctx context.Context, before *model.App, after *model.App) error
}

type AppLifecycleSvc interface {
	Create(ctx context.Context, app *model.App) (*model.App, error)
	Delete(ctx context.Context, appID string) error
	Update(ctx context.Context, app *model.App) (*model.App, error)
	UpdateDetail(ctx context.Context, detail *AppDetail) (*AppDetail, error)
}

type AppQuerySvc interface {
	ReadDetail(ctx context.Context, appID string) (*AppDetail, error)
	Read(ctx context.Context, appID string) (*model.App, error)
	ReadAllByAppIDs(ctx context.Context, appIDs []string) ([]*model.App, error)
	ListPublicApps(ctx context.Context, since string, limit int) ([]*model.App, error)
}

type AppQuerySvcImpl struct {
	appRepo     AppRepository
	displayRepo AppDisplayRepository
}

func NewAppQuerySvcImpl(appRepo AppRepository, displayRepo AppDisplayRepository) *AppQuerySvcImpl {
	return &AppQuerySvcImpl{appRepo: appRepo, displayRepo: displayRepo}
}

func (a *AppQuerySvcImpl) Read(ctx context.Context, appID string) (*model.App, error) {
	return a.appRepo.Find(ctx, appID)
}

func (a *AppQuerySvcImpl) ReadAllByAppIDs(ctx context.Context, appIDs []string) ([]*model.App, error) {
	return a.appRepo.FindAll(ctx, appIDs)
}

type AppDetail struct {
	ID string `json:"id"`

	IsBuiltIn bool `json:"isBuiltIn"`
	IsPrivate bool `json:"isPrivate"`

	Title       string  `json:"title"`
	AvatarURL   *string `json:"avatarUrl,omitempty"`
	Description *string `json:"description,omitempty"`

	ManualURL          *string          `json:"manualUrl,omitempty"`
	DetailDescriptions []map[string]any `json:"detailDescriptions,omitempty"`
	DetailImageURLs    []string         `json:"detailImageUrls,omitempty"`

	I18nMap map[string]*DetailI18n `json:"i18NMap,omitempty"`
}

func (d *AppDetail) app() *model.App {
	i18n := make(map[string]model.I18nFields)
	for locale, origin := range d.I18nMap {
		i18n[locale] = model.I18nFields{
			Title:       origin.Title,
			Description: origin.Description,
		}
	}

	return &model.App{
		ID:          d.ID,
		Title:       d.Title,
		AvatarURL:   d.AvatarURL,
		Description: d.Description,
		I18nMap:     i18n,
		IsPrivate:   d.IsPrivate,
		IsBuiltIn:   d.IsBuiltIn,
	}
}

func (d *AppDetail) display() *model.AppDisplay {
	i18n := make(map[string]model.DisplayI18n)
	for locale, origin := range d.I18nMap {
		i18n[locale] = model.DisplayI18n{
			DetailImageURLs:    origin.DetailImageURLs,
			DetailDescriptions: origin.DetailDescriptions,
			ManualURL:          origin.ManualURL,
		}
	}

	return &model.AppDisplay{
		I18nMap:            i18n,
		ManualURL:          d.ManualURL,
		AppID:              d.ID,
		DetailDescriptions: d.DetailDescriptions,
		DetailImageURLs:    d.DetailImageURLs,
	}
}

type DetailI18n struct {
	model.DisplayI18n
	model.I18nFields
}

func (a *AppQuerySvcImpl) ReadDetail(ctx context.Context, appID string) (*AppDetail, error) {
	return tx.DoReturn(ctx, func(ctx context.Context) (*AppDetail, error) {
		app, err := a.Read(ctx, appID)
		if err != nil {
			return nil, err
		}

		display, err := a.displayRepo.Find(ctx, appID)
		if apierr.IsNotFound(err) {
			display = &model.AppDisplay{AppID: appID}
		} else if err != nil {
			return nil, errors.Wrap(err, "finding app display while reading detail")
		}

		return mergeDetail(app, display), nil
	}, tx.ReadOnly())
}

func mergeDetail(app *model.App, display *model.AppDisplay) *AppDetail {
	i18nMap := make(map[string]*DetailI18n)
	for locale, i18n := range app.I18nMap {
		if _, ok := i18nMap[locale]; !ok {
			i18nMap[locale] = &DetailI18n{}
		}
		i18nMap[locale].I18nFields = i18n
	}

	for locale, i18n := range display.I18nMap {
		if _, ok := i18nMap[locale]; !ok {
			i18nMap[locale] = &DetailI18n{}
		}
		i18nMap[locale].DisplayI18n = i18n
	}

	return &AppDetail{
		ID:                 app.ID,
		IsBuiltIn:          app.IsBuiltIn,
		IsPrivate:          app.IsPrivate,
		Title:              app.Title,
		AvatarURL:          app.AvatarURL,
		Description:        app.Description,
		ManualURL:          display.ManualURL,
		DetailDescriptions: display.DetailDescriptions,
		DetailImageURLs:    display.DetailImageURLs,
		I18nMap:            i18nMap,
	}
}

func (a *AppQuerySvcImpl) ListPublicApps(ctx context.Context, since string, limit int) ([]*model.App, error) {
	return a.appRepo.FindPublicApps(ctx, since, limit)
}

type AppLifecycleSvcImpl struct {
	appRepo     AppRepository
	displayRepo AppDisplayRepository

	installSvc         AppInstallSvc
	querySvc           *InstalledAppQuerySvc
	lifecycleListeners []AppLifeCycleEventListener
}

func NewAppLifecycleSvcImpl(
	appRepo AppRepository,
	displayRepo AppDisplayRepository,
	installSvc AppInstallSvc,
	querySvc *InstalledAppQuerySvc,
	lifecycleHooks []AppLifeCycleEventListener,
) *AppLifecycleSvcImpl {
	return &AppLifecycleSvcImpl{appRepo: appRepo, installSvc: installSvc, querySvc: querySvc, lifecycleListeners: lifecycleHooks, displayRepo: displayRepo}
}

func (a *AppLifecycleSvcImpl) Create(ctx context.Context, app *model.App) (*model.App, error) {
	return tx.DoReturn(ctx, func(ctx context.Context) (*model.App, error) {
		app.ID = uid.New().Hex()

		ret, err := a.appRepo.Save(ctx, app)
		if err != nil {
			return nil, errors.Wrap(err, "saving app while create")
		}

		if err := a.publishCreateEvent(ctx, app); err != nil {
			return nil, errors.Wrap(err, "publish trx event while create")
		}

		return ret, nil
	}, tx.XLock(namespaceApp, app.ID))
}

func (a *AppLifecycleSvcImpl) UpdateDetail(ctx context.Context, detail *AppDetail) (*AppDetail, error) {
	return tx.DoReturn(ctx, func(ctx context.Context) (*AppDetail, error) {
		appBefore, err := a.appRepo.Find(ctx, detail.ID)
		if err != nil {
			return nil, errors.Wrap(err, "finding app while update detail")
		}

		ret, err := a.appRepo.Save(ctx, detail.app())
		if err != nil {
			return nil, errors.Wrap(err, "saving app while updating detail")
		}

		display, err := a.displayRepo.Save(ctx, detail.display())
		if err != nil {
			return nil, errors.Wrap(err, "saving display while updating detail")
		}

		if err := a.publishModifyEvent(ctx, appBefore, ret); err != nil {
			return nil, errors.Wrap(err, "publish trx event while updating detail")
		}

		return mergeDetail(ret, display), nil
	}, tx.XLock(namespaceApp, detail.ID))
}

func (a *AppLifecycleSvcImpl) Update(ctx context.Context, app *model.App) (*model.App, error) {
	return tx.DoReturn(ctx, func(ctx context.Context) (*model.App, error) {
		appBefore, err := a.appRepo.Find(ctx, app.ID)
		if err != nil {
			return nil, errors.Wrap(err, "app ")
		}

		ret, err := a.appRepo.Save(ctx, app)
		if err != nil {
			return nil, errors.Wrap(err, "")
		}

		if err := a.publishModifyEvent(ctx, appBefore, app); err != nil {
			return nil, errors.Wrap(err, "'")
		}

		return ret, nil
	}, tx.XLock(namespaceApp, app.ID))
}

func (a *AppLifecycleSvcImpl) Delete(ctx context.Context, appID string) error {
	return tx.Do(ctx, func(ctx context.Context) error {

		if err := a.uninstallAll(ctx, appID); err != nil {
			return err
		}

		app, err := a.appRepo.Find(ctx, appID)
		if err != nil {
			return err
		}
		if err := a.appRepo.Delete(ctx, appID); err != nil {
			return errors.WithStack(err)
		}
		if err := a.displayRepo.Delete(ctx, appID); err != nil {
			return err
		}

		if err := a.publishDeleteEvent(ctx, app); err != nil {
			return err
		}
		return nil
	}, tx.XLock(namespaceApp, appID))
}

func (a *AppLifecycleSvcImpl) uninstallAll(ctx context.Context, appID string) error {
	installations, err := a.querySvc.QueryInstallationsByAppID(ctx, appID)
	if err != nil {
		return err
	}

	for _, installation := range installations {
		if err := a.installSvc.UnInstallApp(ctx, installation.ID()); err != nil {
			return err
		}
	}

	return nil
}

func (a *AppLifecycleSvcImpl) publishDeleteEvent(ctx context.Context, app *model.App) error {
	for _, h := range a.lifecycleListeners {
		if err := h.OnAppDelete(ctx, app); err != nil {
			return err
		}
	}
	return nil
}

func (a *AppLifecycleSvcImpl) publishModifyEvent(ctx context.Context, before *model.App, after *model.App) error {
	for _, h := range a.lifecycleListeners {
		if err := h.OnAppModify(ctx, before, after); err != nil {
			return err
		}
	}
	return nil
}

func (a *AppLifecycleSvcImpl) publishCreateEvent(ctx context.Context, app *model.App) error {
	for _, h := range a.lifecycleListeners {
		if err := h.OnAppCreate(ctx, app); err != nil {
			return err
		}
	}
	return nil
}
