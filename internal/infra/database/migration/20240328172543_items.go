package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upItems, downItems)
}

func upItems(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	return nil
}

func downItems(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
