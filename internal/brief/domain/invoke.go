package domain

import (
	"context"
	"sync"

	app "github.com/channel-io/ch-app-store/internal/app/domain"
)

type BriefResponses struct {
	Result []string `json:"result"`
}

type BriefResponse struct {
	Result string `json:"result"`
}

type BriefRequest struct {
}

type Invoker struct {
	repo     BriefRepository
	querySvc *app.QuerySvc
	invoker  *app.Invoker[BriefRequest, BriefResponse]
}

func NewInvoker(
	repo BriefRepository,
	querySvc *app.QuerySvc,
	invoker *app.Invoker[BriefRequest, BriefResponse],
) *Invoker {
	return &Invoker{repo: repo, querySvc: querySvc, invoker: invoker}
}

func (i *Invoker) Invoke(ctx context.Context, caller app.Caller, channelID string) (BriefResponses, error) {
	apps, err := i.querySvc.QueryAll(ctx, channelID)
	if err != nil {
		return BriefResponses{}, err
	}

	briefs, err := i.repo.FetchAll(ctx, app.AppIDsOf(apps.AppChannels))
	if err != nil {
		return BriefResponses{}, err
	}

	ch := make(chan string, len(briefs))
	var wg sync.WaitGroup
	wg.Add(len(briefs))

	for _, brief := range briefs {
		brief := brief
		go func() {
			res, err := i.invoker.InvokeChannelFunction(ctx, channelID, app.FunctionRequest[BriefRequest]{
				Endpoint: app.Endpoint{
					AppID:        brief.AppID,
					FunctionName: brief.BriefFunctionName,
				},
				Body: app.Body[BriefRequest]{
					Context: app.ChannelContext{
						Channel: app.Channel{ID: channelID},
						Caller:  caller,
					},
					Params: BriefRequest{},
				},
			})
			if err != nil {
				ch <- res.Result
			}
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	var res []string
	for s := range ch {
		res = append(res, s)
	}

	return BriefResponses{Result: res}, nil
}
