package domain

import (
	"context"
)

type Resource interface {
	GetID() string
	SetID(string)

	GetAppID() string
	SetAppID(string)

	GetName() string
	SetName(string)

	Validate() error
}

type Key struct {
	AppID string
	Name  string
}

type ResourceRepository[R Resource] interface {
	Fetch(ctx context.Context, key Key) (R, error)
	FetchAllByAppID(ctx context.Context, appID string) ([]R, error)

	Delete(ctx context.Context, key Key) error
	Save(ctx context.Context, resource R) (R, error)
}
