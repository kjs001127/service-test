package dto

import (
	app "github.com/channel-io/ch-app-store/internal/app/domain"
)

type AppIDRequest struct {
	AppID string `json:"appId"`
}

type Apps struct {
	Apps []*app.App `json:"apps"`
}
