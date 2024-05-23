package model

type ServerSetting struct {
	FunctionURL *string    `json:"functionUrl,omitempty"`
	WamURL      *string    `json:"wamUrl,omitempty"`
	SigningKey  *string    `json:"signingKey,omitempty"`
	AccessType  AccessType `json:"accessType"`
}

type AccessType string

const (
	AccessType_Internal = AccessType("internal")
	AccessType_External = AccessType("external")
)
