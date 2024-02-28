package main

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/fx/cmdfx/adminfx"
)

func main() {
	fx.New(adminfx.Admin)

	select {}
}
