package domain

import (
	"context"

	app "github.com/channel-io/ch-app-store/internal/app/domain"
)

type RemoteApp struct {
	app.AppData

	RoleID   string
	ClientID string
	Secret   string

	HookURL     *string
	FunctionURL *string
	WamURL      *string
	CheckURL    *string

	requester HttpRequester
}

func (a *RemoteApp) Data() *app.AppData {
	return &a.AppData
}

type RemoteAppRepository interface {
	Index(ctx context.Context, since string, limit int) ([]*RemoteApp, error)
	Fetch(ctx context.Context, appID string) (*RemoteApp, error)
	FindAll(ctx context.Context, appIDs []string) ([]*RemoteApp, error)
	Save(ctx context.Context, app *RemoteApp) (*RemoteApp, error)
	Update(ctx context.Context, app *RemoteApp) (*RemoteApp, error)
	Delete(ctx context.Context, appID string) error
}

type AppRepositoryAdapter struct {
	appRepository RemoteAppRepository
	requester     HttpRequester
}

func NewAppRepositoryAdapter(appRepository RemoteAppRepository, requester HttpRequester) *AppRepositoryAdapter {
	return &AppRepositoryAdapter{appRepository: appRepository, requester: requester}
}

func (i *AppRepositoryAdapter) Index(ctx context.Context, since string, limit int) ([]app.App, error) {
	apps, err := i.appRepository.Index(ctx, since, limit)
	if err != nil {
		return nil, err
	}

	return i.toApps(apps), nil
}

func (i *AppRepositoryAdapter) FindApps(ctx context.Context, appIDs []string) ([]app.App, error) {
	apps, err := i.appRepository.FindAll(ctx, appIDs)
	if err != nil {
		return nil, err
	}

	return i.toApps(apps), nil
}

func (i *AppRepositoryAdapter) FindApp(ctx context.Context, appID string) (app.App, error) {
	one, err := i.appRepository.Fetch(ctx, appID)
	if err != nil {
		return nil, err
	}
	one.requester = i.requester
	return one, nil
}

func (i *AppRepositoryAdapter) toApps(apps []*RemoteApp) []app.App {
	ret := make([]app.App, len(apps))
	for _, a := range apps {
		a.requester = i.requester
		ret = append(ret, a)
	}
	return ret
}

type ClientIDProviderAdapter struct {
	repo RemoteAppRepository
}

func NewClientIDProviderAdapter(repo RemoteAppRepository) *ClientIDProviderAdapter {
	return &ClientIDProviderAdapter{repo: repo}
}

func (c *ClientIDProviderAdapter) FetchClientID(ctx context.Context, appID string) (string, error) {
	found, err := c.repo.Fetch(ctx, appID)
	if err != nil {
		return "", err
	}

	return found.ClientID, nil
}
