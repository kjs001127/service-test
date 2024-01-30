package domain

import (
	"context"
)

type BriefQuerySvc struct {
	repo BriefRepository
}

func NewQueryService(repo BriefRepository) *BriefQuerySvc {
	return &BriefQuerySvc{repo: repo}
}

func (s *BriefQuerySvc) QueryBrief(ctx context.Context, appID string) (*Brief, error) {
	return s.repo.Fetch(ctx, appID)
}
