package httpd

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func HealthHandler(db *gorm.DB)gin.HandlerFunc{
	return func(c *gin.Context) {
		var one int
		if err := db.Raw("SELECT 1").Scan(&one). Error; err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "failed",
				"message": "Failed to connect to database",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"messahe": "Success to connect to database",
		})
	}
}