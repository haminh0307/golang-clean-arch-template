package controller

import (
	"github.com/haminh0307/golang-clean-arch-template/internal/delivery/restapi/handler"
	"github.com/haminh0307/golang-clean-arch-template/internal/delivery/restapi/middleware"

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
