package api

import "prototipo/internal/models"

type GetBillingDetais interface {
	GetVulnerable(id string) ([]models.BillingDetail, error)
}
