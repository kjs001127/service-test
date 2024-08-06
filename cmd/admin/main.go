package main

import (
	"github.com/channel-io/ch-app-store/api/gintoolfx"
	"github.com/channel-io/ch-app-store/cmd/adminfx"
)

func main() {
	gintoolfx.StartServer(adminfx.Admin)
}
