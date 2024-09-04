package command

import (
	"context"
	"encoding/json"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
	"github.com/pkg/errors"

	authgen "github.com/channel-io/ch-app-store/internal/auth/general"
	"github.com/channel-io/ch-app-store/internal/command/model"
	command "github.com/channel-io/ch-app-store/internal/command/svc"
	"github.com/channel-io/ch-app-store/internal/native"
	"github.com/channel-io/ch-app-store/internal/native/localapi/command/action/public"
)

type RegisterRequest struct {
	AppID    string  `json:"appId"`
	Commands CmdDTOs `json:"commands"`
}

type CmdDTOs []CmdDTO
type CmdDTO struct {
	Name  string      `json:"name"`
	Scope model.Scope `json:"scope"`

	Description     *string                  `json:"description"`
	NameDescI18NMap map[string]model.I18nMap `json:"nameDescI18nMap"`

	ActionFunctionName       string  `json:"actionFunctionName"`
	AutoCompleteFunctionName *string `json:"autoCompleteFunctionName"`

	AlfMode        model.AlfMode `json:"alfMode,omitempty"`
	AlfDescription *string       `json:"alfDescription,omitempty"`

	ParamDefinitions model.ParamDefinitions `json:"paramDefinitions"`

	EnabledByDefault *bool `json:"enabledByDefault"`
}

func (d *CmdDTO) toCmd() *model.Command {
	if d.EnabledByDefault == nil {
		defaultEnabled := true
		d.EnabledByDefault = &defaultEnabled
	}
	return &model.Command{
		Name:                     d.Name,
		Scope:                    d.Scope,
		Description:              d.Description,
		NameDescI18NMap:          d.NameDescI18NMap,
		ActionFunctionName:       d.ActionFunctionName,
		AutoCompleteFunctionName: d.AutoCompleteFunctionName,
		ParamDefinitions:         d.ParamDefinitions,
		AlfMode:                  d.AlfMode,
		AlfDescription:           d.AlfDescription,
		EnabledByDefault:         *d.EnabledByDefault,
	}
}

func (d CmdDTOs) toCmds() []*model.Command {
	ret := make([]*model.Command, 0, len(d))
	for _, cmd := range d {
		ret = append(ret, cmd.toCmd())
	}
	return ret
}

func (r *Handler) RegisterCommand(
	ctx context.Context,
	token native.Token,
	request native.FunctionRequest,
) native.FunctionResponse {
	var req RegisterRequest
	if err := json.Unmarshal(request.Params, &req); err != nil {
		return native.WrapCommonErr(err)
	}

	if err := r.authorizeReg(ctx, token, &req); err != nil {
		return native.WrapCommonErr(err)
	}

	if err := r.registerSvc.Register(ctx, &command.CommandRegisterRequest{
		AppID:    req.AppID,
		Commands: req.Commands.toCmds(),
	}); err != nil {
		return native.WrapCommonErr(err)
	}

	return native.Empty()
}

const (
	appScope = "app"
)

func (r *Handler) authorizeReg(ctx context.Context, token native.Token, req *RegisterRequest) error {
	parsedRbac, err := r.rbacParser.Parse(ctx, token.Value)
	if err != nil {
		return err
	}

	if !parsedRbac.CheckAction(authgen.Service(r.serviceName), public.RegisterCommands) {
		return apierr.Unauthorized(errors.New("service, action check fail"))
	}

	if !parsedRbac.CheckScopes(authgen.Scopes{
		appScope: {req.AppID},
	}) {
		return apierr.Unauthorized(errors.New("scope check fail"))
	}

	return nil
}
