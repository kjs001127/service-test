package dto

import (
	rpc "github.com/channel-io/ch-app-store/internal/rpc/domain"
)

type ParamAndContext struct {
	Params  rpc.Params     `json:"params"`
	Context map[string]any `json:"context"` // 세부 정의 필요
}
