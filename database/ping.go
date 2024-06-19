package database

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Ping(ctx context.Context, dsn string) error {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	defer db.Close()

	return db.PingContext(ctx)
}
