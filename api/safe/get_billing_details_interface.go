package segura

import "prototipo/internal/models"

type GetBillingDetails interface {
	GetSafe(id string) ([]models.BillingDetail, error)
}
