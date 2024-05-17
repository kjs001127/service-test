package dto

import "github.com/pkg/errors"

type AppCreateRequest struct {
	Title string `json:"title"`
}

func (r *AppCreateRequest) Validate() error {
	if len(r.Title) < 2 || len(r.Title) > 20 {
		return errors.New("title length should be between 2 and 20")
	}
	return nil
}
