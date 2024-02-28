package domain

import (
	"context"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/channel-io/ch-app-store/generated/models"
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
	appID string,
	req JsonFunctionRequest,
	res JsonFunctionResponse,
) {
	go func() {
		functionLog := &models.FunctionLog{
			AppID:      null.StringFrom(appID),
			Name:       null.StringFrom(req.Method),
			CallerType: null.StringFrom(req.Context.Caller.Type),
			CallerID:   null.StringFrom(req.Context.Caller.ID),
			IsSuccess:  null.BoolFrom(res.Error == nil),
		}
		_ = functionLog.Insert(context.Background(), f.db, boil.Infer())
	}()
}
