package controller

import (
	"log"
	"net/http"
	"stripeIntegration/model"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
	"github.com/stripe/stripe-go/token"
)

func StripePayment(c *gin.Context, secretKey string, stripePay *model.Payment) error {
	stripe.Key = secretKey

	// Create a token using a credit or debit card
	tokenParams := &stripe.TokenParams{
		Card: &stripe.CardParams{
			Number:   stripe.String(stripePay.Card.Number),   // Card number
			ExpMonth: stripe.String(stripePay.Card.ExpMonth), // Expiration month
			ExpYear:  stripe.String(stripePay.Card.ExpYear),  // Expiration year
			CVC:      stripe.String(stripePay.Card.CVC),      // CVC code
		},
	}
	token, err := token.New(tokenParams)
	if err != nil {
		log.Fatalf("Error creating token: %v", err)
	}

	// Create a charge using the token
	params := &stripe.ChargeParams{
		Amount:      stripe.Int64(stripePay.Amount),
		Currency:    stripe.String(stripePay.Currency),
		Description: stripe.String(stripePay.Desc),
		Source: &stripe.SourceParams{
			Token: stripe.String(token.ID),
		},
		ReceiptEmail: stripe.String(stripePay.Email),
		Shipping: &stripe.ShippingDetailsParams{
			Name: stripe.String(stripePay.Name),
			Address: &stripe.AddressParams{
				Line1:      stripe.String(stripePay.Line1),
				Line2:      stripe.String(stripePay.Line2),
				City:       stripe.String(stripePay.City),
				State:      stripe.String(stripePay.State),
				PostalCode: stripe.String(stripePay.PostalCode),
				Country:    stripe.String(stripePay.Country),
			},
		},
	}
	_, err = charge.New(params)
	if err != nil {
		log.Fatalf("Error creating charge: %v", err)
	}

	// c.JSON(http.StatusOK, gin.H{"message": "Charge created successfully"})
	c.JSON(http.StatusOK, params)
	return nil
}
