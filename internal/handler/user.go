package handler

import (
	"fmt"
	"github.com/6a6ydoping/ChitChat/api"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

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
	err = h.service.CreateUser(ctx, &req.User)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &api.Error{
			Code:    -2,
			Message: err.Error(),
		})
		return
	}

	ctx.Status(http.StatusCreated)
}
