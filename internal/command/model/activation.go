package model

import (
	appmodel "github.com/channel-io/ch-app-store/internal/app/model"
)

type Activation struct {
	appmodel.InstallationID
	Enabled bool
}

type ActivationSetting struct {
	AppID           string
	EnableByDefault bool
}
