package domain

import (
	"context"
)

type CheckType string

const (
	CheckTypeConfig = CheckType("config")
)

type CheckRequest struct {
	ChannelId string
	Type      CheckType
	Data      any
}

type CheckReturn struct {
	Success  bool
	Messages map[string]any
}

type HttpAppChecker interface {
	Request(ctx context.Context, url string, req CheckRequest) (CheckReturn, error)
}
