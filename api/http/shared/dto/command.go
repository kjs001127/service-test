package dto

import (
	app "github.com/channel-io/ch-app-store/internal/app/domain"
	cmd "github.com/channel-io/ch-app-store/internal/command/domain"
)

type ArgumentsAndContext struct {
	Arguments cmd.Arguments  `json:"arguments"`
	Context   map[string]any `json:"context"` // 세부 정의 필요
}

type AppsAndCommands struct {
	Apps     []*app.App     `json:"apps"`
	Commands []*cmd.Command `json:"commands"`
}

type ContextAndAutoCompleteArgs struct {
	Context map[string]any       `json:"context"`
	Params  cmd.AutoCompleteArgs `json:"params"`
}

type Context struct {
}

type Channel struct {
}
