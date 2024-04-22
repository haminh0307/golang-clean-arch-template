package handler

import (
	"errors"
	"golang-clean-arch-template/internal/delivery/restapi"
	"golang-clean-arch-template/internal/domain"
	"golang-clean-arch-template/internal/domain/entity"
	"golang-clean-arch-template/internal/domain/interface/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	uc usecase.User
}

func NewUser(u usecase.User) *User {
	return &User{u}
}

// ReadOne.
//
//	@Summary		Read one user.
//	@Description	Read one user.
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			userID	path		string	true	"user id"
//	@Success		200		{object}	restapi.Response{data=object{user=entity.User},error=nil}
//	@Failure		401		{object}	restapi.Response{data=nil}
//	@Failure		403		{object}	restapi.Response{data=nil}
//	@Failure		404		{object}	restapi.Response{data=nil}
//	@Failure		500		{object}	restapi.Response{data=nil}
//	@Router			/users/{userID} [get].
//	@Security		Bearer
func (h *User) ReadOne(ctx *gin.Context) {
	id := entity.ID(ctx.Param("userID"))
	res, err := h.uc.ReadOne(ctx, id)
	switch {
	case err == nil:
		ctx.JSON(http.StatusOK, restapi.Response{Data: map[string]any{"user": res}})
	case errors.Is(err, domain.ErrNotFound):
		ctx.Status(http.StatusNotFound)
	default:
		ctx.JSON(http.StatusInternalServerError, restapi.Response{Error: err.Error()})
	}
}

// UpdateOne.
//
//	@Summary		Update one user.
//	@Description	Update one user.
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			userID	path	string				true	"user id"
//	@Param			update	body	entity.UserToUpdate	true	"user update data"
//	@Success		204
//	@Failure		400	{object}	restapi.Response{data=nil}
//	@Failure		401	{object}	restapi.Response{data=nil}
//	@Failure		403	{object}	restapi.Response{data=nil}
//	@Failure		404	{object}	restapi.Response{data=nil}
//	@Failure		500	{object}	restapi.Response{data=nil}
//	@Router			/users/{userID} [patch].
//	@Security		Bearer
func (h *User) UpdateOne(ctx *gin.Context) {
	id := entity.ID(ctx.Param("userID"))
	var update entity.UserToUpdate

	if err := ctx.ShouldBindJSON(&update); err != nil {
		ctx.JSON(http.StatusBadRequest, restapi.Response{Error: err.Error()})

		return
	}

	err := h.uc.UpdateOne(ctx, id, &update)
	switch {
	case err == nil:
		ctx.Status(http.StatusNoContent)
	case errors.Is(err, domain.ErrNotFound):
		ctx.Status(http.StatusNotFound)
	default:
		ctx.JSON(http.StatusInternalServerError, restapi.Response{Error: err.Error()})
	}
}
