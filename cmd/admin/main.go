package main

import (
	"github.com/channel-io/ch-app-store/cmd/adminfx"

	"go.uber.org/fx"
)

func main() {
	fx.New(adminfx.Admin)

	select {}
}
