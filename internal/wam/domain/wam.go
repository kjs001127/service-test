package domain

import (
	"time"
)

type Wam struct {
	ID string

	AppID string
	Name  string

	CreatedAt time.Time
	UpdatedAt time.Time
}
