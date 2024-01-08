package domain

import (
	"time"

	"github.com/channel-io/ch-app-store/internal/resource/domain"
)

type Wam struct {
	ID string

	AppID string
	Name  string

	CreatedAt time.Time
	UpdatedAt time.Time
}

type WamRepository interface {
	domain.ResourceRepository[*Wam]
}
