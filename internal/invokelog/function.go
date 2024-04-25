package invokelog

import (
	"context"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/channel-io/ch-app-store/generated/models"
	app "github.com/channel-io/ch-app-store/internal/app/svc"
	"github.com/channel-io/ch-app-store/lib/db"
	"github.com/channel-io/ch-app-store/lib/log"
)

type FunctionDBLogger struct {
	db     db.DB
	logger log.ContextAwareLogger
}

func NewFunctionDBLogger(db db.DB, logger log.ContextAwareLogger) *FunctionDBLogger {
	return &FunctionDBLogger{db: db, logger: logger}
}

func (f *FunctionDBLogger) OnInvoke(
	ctx context.Context,
	event app.FunctionInvokeEvent,
) {
	go func() {
		reqCtx, cancel := context.WithTimeout(context.Background(), logTimeout)
		defer cancel()

		functionLog := &models.FunctionLog{
			AppID:      null.StringFrom(event.AppID),
			Name:       null.StringFrom(event.Request.Method),
			CallerType: null.StringFrom(string(event.Request.Context.Caller.Type)),
			CallerID:   null.StringFrom(event.Request.Context.Caller.ID),
			IsSuccess:  null.BoolFrom(event.Response.Error == nil),
		}
		if err := functionLog.Insert(reqCtx, f.db, boil.Infer()); err != nil {
			f.logger.Errorw(reqCtx, "function invoke log insertion fail", "log", functionLog, "err", err)
		}
	}()
}
