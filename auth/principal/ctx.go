package principal

import (
	"context"

	"github.com/channel-io/ch-app-store/auth/general"
)

type Authorizer struct {
	parser           *general.Parser
	exchanger        *general.PrincipalToRBACExchanger
	clientIDProvider ClientIDProvider
}

type ClientIDProvider interface {
	FetchClientID(appID string) (string, error)
}

func (p *Authorizer) Handle(ctx context.Context, req general.Request[Token]) (general.Scopes, error) {
	clientID, err := p.clientIDProvider.FetchClientID(req.AppID)
	if err != nil {
		return nil, err
	}

	if _, err := p.exchanger.Exchange(ctx, req.Token, general.ExtractScopes(req.ChCtx), clientID); err != nil {
		return nil, err
	}

	return general.ExtractScopes(req.ChCtx), nil
}
