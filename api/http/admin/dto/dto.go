package dto

import (
	appmodel "github.com/channel-io/ch-app-store/internal/app/model"
	app "github.com/channel-io/ch-app-store/internal/app/svc"
	cmd "github.com/channel-io/ch-app-store/internal/command/model"
	accountmodel "github.com/channel-io/ch-app-store/internal/permission/model"
)

type InstalledApp struct {
	App             *appmodel.App             `json:"app"`
	AppInstallation *appmodel.AppInstallation `json:"appChannel"`
}

type AppsAndFullCommands struct {
	Apps     []*appmodel.App `json:"apps"`
	Commands []*cmd.Command  `json:"commands"`
}

type RegisterRequest struct {
	EnableByDefault    bool           `json:"enableByDefault"`
	ToggleFunctionName *string        `json:"toggleFunctionName,omitempty"`
	Commands           []*cmd.Command `json:"commands"`
}

type BriefRequest struct {
	Context app.ChannelContext `json:"context"`
}

type ChannelResponse struct {
	Channels []*accountmodel.Channel `json:"channels"`
}
