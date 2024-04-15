package model

type SystemLog struct {
	ID        string   `json:"id"`
	ChatType  ChatType `json:"chatType"`
	ChatId    string   `json:"chatId"`
	AppID     string   `json:"appId"`
	ChannelID string   `json:"channelId"`
	Message   string   `json:"message"`
	CreatedAt int64    `json:"createdAt"`
	ExpiresAt int64    `json:"expiresAt"`
}

type ChatType string

const ChatTypeUserChat = ChatType("userChat")
