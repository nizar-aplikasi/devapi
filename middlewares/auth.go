// File: middlewares/auth.go
package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"devapi/utils/jwtutil"
	"devapi/utils/response"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Ambil token dari header Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.Error("missing or invalid Authorization header"))
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		// Verifikasi token
		claims, err := jwtutil.VerifyAccessToken(tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.Error("invalid or expired token"))
			return
		}

		// (Opsional) simpan ke context jika butuh di handler
		c.Set("username", claims.Username)
		c.Set("roles", claims.Roles)

		// Lanjut ke handler berikutnya
		c.Next()
	}
}
