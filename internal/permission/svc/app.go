package svc

import (
	"context"
	"unicode/utf8"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
	"github.com/pkg/errors"

	appmodel "github.com/channel-io/ch-app-store/internal/app/model"
	appsvc "github.com/channel-io/ch-app-store/internal/app/svc"
	"github.com/channel-io/ch-app-store/lib/db/tx"
)

const maxAppCountPerAccount = 30

type AppModifyRequest struct {
	Title       string  `json:"title"`
	Description *string `json:"description,omitempty"`
	AvatarUrl   *string `json:"avatarUrl,omitempty"`

	DetailImageURLs    []string         `json:"detailImageUrls,omitempty"`
	DetailDescriptions []map[string]any `json:"detailDescriptions,omitempty"`
	ManualURL          *string          `json:"manualUrl,omitempty"`

	I18nMap map[string]*appsvc.DetailI18n `json:"i18nMap,omitempty"`
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

func (r *AppModifyRequest) applyTo(detail *appsvc.AppDetail) *appsvc.AppDetail {
	detail.Title = r.Title
	detail.I18nMap = r.I18nMap
	detail.Description = r.Description
	detail.DetailDescriptions = r.DetailDescriptions
	detail.DetailImageURLs = r.DetailImageURLs
	detail.ManualURL = r.ManualURL
	detail.AvatarURL = r.AvatarUrl

	return detail
}

type AccountAppPermissionSvc interface {
	ListApps(ctx context.Context, accountID string) ([]*appmodel.App, error)
	ReadApp(ctx context.Context, appID string, accountID string) (*appsvc.AppDetail, error)
	CreateApp(ctx context.Context, title string, accountID string) (*appmodel.App, error)
	ModifyApp(ctx context.Context, modifyRequest *AppModifyRequest, appID string, accountID string) (*appsvc.AppDetail, error)
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

func (a *AccountAppPermissionSvcImpl) ListApps(ctx context.Context, accountID string) ([]*appmodel.App, error) {
	appAccounts, err := a.appAccountRepo.FetchAllByAccountID(ctx, accountID)
	if err != nil {
		return nil, err
	}

	var appIDs []string
	for _, appAccount := range appAccounts {
		appIDs = append(appIDs, appAccount.AppID)
	}

	return a.appCrudSvc.ReadAllByAppIDs(ctx, appIDs)
}

func (a *AccountAppPermissionSvcImpl) ReadApp(ctx context.Context, appID string, accountID string) (*appsvc.AppDetail, error) {
	if _, err := a.appAccountRepo.Fetch(ctx, appID, accountID); err != nil {
		return nil, err
	}

	ret, err := a.appCrudSvc.ReadDetail(ctx, appID)
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
			return nil, apierr.Conflict(errors.New("App count per account must be less than 30"))
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

func (a *AccountAppPermissionSvcImpl) ModifyApp(ctx context.Context, modifyRequest *AppModifyRequest, appID string, accountID string) (*appsvc.AppDetail, error) {
	if err := modifyRequest.Validate(); err != nil {
		return nil, err
	}
	return tx.DoReturn(ctx, func(ctx context.Context) (*appsvc.AppDetail, error) {
		_, err := a.appAccountRepo.Fetch(ctx, appID, accountID)
		if err != nil {
			return nil, err
		}

		oldbie, err := a.appCrudSvc.ReadDetail(ctx, appID)
		if err != nil {
			return nil, err
		}

		newbie := modifyRequest.applyTo(oldbie)

		return a.appLifeCycleSvc.UpdateDetail(ctx, newbie)
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
