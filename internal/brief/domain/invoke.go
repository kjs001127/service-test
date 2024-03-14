package domain

import (
	"context"
	"sync"

	"github.com/pkg/errors"

	app "github.com/channel-io/ch-app-store/internal/app/domain"
	"github.com/channel-io/ch-app-store/lib/log"
)

type BriefResponses struct {
	Results []*AppBrief `json:"results"`
}

type AppBrief struct {
	AppId string `json:"appId"`
	Brief string `json:"brief"`
}

type BriefResponse struct {
	Result string `json:"result"`
}

type BriefRequest struct {
	Context   app.ChannelContext
	ChannelID string
}

type EmptyRequest struct {
}

type Invoker struct {
	repo     BriefRepository
	querySvc *app.QuerySvc
	invoker  *app.TypedInvoker[EmptyRequest, BriefResponse]
	logger   log.ContextAwareLogger
}

func NewInvoker(
	repo BriefRepository,
	querySvc *app.QuerySvc,
	invoker *app.TypedInvoker[EmptyRequest, BriefResponse],
	logger log.ContextAwareLogger,
) *Invoker {
	return &Invoker{repo: repo, querySvc: querySvc, invoker: invoker, logger: logger}
}

func (i *Invoker) Invoke(ctx context.Context, req app.ChannelContext) (BriefResponses, error) {
	installedApps, err := i.querySvc.QueryAll(ctx, req.Channel.ID)
	if err != nil {
		return BriefResponses{}, errors.WithStack(err)
	}

	briefs, err := i.repo.FetchAll(ctx, app.AppIDsOf(installedApps.AppChannels))
	if err != nil {
		return BriefResponses{}, errors.WithStack(err)
	}

	i.logger.Infow(ctx, "invoking brief",
		"channelID", req.Channel,
		"appIds", app.AppIDsOf(installedApps.AppChannels),
	)

	ch := make(chan *AppBrief, len(briefs))
	var wg sync.WaitGroup
	wg.Add(len(briefs))

	for _, brief := range briefs {
		brief := brief
		childCtx, cancel := context.WithCancel(ctx)
		go func() {
			defer cancel()
			res := i.invoker.Invoke(childCtx, brief.AppID, app.TypedRequest[EmptyRequest]{
				FunctionName: brief.BriefFunctionName,
				Context:      req,
			})
			if res.Error == nil {
				ch <- &AppBrief{AppId: brief.AppID, Brief: res.Result.Result}
			} else {
				i.logger.Warnw(childCtx, "brief request returned err", "error", res.Error)
			}
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	var res []*AppBrief

	for s := range ch {
		res = append(res, s)
	}

	return BriefResponses{Results: res}, nil
}
