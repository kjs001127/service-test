package dto

import (
	"github.com/channel-io/ch-app-store/api/http/shared/dto"
	"github.com/channel-io/ch-app-store/internal/app/domain"
)

type InstalledAppsWithCommands struct {
	domain.InstalledApps
	Commands []*dto.CommandDTO `json:"commands"`
}

type InstalledAppWithCommands struct {
	domain.InstalledApp
	Commands []*dto.CommandDTO `json:"commands"`
}
