package systemlog

import (
	"context"
	"encoding/json"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
	"github.com/pkg/errors"

	authgen "github.com/channel-io/ch-app-store/internal/auth/general"
	"github.com/channel-io/ch-app-store/internal/native"
	"github.com/channel-io/ch-app-store/internal/systemlog/model"
	"github.com/channel-io/ch-app-store/internal/systemlog/svc"
)

type SystemLog struct {
	svc        *svc.SystemLogSvc
	rbacParser authgen.Parser
}

func NewSystemLog(svc *svc.SystemLogSvc, rbacParser authgen.Parser) *SystemLog {
	return &SystemLog{svc: svc, rbacParser: rbacParser}
}

func (s *SystemLog) RegisterTo(registry native.FunctionRegistry) {
	registry.Register("writeSystemLog", s.WriteLog)
}

func (s *SystemLog) WriteLog(ctx context.Context, token native.Token, req native.FunctionRequest) native.FunctionResponse {
	var log model.SystemLog
	if err := json.Unmarshal(req.Params, &log); err != nil {
		return native.WrapCommonErr(err)
	}

	if err := s.authorize(ctx, token, &log); err != nil {
		return native.WrapCommonErr(err)
	}

	logWritten, err := s.svc.SaveLog(ctx, &log)
	if err != nil {
		return native.WrapCommonErr(err)
	}

	res, err := json.Marshal(logWritten)
	if err != nil {
		return native.WrapCommonErr(err)
	}

	return native.ResultSuccess(res)
}

const (
	channelScope = "channel"
	appScope     = "app"
)

func (s *SystemLog) authorize(ctx context.Context, token native.Token, log *model.SystemLog) error {
	parsedRbac, err := s.rbacParser.Parse(ctx, token.Value)
	if err != nil {
		return err
	}

	if !parsedRbac.CheckScopes(authgen.Scopes{
		channelScope: {log.ChannelID},
		appScope:     {log.AppID},
	}) {
		return apierr.Unauthorized(errors.New("scope check fail"))
	}

	return nil
}
