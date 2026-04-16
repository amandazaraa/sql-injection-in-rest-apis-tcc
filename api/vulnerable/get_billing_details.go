package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	repository GetBillingDetais
}

func NewHandler(repository GetBillingDetais) *Handler {
	return &Handler{repository: repository}
}

func (h *Handler) GetBillingDetails(c *gin.Context) {
	userID := c.Query("user_id")

	billingDetails, err := h.repository.GetVulnerable(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, billingDetails)
}
