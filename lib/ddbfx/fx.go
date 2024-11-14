package ddbfx

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/lib/ddb"
)

var DynamoDB = fx.Options(
	fx.Provide(
		ddb.NewDynamoDB,
	),
)
