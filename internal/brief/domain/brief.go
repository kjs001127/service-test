package domain

import (
	"context"
)

type Brief struct {
	ID                string
	AppID             string
	BriefFunctionName string
}

type BriefRepository interface {
	Fetch(ctx context.Context, appID string) (*Brief, error)
	DeleteAllByAppID(ctx context.Context, appID string) error
	FetchAll(ctx context.Context, appIDs []string) ([]*Brief, error)
}
