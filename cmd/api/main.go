package main

import (
	"log"
	"time"

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
	userRepo := gormrepo.NewUserGorm(gdb)
	linkRepo := gormrepo.NewLinkGorm(gdb)
	clickRepo := gormrepo.NewClickGorm(gdb)

	// services
	aliasGen := service.NewAliasGenerator()
	hasher := service.NewHasher()
	jwtSigner := service.NewJWTSigner(cfg.JWTSecret, time.Duration(cfg.JWTTTLMinutes) * time.Minute)

	// Usecase
	redirectUC := usecase.NewRedirectUC(linkRepo, clickRepo)
	linkUC := usecase.NewLinkUC(linkRepo,aliasGen)
	authUC := usecase.NewAuthUC(userRepo, hasher, jwtSigner)

	// Http Server
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	r.GET("/health", httpd.HealthHandler(gdb))

	httpd.MapRoutes(r, redirectUC, linkUC, authUC, cfg.JWTSecret)

	log.Printf("Listening on :%s", cfg.AppPort)
	if err := r.Run(":"+cfg.AppPort); err != nil{
		log.Fatal(err)
	}
}