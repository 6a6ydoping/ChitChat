package handler

import (
	"errors"
	"github.com/6a6ydoping/ChitChat/api"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const authUserID = "auth_user_id"

func (h *Handler) authMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader("authorization")
		if authorizationHeader == "" {
			err := errors.New("authorization header is not set")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, &api.Error{
				Code:    -1,
				Message: err.Error(),
			})
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := errors.New("authorization header incorrect format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, &api.Error{
				Code:    -2,
				Message: err.Error(),
			})
			return
		}

		userID, err := h.userService.VerifyToken(fields[1])
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, &api.Error{
				Code:    -3,
				Message: err.Error(),
			})
			return
		}

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(204)
			return
		}

		ctx.Set(authUserID, userID)
		ctx.Next()
	}
}
