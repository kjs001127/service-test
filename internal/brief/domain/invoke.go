package domain

import (
	"context"

	app "github.com/channel-io/ch-app-store/internal/app/domain"
)

type BriefResponse struct {
	Brief string
}

type BriefClientRequest struct {
	AppID   string
	Token   app.AuthToken
	Context app.ChannelContext
}

type BriefRequest struct {
	Context app.ChannelContext
}

func (b BriefRequest) ChannelContext() app.ChannelContext {
	return b.Context
}

type InvokeSvc struct {
	repo    BriefRepository
	invoker app.ContextFnInvoker[BriefRequest, BriefResponse]
}

func (s *InvokeSvc) Invoke(ctx context.Context, req BriefClientRequest) (BriefResponse, error) {
	brief, err := s.repo.Fetch(ctx, req.AppID)
	if err != nil {
		return BriefResponse{}, err
	}

	ctxReq := app.Request[BriefRequest]{
		Token: req.Token,
		FunctionRequest: app.FunctionRequest[BriefRequest]{
			AppID:        req.AppID,
			FunctionName: brief.FunctionName,
			Body: BriefRequest{
				Context: req.Context,
			},
		},
	}

	return s.invoker.Invoke(ctx, ctxReq)
}
