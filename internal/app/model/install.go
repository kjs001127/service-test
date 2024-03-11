package model

type AppChannel struct {
	AppID     string    `json:"appId"`
	ChannelID string    `json:"channelId"`
	Configs   ConfigMap `json:"configs"`
}

type ConfigMap map[string]string

type Install struct {
	AppID     string
	ChannelID string
}
