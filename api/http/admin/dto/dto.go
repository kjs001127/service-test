package dto

import (
	app "github.com/channel-io/ch-app-store/internal/app/domain"
	cmd "github.com/channel-io/ch-app-store/internal/command/domain"
)

type AppsAndFullCommands struct {
	Apps     []*app.App     `json:"apps"`
	Commands []*cmd.Command `json:"commands"`
}
