package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upOrders, downOrders)
}

func upOrders(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	return nil
}

func downOrders(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
