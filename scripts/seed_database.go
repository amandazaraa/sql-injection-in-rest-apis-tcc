package main

import (
	"log"
	"prototipo/internal/models"
	"prototipo/internal/repositories"
)

func main() {
	dbPath := "./database.db"
	repo, err := repositories.NewRepository(dbPath)
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	defer repo.Close()

	if err := repo.Init(); err != nil {
		log.Fatalf("Error creating table: %v", err)
	}

	billings := []models.BillingDetail{
		{
			UserID:     "1",
			Name:       "João Silva",
			Email:      "joao@example.com",
			Address:    "Rua A, 123",
			Phone:      "11999999999",
			CreditCard: "1234-5678-9012-3456",
		},
		{
			UserID:     "1",
			Name:       "João Silva",
			Email:      "joao@example.com",
			Address:    "Rua A, 123",
			Phone:      "11999999999",
			CreditCard: "5555-4444-3333-2222",
		},
		{
			UserID:     "1",
			Name:       "João Silva",
			Email:      "joao@example.com",
			Address:    "Rua A, 123",
			Phone:      "11999999999",
			CreditCard: "9999-5555-3333-2222",
		},
		{
			UserID:     "2",
			Name:       "Maria Santos",
			Email:      "maria@example.com",
			Address:    "Rua B, 456",
			Phone:      "11888888888",
			CreditCard: "9876-5432-1098-7654",
		},
		{
			UserID:     "2",
			Name:       "Maria Santos",
			Email:      "maria@example.com",
			Address:    "Rua Z, 123",
			Phone:      "11888888888",
			CreditCard: "1234-1234-1234-1234",
		},
		{
			UserID:     "3",
			Name:       "Pedro Oliveira",
			Email:      "pedro@example.com",
			Address:    "Rua C, 789",
			Phone:      "11777777777",
			CreditCard: "5555-4444-3333-2222",
		},
	}

	for _, billing := range billings {
		err := repo.Insert(billing)
		if err != nil {
			log.Printf("Error inserting record %s: %v", billing.Name, err)
			continue
		}
	}

	log.Println("Successfully registered!")
}
