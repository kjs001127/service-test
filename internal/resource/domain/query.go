package domain

import (
	"context"
)

type QueryService[R Resource] struct {
	repo ResourceRepository[R]
}

func (q *QueryService[R]) Query(ctx context.Context, key Key) (R, error) {
	return q.repo.Fetch(ctx, key)
}
