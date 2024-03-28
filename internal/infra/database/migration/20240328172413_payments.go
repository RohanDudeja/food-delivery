package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upPayments, downPayments)
}

func upPayments(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	return nil
}

func downPayments(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
