package models

import (
	"database/sql"
	"embed"
	"github.com/pressly/goose/v3"
)

//go:embed sql/*.sql
var EmbedMigrations embed.FS

func DBMigrations(db *sql.DB) error {

	goose.SetBaseFS(EmbedMigrations)

	if err := goose.SetDialect("mysql"); err != nil {
		return err
	}
	if err := goose.Up(db, "sql"); err != nil {
		return err
	}
	if err := goose.Status(db, "sql"); err != nil {
		return err
	}

	return nil
}
