package repositories

import (
	"fmt"
	"prototipo/internal/models"
)

func (r *Repository) Insert(b models.BillingDetail) error {
	query := `
	INSERT INTO BILLING_DETAILS (
		user_id,
		name,
		email,
		address,
		phone,
		credit_card
	) VALUES (?, ?, ?, ?, ?, ?)
	`

	_, err := r.db.Exec(
		query,
		b.UserID,
		b.Name,
		b.Email,
		b.Address,
		b.Phone,
		b.CreditCard,
	)

	if err != nil {
		return fmt.Errorf("Error saving billing details: %v", err)
	}

	return nil
}
