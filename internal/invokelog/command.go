package invokelog

import (
	"context"
	"time"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/channel-io/ch-app-store/generated/models"
	command "github.com/channel-io/ch-app-store/internal/command/svc"
	"github.com/channel-io/ch-app-store/lib/db"
	"github.com/channel-io/ch-app-store/lib/log"
)

const logTimeout = 10 * time.Second

type CommandDBLogger struct {
	db     db.DB
	logger log.ContextAwareLogger
}

func NewCommandDBLogger(src db.DB, logger log.ContextAwareLogger) *CommandDBLogger {
	return &CommandDBLogger{db: src, logger: logger}
}

func (c *CommandDBLogger) OnInvoke(ctx context.Context, event command.CommandInvokeEvent) {
	go func() {
		reqCtx, cancel := context.WithTimeout(context.Background(), logTimeout)
		defer cancel()

		messageID := event.Request.Trigger.Attributes["messageId"]
		cmdLog := &models.CommandLog{
			AppID:            null.StringFrom(event.Request.AppID),
			ChannelID:        null.StringFrom(event.Request.ChannelID),
			CommandID:        null.StringFrom(event.ID),
			ChatType:         null.StringFrom(event.Request.Chat.Type),
			ChatID:           null.StringFrom(event.Request.Chat.ID),
			CallerType:       null.StringFrom(event.Request.Caller.Type),
			CallerID:         null.StringFrom(event.Request.Caller.ID),
			TriggerType:      null.StringFrom(event.Request.Trigger.Type),
			TriggerMessageID: null.NewString(messageID, messageID != ""),
			IsSuccess:        null.BoolFrom(event.Err == nil),
		}
		if err := cmdLog.Insert(reqCtx, c.db, boil.Infer()); err != nil {
			c.logger.Errorw(reqCtx, "command invoke log insertion fail", "log", cmdLog, "err", err)
		}
	}()
}
