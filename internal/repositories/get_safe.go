package repositories

import (
	"fmt"
	"prototipo/internal/models"
)

func (r *Repository) GetSafe(userID string) ([]models.BillingDetail, error) {

	query := `
		SELECT id, user_id, name, email, address, phone, credit_card
		FROM BILLING_DETAILS
		WHERE user_id = ?
	`

	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("Error retrieving billing details")
	}
	defer rows.Close()

	var results []models.BillingDetail

	for rows.Next() {
		var b models.BillingDetail

		err := rows.Scan(
			&b.ID,
			&b.UserID,
			&b.Name,
			&b.Email,
			&b.Address,
			&b.Phone,
			&b.CreditCard,
		)
		if err != nil {
			return nil, fmt.Errorf("Error scanning")
		}

		results = append(results, b)
	}

	if len(results) == 0 {
		return nil, fmt.Errorf("Record not found")
	}

	return results, nil
}
