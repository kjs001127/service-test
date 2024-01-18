package domain

import (
	"context"
	"fmt"
)

type QuerySvc struct {
	repo CommandRepository
}

func NewQueryService(repo CommandRepository) *QuerySvc {
	return &QuerySvc{repo: repo}
}

func (s *QuerySvc) QueryCommands(ctx context.Context, query Query) ([]*Command, error) {
	cmds, err := s.repo.FetchByQuery(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("fetchByQuery failed. cause: %w", err)
	}

	return cmds, nil
}
