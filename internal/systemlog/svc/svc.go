package svc

import (
	"context"
	"time"

	"github.com/channel-io/go-lib/pkg/uid"
	"github.com/pkg/errors"

	"github.com/channel-io/ch-app-store/internal/systemlog/model"
)

type Order string

const (
	OrderAsc  = Order("asc")
	OrderDesc = Order("desc")
)

const TTL = 24 * time.Hour * 7
const maxMessageSize = 2000

type SystemLogSvc struct {
	repo SystemLogRepository
}

func NewSystemLogSvc(repo SystemLogRepository) *SystemLogSvc {
	return &SystemLogSvc{repo: repo}
}

func (s *SystemLogSvc) SaveLog(ctx context.Context, log *model.SystemLog) (*model.SystemLog, error) {
	log.ID = uid.New().Hex()
	log.CreatedAt = time.Now().UnixMilli()
	log.ExpiresAt = time.Now().Add(TTL).Unix()

	if len(log.Message) >= maxMessageSize {
		log.Message = log.Message[:maxMessageSize]
	}

	if err := s.repo.Save(ctx, log); err != nil {
		return nil, errors.Wrap(err, "ddb save log")
	}

	return log, nil
}

func (s *SystemLogSvc) QueryLog(ctx context.Context, request *QueryRequest) ([]*model.SystemLog, error) {
	ret, err := s.repo.Query(ctx, request)
	if err != nil {
		return nil, errors.Wrap(err, "ddb query log")
	}

	return ret, nil
}

type QueryRequest struct {
	ChatId   string         `json:"chatId"`
	ChatType model.ChatType `json:"chatType"`
	CursorID string         `json:"cursorId"`
	Order    Order          `json:"order"`
	Limit    int32          `json:"limit"`
}
