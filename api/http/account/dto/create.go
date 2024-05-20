package dto

import (
	"unicode/utf8"

	"github.com/pkg/errors"
)

type AppCreateRequest struct {
	Title string `json:"title"`
}

func (r *AppCreateRequest) Validate() error {
	if utf8.RuneCountInString(r.Title) < 2 || utf8.RuneCountInString(r.Title) > 20 {
		return errors.New("title length should be between 2 and 20")
	}
	return nil
}
