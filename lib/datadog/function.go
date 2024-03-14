package datadog

import (
	"context"

	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/ext"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"

	app "github.com/channel-io/ch-app-store/internal/app/domain"
	"github.com/channel-io/ch-app-store/lib/log"
)

type MethodSpanTagger struct {
	logger log.ContextAwareLogger
}

func NewMethodSpanTagger(logger log.ContextAwareLogger) *MethodSpanTagger {
	return &MethodSpanTagger{logger: logger}
}

func (d *MethodSpanTagger) OnInvoke(ctx context.Context, event app.FunctionInvokeEvent) {
	span, ok := tracer.SpanFromContext(ctx)
	if !ok {
		d.logger.Warn(ctx, "span not found on MethodSpanTagger")
	}

	span.SetTag(ext.RPCService, event.AppID)
	span.SetTag(ext.RPCMethod, event.Request.Method)
}
