package httpd

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct{
	jwt.RegisteredClaims
}
func JWTAuthMiddleware(secret string) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		h := ctx.GetHeader("Authorization")
		if h == "" || !strings.HasPrefix(h, "Bearer ") {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "missing token",
			})
			return 
		}
		tokenStr := strings.TrimPrefix(h, "Bearer ")

		claims := &Claims{}
		tok, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (any, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrTokenUnverifiable
			}
			return []byte(secret), nil
		})
		if err != nil{
			if errors.Is(err, jwt.ErrTokenExpired){
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error": "token expired",
				})
				return 
			}
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token: " + err.Error(),
			})
			return
		}

		
		if !tok.Valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token",
			})
			return
		}

		sub := claims.Subject
		if sub == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "subject missing"})
			return
		}

		// Save userID to ctx
		ctx.Set("userID", sub)
		ctx.Next()
	}
}
