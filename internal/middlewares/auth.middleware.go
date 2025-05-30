package middlewares

import (
	"go_ecommerce_backend/pkg/respone"

	"github.com/gin-gonic/gin"
)

func AuthenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != "valid-token" {
			respone.ErroRespone(c, respone.ErrInvalidToken, "")
			c.Abort()
			return
		}
		c.Next()
	}
}
