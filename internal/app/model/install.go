package model

type AppInstallation struct {
	AppID     string `json:"appId"`
	ChannelID string `json:"channelId"`
}

func (i *AppInstallation) ID() InstallationID {
	return InstallationID{
		AppID:     i.AppID,
		ChannelID: i.ChannelID,
	}
}

type InstallationID struct {
	AppID     string `json:"appId"`
	ChannelID string `json:"channelId"`
}
