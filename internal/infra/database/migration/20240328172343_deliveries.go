package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upDeliveries, downDeliveries)
}

func upDeliveries(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	return nil
}

func downDeliveries(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
