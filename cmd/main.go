package main

import (
	"go.uber.org/fx"

	appstorefx "github.com/channel-io/ch-app-store/fx"
)

func main() {
	startModule()
}

func startModule() {
	fx.New(appstorefx.Option)

	select {}
}
