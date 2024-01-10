package domain

import (
	"context"

	"github.com/volatiletech/null/v8"
)

type QuerySvc struct {
}

type InstallQueryResult struct {
	Result  bool
	Message null.String
}

func (s *QuerySvc) CheckInstallable(ctx context.Context, info InstallInfo) (InstallQueryResult, error) {
	panic("")
}
