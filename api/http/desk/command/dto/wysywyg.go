package dto

import (
	"github.com/channel-io/ch-app-store/api/http/desk/dto"
)

type WysiwygView struct {
	Apps     []*dto.SimpleAppView     `json:"apps"`
	Commands []*dto.SimpleCommandView `json:"commands"`
}
