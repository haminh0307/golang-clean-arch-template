package controller

import (
	"rainbow-love-memory/internal/delivery/restapi/handler"
	"rainbow-love-memory/internal/delivery/restapi/middleware"

	"github.com/gin-gonic/gin"
)

type user struct {
	handler       *handler.User
	authenticator *middleware.Authentication
}

func (c *user) RegisterRoutes(router gin.IRouter) {
	userGroup := router.Group("/users/:userID", c.authenticator.Authenticate)
	{
		userGroup.GET("", c.handler.ReadOne)
		userGroup.PATCH("", c.handler.UpdateOne)
	}
}

func NewUser(handler *handler.User, authenticator *middleware.Authentication) *user {
	return &user{handler, authenticator}
}
