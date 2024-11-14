package dto

import (
	"github.com/channel-io/ch-app-store/api/http/desk/dto"
)

type AppStoreView struct {
	App      *dto.DetailedAppView     `json:"app"`
	Commands []*dto.SimpleCommandView `json:"commands"`
}
