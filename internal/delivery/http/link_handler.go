package httpd

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hajir.muhajir/shorty-api/internal/usecase"
)

func CreateLinkHandler(uc *usecase.LinkUC) gin.HandlerFunc {
	type req struct {
		OriginalURL string     `json:"original_url"`
		Alias       *string    `json:"alias"`
		ExpiresAt   *time.Time `json:"expires_at"`
		MaxClicks   *int       `json:"max_clicks"`
	}

	return func(c *gin.Context) {
		var body req
		if err := c.ShouldBindBodyWithJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid body",
			})
			return
		}

		userID := "0f9d3472-149e-4814-8bd5-33cd2009b472" //hadcode dummy

		link, err := uc.Create(c.Request.Context(), usecase.CreateLinkRequest{
			UserID:      userID,
			OriginalURL: body.OriginalURL,
			Alias:       body.Alias,
			ExpiresAt:   body.ExpiresAt,
			MaxClicks:   body.MaxClicks,
		})

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusCreated, link)
	}
}
