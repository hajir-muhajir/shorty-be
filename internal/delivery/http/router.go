package httpd

import (
	"github.com/gin-gonic/gin"
	"github.com/hajir.muhajir/shorty-api/internal/usecase"
)

func MapRoutes(r *gin.Engine, redirectUc *usecase.RedirectUC){
	r.GET("/:alias", RedirectHandler(redirectUc))
}
