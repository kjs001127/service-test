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

type CommandView struct {
	AppID        string
	FunctionName string

	Name        string
	Description string
}

func (s *QueryService) QueryCommands(ctx context.Context, query Query) ([]*CommandView, error) {
	cmds, err := s.repo.FetchByQuery(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("fetchByQuery failed. cause: %w", err)
	}

	views := make([]*CommandView, len(cmds))
	for _, cmd := range cmds {
		views = append(views, s.viewOf(cmd))
	}

	return views, nil
}

func (s *QueryService) viewOf(cmd *Command) *CommandView {
	return &CommandView{AppID: cmd.AppID, Name: cmd.Name, FunctionName: cmd.FunctionName, Description: cmd.Description}
}
