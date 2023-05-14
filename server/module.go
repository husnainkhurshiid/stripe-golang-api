package server

import (
	"stripeIntegration/api"

	"github.com/gin-gonic/gin"
)

type ServerImpl interface {
	StripePayment(c *gin.Context)
}
type Server struct {
	api api.StripeAPI
}

func NewServer() *Server {
	api := api.NewStripeAPIImpl()
	return &Server{
		api: api,
	}
}

func NewServerImpl(r *gin.Engine) *gin.Engine {
	server := NewServer()

	listRoute := r.Group("/payment")
	listRoute.POST("/charge", server.Payment)
	return r
}

var _ServerImpl = &Server{}
