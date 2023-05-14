package main

import (
	"log"
	"net/http"
	"stripeIntegration/server"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.New()
	log.Fatal(http.ListenAndServe(":8080", server.NewServerImpl(router)))

}
