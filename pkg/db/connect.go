package db

import (
	"database/sql"
	"log/slog"

	_ "github.com/lib/pq"
)

var (
	DB_DSN string
)

func init() {
	DB_DSN = "postgres://postgres:password@db:5432/flickgame?sslmode=disable"
}

func Connect() *sql.DB {
	db, err := sql.Open("postgres", DB_DSN)
	if err != nil {
		slog.Error("failed to connect database", "err=", err)
	}

	if err = db.Ping(); err != nil {
		slog.Error("failed to ping database", "err=", err)
	}

	return db
}
