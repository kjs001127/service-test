package dto

import (
	"github.com/channel-io/ch-app-store/internal/app/model"
)

type AppCreateRequest struct {
	Title string `json:"title"`
}

func (r *AppCreateRequest) ConvertToApp() *model.App {
	return &model.App{
		Title: r.Title,
	}
}
