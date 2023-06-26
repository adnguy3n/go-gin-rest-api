package middlewares

import (
	"go-gin-rest-api/src/authenthication"

	"github.com/gin-gonic/gin"
)

/*
 * Authenthication method.
 */
func Authenthicate() gin.HandlerFunc {
	return func(context *gin.Context) {
		// If no token found at the header, throw a 401 error.
		tokenString := context.GetHeader("Authorization")
		if tokenString == "" {
			context.JSON(401, gin.H{"error": "request does not contain an access token"})
			context.Abort()
			return
		}

		// Validate Token. If token is found to be invalid or expired, throw 401.
		err := authenthication.ValidateToken(tokenString)
		if err != nil {
			context.JSON(401, gin.H{"error": err.Error()})
			context.Abort()
			return
		}

		context.Next()
	}
}
