package domain

import (
	"context"
	"fmt"
)

type QueryService struct {
	repo CommandRepository
}

func NewQueryService(repo CommandRepository) *QueryService {
	return &QueryService{repo: repo}
}

func (s *QueryService) QueryCommands(ctx context.Context, query Query) ([]*Command, error) {
	cmds, err := s.repo.FetchByQuery(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("fetchByQuery failed. cause: %w", err)
	}

	return cmds, nil
}
