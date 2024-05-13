package model

type AppRole struct {
	Credentials *Credentials
	RoleID      string
	Type        RoleType
	AppID       string
}

type Credentials struct {
	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
}

type Claim struct {
	Service string `json:"service"`
	Action  string `json:"action"`
}

type RoleType string

const (
	RoleTypeApp     = RoleType("app")
	RoleTypeChannel = RoleType("channel")
	RoleTypeUser    = RoleType("user")
	RoleTypeManager = RoleType("manager")
)

var AvailableRoleTypes = []RoleType{RoleTypeApp, RoleTypeChannel, RoleTypeUser, RoleTypeManager}
