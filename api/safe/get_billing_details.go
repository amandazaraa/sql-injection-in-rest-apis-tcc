package segura

import (
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

var idRegex = regexp.MustCompile(`^[1-9][0-9]{0,5}$`)

type Handler struct {
	repository GetBillingDetails
}

func NewHandler(repository GetBillingDetails) *Handler {
	return &Handler{repository: repository}
}

func (h *Handler) GetBillingDetails(c *gin.Context) {
	userID := c.Query("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "User ID is required"})
		return
	}

	if !idRegex.MatchString(userID) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is invalid"})
		return
	}

	billingDetails, err := h.repository.GetSafe(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Error to get safe billing details",
		})
	}

	c.JSON(http.StatusOK, billingDetails)
}
