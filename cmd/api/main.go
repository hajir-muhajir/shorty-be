package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hajir.muhajir/shorty-be/internal/config"
	"github.com/hajir.muhajir/shorty-be/internal/db"
	httpd "github.com/hajir.muhajir/shorty-be/internal/delivery/http"
	gormrepo "github.com/hajir.muhajir/shorty-be/internal/repository/gorm"
	"github.com/hajir.muhajir/shorty-be/internal/service"
	"github.com/hajir.muhajir/shorty-be/internal/usecase"
	"github.com/joho/godotenv"
)

func main()  {
	err := godotenv.Load()
	if err != nil{
		log.Fatal("Failed to load env")
	}
	cfg := config.Load()

	gdb, err := db.Open(cfg)
	if err != nil{
		log.Fatalf("DB Open: %v", err)
	}
	sqlDB, _ := gdb.DB()
	defer sqlDB.Close()

	// Repositories
	linkRepo := gormrepo.NewLinkGorm(gdb)
	clickRepo := gormrepo.NewClickGorm(gdb)

	// Usecase
	redirectUC := usecase.NewRedirectUC(linkRepo, clickRepo)
	aliasGen := service.NewAliasGenerator()
	linkUC := usecase.NewLinkUC(linkRepo,aliasGen)

	// Http Server
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	r.GET("/health", func(c *gin.Context){
		var one int
		if err := gdb.Raw("SELECT 1").Scan(&one). Error; err != nil{
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
	})

	httpd.MapRoutes(r, redirectUC, linkUC)

	log.Printf("Listening on :%s", cfg.AppPort)
	if err := r.Run(":"+cfg.AppPort); err != nil{
		log.Fatal(err)
	}
}