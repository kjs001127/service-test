package dto

import (
	app "github.com/channel-io/ch-app-store/internal/app/domain"
	cmd "github.com/channel-io/ch-app-store/internal/command/domain"
)

type AppsAndFullCommands struct {
	Apps     []*app.App     `json:"apps"`
	Commands []*cmd.Command `json:"commands"`
}

type RegisterRequest struct {
	Commands []*cmd.Command `json:"commands"`
}

type BriefRequest struct {
	Context app.ChannelContext `json:"context"`
}
