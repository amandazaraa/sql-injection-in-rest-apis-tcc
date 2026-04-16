package models

type BillingDetail struct {
	ID         int    `json:"id"`
	UserID     string `json:"user_id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Address    string `json:"address"`
	Phone      string `json:"phone"`
	CreditCard string `json:"credit_card"`
}
