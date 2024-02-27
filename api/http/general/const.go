package general

import (
	"github.com/channel-io/ch-app-store/internal/auth/general"
)

// TODO @Camel Scope, Service const 정의 위치 고민
const (
	AppStoreService = general.Service("appstore.channel.io")

	ChannelScope = "channel"
	AppScope     = "app"
)
