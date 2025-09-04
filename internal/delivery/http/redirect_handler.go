package httpd

import (
	"net"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hajir.muhajir/shorty-api/internal/usecase"
)

func RedirectHandler(uc *usecase.RedirectUC) gin.HandlerFunc {
	return func(c *gin.Context) {
		alias := c.Param("alias")

		link, err := uc.Resolve(c.Request.Context(), alias)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		ip := clientIp(c)
		ref := c.Request.Referer()
		ua := c.Request.UserAgent()

		_ = uc.LogClick(c.Request.Context(), link, ip, ref, ua)

		c.Redirect(http.StatusFound, link.OriginalURL)
	}
}

func clientIp(c *gin.Context) string {
	xff := c.GetHeader("X-Forwarded-For")
	if xff != "" {
		parts := strings.Split(xff, ",")
		if len(parts) > 0 {
			return strings.TrimSpace(parts[0])
		}
	}
	host, _, err := net.SplitHostPort(c.Request.RemoteAddr)
	if err == nil {
		return host
	}
	return c.ClientIP()
}
