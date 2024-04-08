package controller

import (
	"rainbow-love-memory/internal/delivery/restapi/handler"

	"github.com/gin-gonic/gin"
)

type authentication struct {
	handler *handler.Authentication
}

func (c *authentication) RegisterRoutes(router gin.IRouter) {
	authenticationGroup := router.Group("/authentication")
	{
		authenticationGroup.POST("/signin", c.handler.SignIn)
		authenticationGroup.POST("/signup", c.handler.SignUp)
	}
}

func NewAuthentication(handler *handler.Authentication) *authentication {
	return &authentication{handler}
}
