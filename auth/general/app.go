package general

import (
	"context"
	"errors"

	"github.com/channel-io/ch-app-store/auth/appauth"
	"github.com/channel-io/ch-app-store/auth/principal"
)

type Authorizer struct {
	parser *Parser
}

func NewAuthorizer(parser *Parser) *Authorizer {
	return &Authorizer{parser: parser}
}

func (g *Authorizer) Handle(ctx context.Context, req appauth.AppUseRequest[Token]) (appauth.Authorities, error) {
	rbac, err := g.parser.Parse(ctx, req.Token.Value())
	if err != nil {
		return nil, err
	}

	if ok := rbac.CheckAuthorities(req.ChCtx.RequiredAuthorities()); !ok {
		return nil, errors.New("scope check fail")
	}

	if ok := rbac.CheckAction(Service(req.AppID), Action(req.FunctionName)); !ok {
		return nil, errors.New("action check fail")
	}

	return rbac.Authorities, nil
}

type PrincipalAuthorizer struct {
	parser           *Parser
	exchanger        *RBACExchanger
	clientIDProvider ClientIDProvider
}

func NewPrincipalAuthorizer(
	parser *Parser,
	exchanger *RBACExchanger,
	clientIDProvider ClientIDProvider,
) *PrincipalAuthorizer {
	return &PrincipalAuthorizer{parser: parser, exchanger: exchanger, clientIDProvider: clientIDProvider}
}

type ClientIDProvider interface {
	FetchClientID(ctx context.Context, appID string) (string, error)
}

func (p *PrincipalAuthorizer) Handle(ctx context.Context, req appauth.AppUseRequest[principal.Token]) (appauth.Authorities, error) {
	clientID, err := p.clientIDProvider.FetchClientID(ctx, req.AppID)
	if err != nil {
		return nil, err
	}

	rbac, err := p.exchanger.Exchange(ctx, req.Token, req.ChCtx.RequiredAuthorities(), clientID)
	if err != nil {
		return nil, err
	}

	parsed, err := p.parser.Parse(ctx, rbac.AccessToken)
	if err != nil {
		return nil, err
	}

	if ok := parsed.CheckAction(Service(req.AppID), Action(req.FunctionName)); !ok {
		return nil, errors.New("action check fail")
	}

	return req.ChCtx.RequiredAuthorities(), nil
}
