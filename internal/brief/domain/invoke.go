package domain

import (
	"context"

	app "github.com/channel-io/ch-app-store/internal/app/domain"
)

type BriefResponse struct {
	Brief string
}

type BriefRequest struct {
	AppID     string
	ChannelID string
	app.Body
}

type InvokeSvc struct {
	repo    BriefRepository
	invoker app.Invoker[BriefResponse]
}

func (s *InvokeSvc) Invoke(ctx context.Context, req BriefRequest) (BriefResponse, error) {
	brief, err := s.repo.Fetch(ctx, req.AppID)
	if err != nil {
		return BriefResponse{}, err
	}

	ctxReq := app.FunctionRequest{
		Endpoint: app.Endpoint{
			AppID:        req.AppID,
			FunctionName: brief.BriefFunctionName,
		},
		Body: req.Body,
	}

	return s.invoker.InvokeChannelFunction(ctx, req.ChannelID, ctxReq)
}
