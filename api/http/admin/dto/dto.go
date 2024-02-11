package dto

import (
	app "github.com/channel-io/ch-app-store/internal/app/domain"
	brief "github.com/channel-io/ch-app-store/internal/brief/domain"
	cmd "github.com/channel-io/ch-app-store/internal/command/domain"
)

type AppResource struct {
	App      *app.App
	Brief    *brief.Brief
	Commands []*cmd.Command
}

type AppResources struct {
	Apps     []*app.App
	Commands []*cmd.Command
	Briefs   []*brief.Brief
}
