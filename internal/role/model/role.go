package model

type AppRole struct {
	ID      string   `json:"id"`
	AppID   string   `json:"appId"`
	Type    RoleType `json:"type"`
	Version int      `json:"version"`
	RoleID  string   `json:"roleId"`
	*Credentials
}

type ChannelRoleAgreement struct {
	ChannelID string
	AppRoleID string
}

type Credentials struct {
	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
}

type Claim struct {
	Service string `json:"service"`
	Action  string `json:"action"`
}

func (c *Claim) Equal(other *Claim) bool {
	return c.Action == other.Action && c.Service == other.Service
}

type Claims []*Claim

type RoleType string

const (
	RoleTypeApp     = RoleType("app")
	RoleTypeChannel = RoleType("channel")
	RoleTypeUser    = RoleType("user")
	RoleTypeManager = RoleType("manager")
)

func (t RoleType) NeedsAgreement() bool {
	return t != RoleTypeApp
}

var AvailableRoleTypes = []RoleType{RoleTypeApp, RoleTypeChannel, RoleTypeUser, RoleTypeManager}
