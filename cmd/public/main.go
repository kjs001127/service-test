package main

import (
	"github.com/channel-io/ch-app-store/api/gintoolfx"
	"github.com/channel-io/ch-app-store/cmd/publicfx"
)

func main() {
	gintoolfx.StartServer(publicfx.Public)
}
