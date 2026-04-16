package repositories

import (
	"fmt"
)

func (r *Repository) Init() error {
	query := `
	CREATE TABLE IF NOT EXISTS BILLING_DETAILS (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id TEXT NOT NULL,
		name TEXT NOT NULL,
		email TEXT NOT NULL,
		address TEXT,
		phone TEXT,
		credit_card TEXT
	);
	`

	_, err := r.db.Exec(query)
	if err != nil {
		return fmt.Errorf("Error creating billing details table: %v", err)
	}

	return nil
}
