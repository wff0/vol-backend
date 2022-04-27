package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vol-backend/model/response"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		_, e := context.Request.Cookie("userID")
		if e == nil {
			context.Next()
		} else {
			context.Abort()
			response.Result(http.StatusUnauthorized, map[string]interface{}{}, "用户未登录", context)
		}

	}
}
