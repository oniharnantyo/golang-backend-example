package migration

import (
	"database/sql"
	"fmt"

	migrate "github.com/rubenv/sql-migrate"
)

func Up(db *sql.DB) error {
	migrations := &migrate.FileMigrationSource{
		Dir: "database/migration/files",
	}

	n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		return err
	}

	fmt.Printf("Applied %d migrations!\n", n)

	return nil

}
