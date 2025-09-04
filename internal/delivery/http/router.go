package httpd

import (
	"github.com/gin-gonic/gin"
	"github.com/hajir.muhajir/shorty-be/internal/usecase"
)

func MapRoutes(r *gin.Engine, redirectUc *usecase.RedirectUC, linkUc *usecase.LinkUC){
	r.GET("/:alias", RedirectHandler(redirectUc))

	api := r.Group("/api")
	api.POST("/links", CreateLinkHandler(linkUc))
}
