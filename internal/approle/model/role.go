package model

import (
	"github.com/channel-io/ch-proto/auth/v1/go/model"
)

type AppRole struct {
	*model.RoleCredentials
	RoleID string
	Type   RoleType
	AppID  string
}

type RoleType string
