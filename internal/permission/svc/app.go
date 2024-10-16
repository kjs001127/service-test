package svc

import (
	"context"
	"unicode/utf8"

	"github.com/pkg/errors"

	appmodel "github.com/channel-io/ch-app-store/internal/app/model"
	appsvc "github.com/channel-io/ch-app-store/internal/app/svc"
	"github.com/channel-io/ch-app-store/lib/db/tx"
)

const maxAppCountPerAccount = 30

type AppModifyRequest struct {
	Title              string                         `json:"title"`
	Description        *string                        `json:"description,omitempty"`
	DetailImageURLs    []string                       `json:"detailImageUrls,omitempty"`
	DetailDescriptions []map[string]any               `json:"detailDescriptions,omitempty"`
	ManualURL          *string                        `json:"manualUrl,omitempty"`
	AvatarUrl          *string                        `json:"avatarUrl,omitempty"`
	I18nMap            map[string]appmodel.I18nFields `json:"i18nMap,omitempty"`
}

func (r *AppModifyRequest) Validate() error {
	err := validateTitle(r.Title)
	if err != nil {
		return err
	}

	err = validateDescription(r.Description)
	if err != nil {
		return err
	}

	if r.I18nMap == nil {
		return nil
	}

	for _, item := range r.I18nMap {
		err = validateTitle(item.Title)
		if err != nil {
			return err
		}

		err = validateDescription(&item.Description)
		if err != nil {
			return err
		}
	}
	return nil
}

func validateTitle(title string) error {
	if utf8.RuneCountInString(title) < 2 || utf8.RuneCountInString(title) > 20 {
		return errors.New("Title length should be between 2 and 20")
	}
	return nil
}

func validateDescription(description *string) error {
	if description != nil && utf8.RuneCountInString(*description) > 200 {
		return errors.New("Description length should be less than 200")
	}
	return nil
}

func (r *AppModifyRequest) applyTo(target *appmodel.App) *appmodel.App {
	target.Title = r.Title
	target.I18nMap = r.I18nMap
	target.Description = r.Description
	target.AvatarURL = r.AvatarUrl
	return target
}

type AccountAppPermissionSvc interface {
	ReadApp(ctx context.Context, appID string, accountID string) (*appmodel.App, error)
	CreateApp(ctx context.Context, title string, accountID string) (*appmodel.App, error)
	ModifyApp(ctx context.Context, modifyRequest *AppModifyRequest, appID string, accountID string) (*appmodel.App, error)
	DeleteApp(ctx context.Context, appID string, accountID string) error
}

type AccountAppPermissionSvcImpl struct {
	appCrudSvc      appsvc.AppQuerySvc
	appLifeCycleSvc appsvc.AppLifecycleSvc
	appAccountRepo  AppAccountRepo
}

func NewAccountAppPermissionSvc(
	appCrudSvc appsvc.AppQuerySvc,
	appLifecycleSvc appsvc.AppLifecycleSvc,
	appAccountRepo AppAccountRepo,
) *AccountAppPermissionSvcImpl {
	return &AccountAppPermissionSvcImpl{
		appCrudSvc:      appCrudSvc,
		appAccountRepo:  appAccountRepo,
		appLifeCycleSvc: appLifecycleSvc,
	}
}

func (a *AccountAppPermissionSvcImpl) ReadApp(ctx context.Context, appID string, accountID string) (*appmodel.App, error) {
	if _, err := a.appAccountRepo.Fetch(ctx, appID, accountID); err != nil {
		return nil, err
	}

	ret, err := a.appCrudSvc.Read(ctx, appID)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (a *AccountAppPermissionSvcImpl) CreateApp(ctx context.Context, title string, accountID string) (*appmodel.App, error) {
	return tx.DoReturn(ctx, func(ctx context.Context) (*appmodel.App, error) {
		res, err := a.appAccountRepo.CountByAccountID(ctx, accountID)
		if err != nil {
			return nil, err
		}

		if res >= maxAppCountPerAccount {
			return nil, errors.New("App count per account must be less than 30")
		}

		createApp := appmodel.App{
			Title:     title,
			IsBuiltIn: false,
		}

		app, err := a.appLifeCycleSvc.Create(ctx, &createApp)
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

func (a *AccountAppPermissionSvcImpl) ModifyApp(ctx context.Context, modifyRequest *AppModifyRequest, appID string, accountID string) (*appmodel.App, error) {
	if err := modifyRequest.Validate(); err != nil {
		return nil, err
	}
	return tx.DoReturn(ctx, func(ctx context.Context) (*appmodel.App, error) {
		_, err := a.appAccountRepo.Fetch(ctx, appID, accountID)
		if err != nil {
			return nil, err
		}

		oldbie, err := a.appCrudSvc.Read(ctx, appID)
		if err != nil {
			return nil, err
		}

		newbie := modifyRequest.applyTo(oldbie)

		ret, err := a.appLifeCycleSvc.Update(ctx, newbie)
		if err != nil {
			return nil, err
		}

		return ret, nil
	})
}

func (a *AccountAppPermissionSvcImpl) DeleteApp(ctx context.Context, appID string, accountID string) error {
	return tx.Do(ctx, func(ctx context.Context) error {
		if _, err := a.appAccountRepo.Fetch(ctx, appID, accountID); err != nil {
			return err
		}
		err := a.appLifeCycleSvc.Delete(ctx, appID)
		if err != nil {
			return err
		}
		return a.appAccountRepo.Delete(ctx, appID, accountID)
	})
}
