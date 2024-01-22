package domain

import (
	"context"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
	"github.com/friendsofgo/errors"

	rpc "github.com/channel-io/ch-app-store/internal/rpc/domain"
)

type CommandRequest struct {
	Key
	Arguments Arguments
	Context   map[string]any
}

type CmdHandlerFactory struct {
	repo CommandRepository

	paramValidator  *ArgsValidator
	actionValidator *ActionValidator
}

func NewCmdHandlerFactory(
	repo CommandRepository,
	argsValidator *ArgsValidator,
	actionValidator *ActionValidator,
) *CmdHandlerFactory {
	return &CmdHandlerFactory{repo: repo, paramValidator: argsValidator, actionValidator: actionValidator}
}

func (c *CmdHandlerFactory) NewHandler(ctx context.Context, req CommandRequest) (*CmdHandler, error) {
	cmd, err := c.repo.Fetch(ctx, req.Key)
	if err != nil {
		return nil, errors.Wrap(err, "fetch cmd fail")
	}

	return &CmdHandler{
		target:          cmd,
		paramValidator:  c.paramValidator,
		actionValidator: c.actionValidator,
	}, nil
}

type CmdHandler struct {
	target *Command

	paramValidator  *ArgsValidator
	actionValidator *ActionValidator
}

func (v *CmdHandler) Handle(ctx context.Context, requester rpc.Requester, req CommandRequest) (Action, error) {
	if err := v.paramValidator.ValidateArgs(v.target.ParamDefinitions, req.Arguments); err != nil {
		return Action{}, apierr.BadRequest(err)
	}

	ret := new(Action)
	if err := requester.SendRequest(ctx, v.toRpcRequest(req), ret); err != nil {
		return Action{}, nil
	}

	if err := v.actionValidator.ValidateAction(*ret); err != nil {
		return Action{}, apierr.BadRequest(err)
	}

	return *ret, nil
}

func (v *CmdHandler) toRpcRequest(req CommandRequest) rpc.RpcRequest {
	return rpc.RpcRequest{
		AppID:        v.target.AppID,
		FunctionName: v.target.FunctionName,
		Arguments:    req.Arguments,
		ExtraFields:  withContext(req.Context),
	}
}

type AutoCompleteRequest struct {
	Key
	Context   map[string]any
	Arguments AutoCompleteArgs
}

type AutoCompleteHandlerFactory struct {
	repo CommandRepository
}

func NewAutoCompleteHandlerFactory(repo CommandRepository) *AutoCompleteHandlerFactory {
	return &AutoCompleteHandlerFactory{repo: repo}
}

func (f *AutoCompleteHandlerFactory) NewHandler(ctx context.Context, req AutoCompleteRequest) (*AutoCompleteHandler, error) {
	cmd, err := f.repo.Fetch(ctx, req.Key)
	if err != nil {
		return nil, errors.Wrap(err, "fetch cmd fail")
	}

	return &AutoCompleteHandler{
		target: cmd,
	}, nil
}

type AutoCompleteHandler struct {
	target *Command
}

func (v *AutoCompleteHandler) Handle(
	ctx context.Context,
	requester rpc.Requester,
	req AutoCompleteRequest,
) (Choices, error) {
	if err := req.Arguments.validate(); err != nil {
		return nil, errors.Wrap(err, "autoComplete arguments are not valid")
	}

	if !v.target.AutoCompleteFunctionName.Valid {
		return nil, errors.New("autoCompleteFunctionName is empty")
	}

	ret := new(Choices)
	if err := requester.SendRequest(ctx, v.toRpcRequest(req), ret); err != nil {
		return nil, err
	}

	if err := ret.validate(); err != nil {
		return nil, apierr.BadRequest(err)
	}

	return *ret, nil
}

func (v *AutoCompleteHandler) toRpcRequest(req AutoCompleteRequest) rpc.RpcRequest {
	return rpc.RpcRequest{
		AppID:        v.target.AppID,
		FunctionName: v.target.AutoCompleteFunctionName.String,
		Arguments:    req.Arguments,
		ExtraFields:  withContext(req.Context),
	}
}

const contextHeaderKey = "context"

func withContext(context map[string]any) map[string]any {
	return map[string]any{contextHeaderKey: context}
}
