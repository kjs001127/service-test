package model

type Urls struct {
	FunctionURL *string `json:"functionUrl,omitempty"`
	WamURL      *string `json:"wamUrl,omitempty"`
	// SigningKey  *string `json:"signingKey,omitempty"`
}
