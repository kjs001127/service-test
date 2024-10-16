package svc

import (
	"context"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
	"github.com/pkg/errors"

	appmodel "github.com/channel-io/ch-app-store/internal/app/model"
	"github.com/channel-io/ch-app-store/internal/app/svc"
	"github.com/channel-io/ch-app-store/internal/appwidget/model"
	"github.com/channel-io/ch-app-store/internal/auth/principal/account"
	"github.com/channel-io/ch-app-store/internal/auth/principal/session"
)

type AppWidgetInvoker interface {
	InvokeFrontWidget(ctx context.Context, invoker *session.UserRequester, appWidgetID string) (*Action, error)
	InvokeDeskWidget(ctx context.Context, invoker *account.ManagerRequester, appWidgetID string, request AppWidgetRequest) (*Action, error)
	IsInvocable(ctx context.Context, installation appmodel.InstallationID, appWidgetID string, scope model.Scope) (*model.AppWidget, error)
}

type AppWidgetRequest struct {
	Chat     Chat   `json:"chat"`
	Language string `json:"language"`
}

type Chat struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

type Action struct {
	Type       string         `json:"type"`
	Attributes map[string]any `json:"attributes"`
}

type AppWidgetInvokerImpl struct {
	repo            AppWidgetRepository
	installQuerySvc *svc.InstalledAppQuerySvc
	delegate        svc.TypedInvoker[*AppWidgetRequest, *Action]
}

func NewAppWidgetInvokerImpl(repo AppWidgetRepository, installQuerySvc *svc.InstalledAppQuerySvc, delegate svc.Invoker) *AppWidgetInvokerImpl {
	return &AppWidgetInvokerImpl{
		repo:            repo,
		installQuerySvc: installQuerySvc,
		delegate:        svc.NewTypedInvoker[*AppWidgetRequest, *Action](delegate),
	}
}

func (i *AppWidgetInvokerImpl) InvokeFrontWidget(ctx context.Context, invoker *session.UserRequester, appWidgetID string) (*Action, error) {
	widget, err := i.repo.FetchByIDAndScope(ctx, appWidgetID, model.ScopeFront)
	if err != nil {
		return nil, err
	}

	resp := i.delegate.Invoke(ctx, widget.AppID, svc.TypedRequest[*AppWidgetRequest]{
		FunctionName: widget.ActionFunctionName,
		Context: svc.ChannelContext{
			Channel: svc.Channel{
				ID: invoker.ChannelID,
			},
			Caller: svc.Caller{
				ID:   invoker.ID,
				Type: svc.CallerTypeUser,
			},
		},
		Params: &AppWidgetRequest{
			Language: invoker.Language,
		},
	})

	if resp.IsError() {
		return nil, resp.Error
	}

	return resp.Result, nil
}

func (i *AppWidgetInvokerImpl) InvokeDeskWidget(ctx context.Context, invoker *account.ManagerRequester, appWidgetID string, request AppWidgetRequest) (*Action, error) {
	widget, err := i.repo.FetchByIDAndScope(ctx, appWidgetID, model.ScopeDesk)
	if err != nil {
		return nil, err
	}

	request.Language = invoker.Language

	resp := i.delegate.Invoke(ctx, widget.AppID, svc.TypedRequest[*AppWidgetRequest]{
		FunctionName: widget.ActionFunctionName,
		Context: svc.ChannelContext{
			Channel: svc.Channel{
				ID: invoker.ChannelID,
			},
			Caller: svc.Caller{
				ID:   invoker.ID,
				Type: svc.CallerTypeManager,
			},
		},
		Params: &request,
	})

	if resp.IsError() {
		return nil, resp.Error
	}

	return resp.Result, nil
}

func (i *AppWidgetInvokerImpl) IsInvocable(ctx context.Context, install appmodel.InstallationID, appWidgetID string, scope model.Scope) (*model.AppWidget, error) {
	widget, err := i.repo.FetchByIDAndScope(ctx, appWidgetID, scope)
	if err != nil {
		return nil, err
	}

	exists, err := i.installQuerySvc.CheckInstall(ctx, install)

	if err != nil {
		return nil, err
	} else if !exists {
		return nil, apierr.NotFound(errors.New("app is not installed on channel"))
	}

	return widget, nil
}
