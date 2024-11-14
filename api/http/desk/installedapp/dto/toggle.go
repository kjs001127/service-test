package dto

import (
	cmd "github.com/channel-io/ch-app-store/internal/command/model"
)

type CommandToggleRequest struct {
	Scope    cmd.Scope `json:"scope"`
	Name     string    `json:"name"`
	Enabled  bool      `json:"enabled"`
	Language string    `json:"language"`
}
