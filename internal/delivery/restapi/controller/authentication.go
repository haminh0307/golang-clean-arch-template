package controller

import (
	"github.com/haminh0307/golang-clean-arch-template/internal/delivery/restapi/handler"

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
