package dto

import (
	appmodel "github.com/channel-io/ch-app-store/internal/app/model"
	app "github.com/channel-io/ch-app-store/internal/app/svc"
	cmd "github.com/channel-io/ch-app-store/internal/command/model"
)

type InstalledApp struct {
	App        *appmodel.App        `json:"app"`
	AppChannel *appmodel.AppChannel `json:"appChannel"`
}

type AppsAndFullCommands struct {
	Apps     []*appmodel.App `json:"apps,omitempty"`
	Commands []*cmd.Command  `json:"commands,omitempty"`
}

type RegisterRequest struct {
	Commands []*cmd.Command `json:"commands"`
}

type BriefRequest struct {
	Context app.ChannelContext `json:"context"`
}
