package domain

import (
	"context"
)

type WamSvc struct {
}

type WamKey struct {
	AppID string
	Name  string
}

func (s *WamSvc) GetWamUrl(ctx context.Context, key WamKey) string {
	return "" // cloudfront URL
}

func (s *WamSvc) UpdateWam(ctx context.Context, key WamKey) error {
	// invalidate cloudfront cache
	return nil
}
