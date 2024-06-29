package main

import (
	"github.com/aixoio/speed-notes/server/env"
	"github.com/aixoio/speed-notes/server/router"
)

func main() {
	cfg := env.LoadConfig()

	router.StartRouter(cfg)
}
