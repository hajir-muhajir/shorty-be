package httpd

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hajir.muhajir/shorty-be/internal/usecase"
)

func RegisterHandler(uc *usecase.AuthUC) gin.HandlerFunc {
	type req struct {
		Email    string
		Password string
	}
	type res struct {
		Token string `json:"token"`
	}

	return func(c *gin.Context) {
		var r req
		if err := c.ShouldBindJSON(&r); err != nil || r.Email == "" || r.Password == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
			return
		}
		tok, err := uc.Register(c.Request.Context(), r.Email, r.Password)
		if err != nil {
			code := http.StatusBadRequest
			if err == usecase.ErrConflict {
				code = http.StatusConflict

			}

			c.JSON(code, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, res{Token: tok})
	}
}

func LoginHandler(uc *usecase.AuthUC) gin.HandlerFunc {
	type req struct {
		Email    string
		Password string
	}
	type res struct {
		Token string `json:"token"`
	}

	return func(c *gin.Context) {
		var r req
		if err := c.ShouldBindJSON(&r); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
			return
		}
		tok, err := uc.Login(c.Request.Context(), r.Email, r.Password)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": usecase.ErrUnauthorized.Error()})
			return
		}
		c.JSON(http.StatusOK, res{Token: tok})

	}
}
