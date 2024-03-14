package invokelog

import (
	"context"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/channel-io/ch-app-store/generated/models"
	app "github.com/channel-io/ch-app-store/internal/app/domain"
	"github.com/channel-io/ch-app-store/lib/db"
)

type FunctionDBLogger struct {
	db db.DB
}

func NewFunctionDBLogger(db db.DB) *FunctionDBLogger {
	return &FunctionDBLogger{db: db}
}

func (f *FunctionDBLogger) OnInvoke(
	ctx context.Context,
	event app.FunctionInvokeEvent,
) {
	go func() {
		functionLog := &models.FunctionLog{
			AppID:      null.StringFrom(event.AppID),
			Name:       null.StringFrom(event.Request.Method),
			CallerType: null.StringFrom(event.Request.Context.Caller.Type),
			CallerID:   null.StringFrom(event.Request.Context.Caller.ID),
			IsSuccess:  null.BoolFrom(event.Response.Error == nil),
		}
		_ = functionLog.Insert(context.Background(), f.db, boil.Infer())
	}()
}
