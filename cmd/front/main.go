package main

import (
	"github.com/channel-io/ch-app-store/cmd/publicfx"

	"go.uber.org/fx"
)

func main() {
	fx.New(publicfx.Public)

	select {}
}
