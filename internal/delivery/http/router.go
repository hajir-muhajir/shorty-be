package httpd

import (
	"github.com/gin-gonic/gin"
	"github.com/hajir.muhajir/shorty-be/internal/usecase"
)

func MapRoutes(r *gin.Engine,
	redirectUc *usecase.RedirectUC,
	linkUc *usecase.LinkUC,
	authUC *usecase.AuthUC,
	){
	r.GET("/:alias", RedirectHandler(redirectUc))

	api := r.Group("/api")
	
	// auth
	api.POST("/auth/register", RegisterHandler(authUC))
	api.POST("/auth/login", LoginHandler(authUC))

	// links
	api.POST("/links", CreateLinkHandler(linkUc))
}
