package main

import (
	"github.com/vinc-dev/promo-strategy/transport/container"
	httpServer "github.com/vinc-dev/promo-strategy/transport/http/server"
)

func main() {
	http := httpServer.CreateHttpServer(container.CreateAppContainer())
	http.Serve()
}
