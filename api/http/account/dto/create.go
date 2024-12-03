package dto

import (
	"unicode/utf8"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
	"github.com/pkg/errors"
)

type AppCreateRequest struct {
	Title string `json:"title"`
}

func (r *AppCreateRequest) Validate() error {
	if utf8.RuneCountInString(r.Title) < 2 || utf8.RuneCountInString(r.Title) > 20 {
		return apierr.UnprocessableEntity(errors.New("{app.title.invalid.length}"))
	}
	return nil
}
