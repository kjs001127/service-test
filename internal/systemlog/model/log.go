package model

type SystemLog struct {
	Id        string `json:"id"`
	ChatType  string `json:"chatType"`
	ChatId    string `json:"chatId"`
	AppID     string `json:"appId"`
	ChannelID string `json:"channelId"`
	Message   string `json:"message"`
	CreatedAt int64  `json:"createdAt"`
	ExpiresAt int64  `json:"expiresAt"`
}
