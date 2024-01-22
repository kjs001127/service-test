package dto

import (
	cmd "github.com/channel-io/ch-app-store/internal/command/domain"
)

type ArgumentsAndContext struct {
	Arguments cmd.Arguments  `json:"arguments"`
	Context   map[string]any `json:"context"` // 세부 정의 필요
}

type Context struct {
}

type Channel struct {
}
