package model

type AppInstallation struct {
	AppID     string `json:"appId"`
	ChannelID string `json:"channelId"`
}

type InstallationID struct {
	AppID     string
	ChannelID string
}
