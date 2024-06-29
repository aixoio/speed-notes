package postgresql

import (
	"database/sql"

	"github.com/aixoio/speed-notes/server/env"
)

func Connect(cfg env.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.Postgresql_url)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
