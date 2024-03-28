package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upAddresses, downAddresses)
}

func upAddresses(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	return nil
}

func downAddresses(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
