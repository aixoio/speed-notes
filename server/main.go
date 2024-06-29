package main

import (
	"github.com/aixoio/speed-notes/server/db/postgresql"
	"github.com/aixoio/speed-notes/server/env"
	"github.com/aixoio/speed-notes/server/router"
)

func main() {
	cfg := env.LoadConfig()

	db, err := postgresql.Connect(cfg)
	if err != nil {
		panic(err)
	}

	router.StartRouter(cfg, db)
}
