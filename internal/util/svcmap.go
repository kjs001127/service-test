package util

type Host string

func (h Host) String() string {
	return string(h)
}

type ServiceKey string

func (s ServiceKey) String() string {
	return string(s)
}

const (
	CHANNEL_API       = ServiceKey("channel-api")
	DOCUMENT_API      = ServiceKey("document-api")
	CHANNEL_ADMIN_API = ServiceKey("channel-admin-api")
	APP_STORE_API     = ServiceKey("app-store-api")
)

type ServiceMap map[ServiceKey]Host
