package event

import (
	"context"
)

type RequestListener[REQ any, RES any] interface {
	OnInvoke(ctx context.Context, appID string, req REQ, res RES)
}
