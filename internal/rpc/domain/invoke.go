package domain

import (
	"context"
)

type InvokeSvc[REQ any, RET any] interface {
	Invoke(ctx context.Context, originReq REQ) (RET, error)
}
