package svc

import (
	"context"

	appmodel "github.com/channel-io/ch-app-store/internal/app/model"
	app "github.com/channel-io/ch-app-store/internal/app/svc"
	"github.com/channel-io/ch-app-store/internal/permission/repo"
	"github.com/channel-io/ch-app-store/lib/db/tx"
)

type AccountAppPermissionSvc interface {
	CreateApp(ctx context.Context, title string, accountID string) (*appmodel.App, error)
	ModifyApp(ctx context.Context, modifyRequest AppModifyRequest, appID string, accountID string) (*appmodel.App, error)
	DeleteApp(ctx context.Context, appID string, accountID string) error
}

type AccountAppPermissionSvcImpl struct {
	appCrudSvc     app.AppCrudSvc
	appAccountRepo repo.AppAccountRepo
}

type AppModifyRequest struct {
	Title              string                `json:"title"`
	Description        *string               `json:"description,omitempty"`
	DetailImageURLs    []string              `json:"detailImageUrls,omitempty"`
	DetailDescriptions []map[string]any      `json:"detailDescriptions,omitempty"`
	I18nMap            map[string]I18nFields `json:"i18nMap,omitempty"`
}

type I18nFields struct {
	Title             string   `json:"title"`
	DetailImageURLs   []string `json:"detailImageUrls,omitempty"`
	DetailDescription string   `json:"detailDescription,omitempty"`
	Description       string   `json:"description,omitempty"`
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
	}
}

func NewAccountAppPermissionSvc(
	appCrudSvc app.AppCrudSvc,
	appAccountRepo repo.AppAccountRepo,
) *AccountAppPermissionSvcImpl {
	return &AccountAppPermissionSvcImpl{
		appCrudSvc:     appCrudSvc,
		appAccountRepo: appAccountRepo,
	}
}

func (a *AccountAppPermissionSvcImpl) CreateApp(ctx context.Context, title string, accountID string) (*appmodel.App, error) {
	return tx.DoReturn(ctx, func(ctx context.Context) (*appmodel.App, error) {
		createApp := appmodel.App{
			Title: title,
		}

		app, err := a.appCrudSvc.Create(ctx, &createApp)
		if err != nil {
			return nil, err
		}

		err = a.appAccountRepo.Save(ctx, app.ID, accountID)
		if err != nil {
			return nil, err
		}
		return app, nil
	})
}

func (a *AccountAppPermissionSvcImpl) ModifyApp(ctx context.Context, modifyRequest AppModifyRequest, appID string, accountID string) (*appmodel.App, error) {
	return tx.DoReturn(ctx, func(ctx context.Context) (*appmodel.App, error) {
		_, err := a.appAccountRepo.Fetch(ctx, appID, accountID)
		if err != nil {
			return nil, err
		}

		_, err = a.appCrudSvc.Read(ctx, appID)
		if err != nil {
			return nil, err
		}

		converted := modifyRequest.ConvertToApp(appID)

		return a.appCrudSvc.Update(ctx, converted)
	})
}

func (a *AccountAppPermissionSvcImpl) DeleteApp(ctx context.Context, appID string, accountID string) error {
	return tx.Do(ctx, func(ctx context.Context) error {
		if _, err := a.appAccountRepo.Fetch(ctx, appID, accountID); err != nil {
			return err
		}
		err := a.appCrudSvc.Delete(ctx, appID)
		if err != nil {
			return err
		}
		return a.appAccountRepo.Delete(ctx, appID, accountID)
	})
}
