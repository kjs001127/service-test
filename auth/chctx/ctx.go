package chctx

type ChannelContext struct {
	Channel struct {
		ID string `json:"id"`
	}
	Caller struct {
		Type string `json:"type"`
		ID   string `json:"id"`
	}
	Chat struct {
		Type string `json:"type"`
		ID   string `json:"id"`
	}
}
