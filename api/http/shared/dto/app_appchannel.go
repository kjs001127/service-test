package dto

import (
	app "github.com/channel-io/ch-app-store/internal/app/domain"
	appchannel "github.com/channel-io/ch-app-store/internal/appchannel/domain"
)

type AppAndAppChannel struct {
	App        app.App               `json:"app"`
	AppChannel appchannel.AppChannel `json:"appChannel"`
}
