package api

import (
	"stripeIntegration/controller"
	"stripeIntegration/model"

	"github.com/gin-gonic/gin"
)

func (api *StripeAPIImpl) StripePayment(c *gin.Context, stripePay model.Payment) error {
	secretKey := api.StripeKey
	if err := controller.StripePayment(c, secretKey, &stripePay); err != nil {
		return err
	}
	return nil
}
