package dto

import (
	app "github.com/channel-io/ch-app-store/internal/app/domain"
)

type JsonRPCRequest struct {
	ID      string             `json:"id"`
	JsonRPC string             `json:"jsonrpc"`
	Params  any                `json:"params"`
	Context app.ChannelContext `json:"context"`
}
