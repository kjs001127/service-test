package domain

import (
	"context"

	rpc "github.com/channel-io/ch-app-store/internal/rpc/domain"
)

// RpcRepository is an Adapter between CommandRepository and rpc.domain.RpcRepository
type RpcRepository struct {
	repo            CommandRepository
	resultValidator *ActionValidator
}

func NewRpcRepository(repo CommandRepository, resultValidator *ActionValidator) *RpcRepository {
	return &RpcRepository{repo: repo, resultValidator: resultValidator}
}

func (r *RpcRepository) Fetch(ctx context.Context, key rpc.Key) (rpc.Rpc, error) {
	cmd, err := r.repo.Fetch(ctx, Key{AppID: key.AppID, Name: key.Name})
	if err != nil {
		return nil, err
	}

	return &CommandValidator{
		ID:              cmd.ID,
		ParamValidator:  cmd.ParamDefinitions,
		ResultValidator: r.resultValidator,
	}, nil
}

// CommandValidator is a Command specific implementation of rpc.domain.Rpc
type CommandValidator struct {
	ID              string
	ParamValidator  ParamDefinitions
	ResultValidator *ActionValidator
}

func (c *CommandValidator) ValidateParams(params rpc.Params) error {
	return c.ParamValidator.validateParamInput(params)
}

func (c *CommandValidator) ValidateResult(res rpc.Result) error {
	return c.ResultValidator.validate(res)
}
