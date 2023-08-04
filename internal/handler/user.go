package handler

import (
	"fmt"
	"github.com/6a6ydoping/ChitChat/api"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// createUser registration new user
//
//	@Summary		Create user
//	@Description	Create new user
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			req	body	api.RegisterRequest	true	"req body"
//
//	@Success		201
//	@Failure		400	{object}	api.Error
//	@Failure		500	{object}	api.Error
//	@Router			/user/register [post]
func (h *Handler) createUser(ctx *gin.Context) {
	var req api.RegisterRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("bind json err: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}
	fmt.Println(req)
	err = h.userService.CreateUser(ctx, &req.User)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &api.Error{
			Code:    -2,
			Message: err.Error(),
		})
		return
	}

	ctx.Status(http.StatusCreated)
}

// loginUser log in as an existing user
//
//	@Summary		Login user
//	@Description	Log in as an existing user and get auth token
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			req	body	api.LoginRequest	true	"req body"
//
//	@Success		200
//	@Failure		400	{object}	api.Error
//	@Failure		500	{object}	api.Error
//	@Router			/user/login [post]
func (h *Handler) loginUser(ctx *gin.Context) {
	var req api.LoginRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("bind json err: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}

	accessToken, err := h.userService.Login(ctx, req.Username, req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &api.Error{
			Code:    -2,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &api.Ok{
		Code:    0,
		Message: "success",
		Data:    accessToken,
	})
}
