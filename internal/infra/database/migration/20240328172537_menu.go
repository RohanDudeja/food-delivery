package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upMenu, downMenu)
}

func upMenu(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	return nil
}

func downMenu(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
