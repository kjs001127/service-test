package model

type Channel struct {
	Id              string                      `json:"id"`
	Name            string                      `json:"name,omitempty"`
	Description     string                      `json:"description,omitempty"`
	NameDescI18NMap map[string]*ChannelNameDesc `json:"nameDescI18nMap,omitempty"`
	AvatarUrl       string                      `json:"avatarUrl,omitempty"`
}

type ChannelNameDesc struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}
