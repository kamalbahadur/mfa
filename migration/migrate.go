package migration

import (
	"database/sql"
	"embed"
	"github.com/pressly/goose/v3"
)

//go:embed ddl/*.sql
//go:embed dml/*.sql
var migrations embed.FS

func MigrateTo(db *sql.DB, dialect string) error {
	goose.SetBaseFS(migrations)
	err := goose.SetDialect(dialect)
	if err != nil {
		return err
	}
	err = goose.Up(db, "ddl")
	if err != nil {
		return err
	}
	err = goose.Up(db, "dml", goose.WithNoVersioning())
	if err != nil {
		return err
	}
	return nil
}
