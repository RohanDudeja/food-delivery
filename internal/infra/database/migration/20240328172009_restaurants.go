package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upRestaurants, downRestaurants)
}

func upRestaurants(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	return nil
}

func downRestaurants(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
