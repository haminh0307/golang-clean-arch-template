package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/haminh0307/golang-clean-arch-template/internal/delivery/restapi"
	"github.com/haminh0307/golang-clean-arch-template/internal/delivery/restapi/request"
	"github.com/haminh0307/golang-clean-arch-template/internal/domain"
	"github.com/haminh0307/golang-clean-arch-template/internal/domain/entity"
	"github.com/haminh0307/golang-clean-arch-template/internal/domain/interface/usecase"

	"github.com/gin-gonic/gin"
)

type Authentication struct {
	uc usecase.Authentication
}

func NewAuthentication(a usecase.Authentication) *Authentication {
	return &Authentication{a}
}

// SignUp.
//
//	@Summary		Sign up new account.
//	@Description	Sign up new account.
//	@Tags			authentication
//	@Accept			json
//	@Produce		json
//	@Param			user	body	entity.UserToCreate	true	"User to be created"
//	@Success		201
//	@Failure		400	{object}	restapi.Response{data=nil}
//	@Failure		500	{object}	restapi.Response{data=nil}
//	@Header			204	{string}	Location	"/users/:id"
//	@Router			/authentication/signup [post].
func (h *Authentication) SignUp(ctx *gin.Context) {
	var user entity.UserToCreate
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, restapi.Response{Error: err.Error()})

		return
	}

	id, err := h.uc.SignUp(ctx, &user)
	switch {
	case err == nil:
		ctx.Header("Location", fmt.Sprintf("/users/%s", id.String()))
		ctx.Status(http.StatusCreated)
	default:
		ctx.JSON(http.StatusInternalServerError, restapi.Response{Error: err.Error()})
	}
}

// SignIn.
//
//	@Summary		Sign in.
//	@Description	Sign in.
//	@Tags			authentication
//	@Accept			json
//	@Produce		json
//	@Param			credentials	body		request.SignIn	true	"User credentials"
//	@Success		200			{object}	restapi.Response{data=object{token=string},error=nil}
//	@Failure		400			{object}	restapi.Response{data=nil}
//	@Failure		401			{object}	restapi.Response{data=nil}
//	@Failure		404			{object}	restapi.Response{data=nil}
//	@Failure		500			{object}	restapi.Response{data=nil}
//	@Router			/authentication/signin [post].
func (h *Authentication) SignIn(ctx *gin.Context) {
	var credentials request.SignIn

	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		ctx.JSON(http.StatusBadRequest, restapi.Response{Error: err.Error()})

		return
	}

	token, err := h.uc.SignIn(ctx, credentials.Username, credentials.Password)
	switch {
	case err == nil:
		ctx.JSON(http.StatusOK, restapi.Response{Data: map[string]any{"token": token}})
	case errors.Is(err, domain.ErrNotFound):
		ctx.Status(http.StatusNotFound)
	case errors.Is(err, domain.ErrWrongCredentials):
		ctx.JSON(http.StatusUnauthorized, restapi.Response{Error: err.Error()})
	default:
		ctx.JSON(http.StatusInternalServerError, restapi.Response{Error: err.Error()})
	}
}
