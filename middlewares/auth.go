package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/salimmia/events-go/utils"
)

func Authenticate(context *gin.Context){
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message" : "Not authorized"})
		return
	}

	userId, email, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message" : "Unauthorized.", "error" : err.Error()})
		return
	}

	context.Set("user_id", userId)
	context.Set("email", email)
	context.Next()
}