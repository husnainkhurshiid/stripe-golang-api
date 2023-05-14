package main

import (
	"log"
	"stripeIntegration/controller"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.POST("/charge", controller.StripePayment)
	log.Fatal(router.Run(":8080"))

}
