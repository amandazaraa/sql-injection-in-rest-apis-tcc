package main

import (
	"log"
	safe "prototipo/api/safe"
	vulnerable "prototipo/api/vulnerable"
	"prototipo/internal/repositories"

	"github.com/gin-gonic/gin"
)

func main() {
	repository, err := repositories.NewRepository("./database.db")
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	defer repository.Close()

	handlerVulnerable := vulnerable.NewHandler(repository)
	handlerSafe := safe.NewHandler(repository)

	r := gin.Default()
	r.GET("/billing/vulnerable/get", handlerVulnerable.GetBillingDetails)
	r.GET("/billing/safe/get", handlerSafe.GetBillingDetails)
	r.Run(":8080")
}
