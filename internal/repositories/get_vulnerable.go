package repositories

import (
	"fmt"
	"prototipo/internal/models"
)

func (r *Repository) GetVulnerable(userID string) ([]models.BillingDetail, error) {

	query := fmt.Sprintf(`
		SELECT id, user_id, name, email, address, phone, credit_card
		FROM BILLING_DETAILS
		WHERE user_id = %s
	`, userID)

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("Error retrieving billing details: %v", err)
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
			return nil, fmt.Errorf("Error scanning: %v", err)
		}

		results = append(results, b)
	}

	if len(results) == 0 {
		return nil, fmt.Errorf("Billing details not found with User ID: %s", userID)
	}

	return results, nil
}
