package main

import (
	"go.uber.org/fx"

	appstorefx "github.com/channel-io/ch-app-store/fx"
)

func main() {
	fx.New(appstorefx.Option)

	select {}
}
