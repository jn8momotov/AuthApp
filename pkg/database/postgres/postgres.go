package postgres

import (
	"AuthApp/internal/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Init(config config.Postgres) (*sqlx.DB, error) {
	db, error := sqlx.Open("postgres", config.DataSource())
	if error != nil {
		return nil, error
	}

	error = db.Ping()
	if error != nil {
		return nil, error
	}

	return db, nil
}
