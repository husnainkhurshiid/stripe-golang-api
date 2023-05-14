package server

import (
	"stripeIntegration/model"

	"github.com/gin-gonic/gin"
)

func (s *Server) Payment(c *gin.Context) {

	var payment model.Payment
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(400, gin.H{"error": "Error binding JSON"})
		return
	}

	if err := s.api.StripePaymentAPI(c, payment); err != nil {
		c.JSON(400, gin.H{"error": "Error creating charge"})
		return
	}
}
