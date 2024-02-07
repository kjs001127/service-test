package dto

import (
	"github.com/channel-io/ch-app-store/auth/appauth"
)

type JsonRPCRequest struct {
	ID      string                 `json:"id"`
	JsonRPC string                 `json:"jsonrpc"`
	Params  any                    `json:"params"`
	Context appauth.ChannelContext `json:"context"`
}
