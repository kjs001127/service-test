package ddbfx

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/lib/dynamodb"
)

var DynamoDB = fx.Options(
	fx.Provide(
		dynamodb.NewDynamoDB,
	),
)
