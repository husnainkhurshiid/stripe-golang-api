package api

import (
	"stripeIntegration/model"

	"github.com/gin-gonic/gin"
)

type StripeAPI interface {
	StripePayment(c *gin.Context, stripePay model.Payment) error
}

type StripeAPIImpl struct {
	StripeKey string
}

func NewStripeAPIImpl() *StripeAPIImpl {

	return &StripeAPIImpl{
		StripeKey: "YOUR_STRIPE_SECRET_KEY",
	}
}

var _ StripeAPI = &StripeAPIImpl{}
