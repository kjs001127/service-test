package domain

import (
	"context"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/channel-io/ch-app-store/generated/models"
	"github.com/channel-io/ch-app-store/internal/app/domain"
	"github.com/channel-io/ch-app-store/lib/db"
)

type CommandDBLogger struct {
	db db.DB
}

func NewCommandDBLogger(src db.DB) *CommandDBLogger {
	return &CommandDBLogger{db: src}
}

func (c *CommandDBLogger) OnInvoke(ctx context.Context, cmdID string, req CommandRequest, res domain.TypedResponse[Action]) {
	cmdLog := &models.CommandLog{
		AppID:       null.StringFrom(req.AppID),
		ChannelID:   null.StringFrom(req.Channel.ID),
		CommandID:   null.StringFrom(cmdID),
		ChatType:    null.StringFrom(req.Chat.Type),
		ChatID:      null.StringFrom(req.Chat.ID),
		CallerType:  null.StringFrom(req.Caller.Type),
		CallerID:    null.StringFrom(req.Caller.ID),
		TriggerType: null.StringFrom(req.Trigger.Type),
		IsSuccess:   null.BoolFrom(res.Error != nil),
	}
	_ = cmdLog.Insert(ctx, c.db, boil.Infer())
}
