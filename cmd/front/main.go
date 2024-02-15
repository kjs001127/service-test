package main

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/fx/publicfx"
)

func main() {
	fx.New(publicfx.Public)

	select {}
}
