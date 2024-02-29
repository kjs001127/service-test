package log

import (
	"context"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/channel-io/ch-app-store/generated/models"
	command "github.com/channel-io/ch-app-store/internal/command/domain"
	"github.com/channel-io/ch-app-store/lib/db"
)

type CommandDBLogger struct {
	db db.DB
}

func NewCommandDBLogger(src db.DB) *CommandDBLogger {
	return &CommandDBLogger{db: src}
}

func (c *CommandDBLogger) OnInvoke(ctx context.Context, event command.CommandInvokeEvent) {
	go func() {
		messageID := event.Request.Trigger.Attributes["messageId"]
		cmdLog := &models.CommandLog{
			AppID:            null.StringFrom(event.Request.AppID),
			ChannelID:        null.StringFrom(event.Request.Caller.ChannelID),
			CommandID:        null.StringFrom(event.ID),
			ChatType:         null.StringFrom(event.Request.Chat.Type),
			ChatID:           null.StringFrom(event.Request.Chat.ID),
			CallerType:       null.StringFrom(event.Request.Caller.Type),
			CallerID:         null.StringFrom(event.Request.Caller.ID),
			TriggerType:      null.StringFrom(event.Request.Trigger.Type),
			TriggerMessageID: null.NewString(messageID, messageID != ""),
			IsSuccess:        null.BoolFrom(event.Err == nil),
		}
		_ = cmdLog.Insert(context.Background(), c.db, boil.Infer())
	}()
}
