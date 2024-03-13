package model

type Installation struct {
	AppID     string    `json:"appId"`
	ChannelID string    `json:"channelId"`
	Configs   ConfigMap `json:"configs"`
}

type ConfigMap map[string]string

type InstallationID struct {
	AppID     string
	ChannelID string
}
