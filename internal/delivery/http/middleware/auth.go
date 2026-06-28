package middleware

import (
	"strings"

	"Hermes/internal/delivery/http/response"

	"github.com/gin-gonic/gin"
)

// JWTMiddleware validates the incoming JWT token.
func JWTMiddleware() gin.HandlerFunc {
	return func(g *gin.Context) {

		jwt_token := strings.TrimSpace(g.GetHeader("Authorization"))
		if jwt_token == "" {
			response.Failure("No verification token provided.", "Authentication requried.")
		}

		g.Next()
	}
}
