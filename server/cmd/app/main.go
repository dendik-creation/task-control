package main

import (
	"log"
	"os"

	"github.com/dendik-creation/task-control/internal/config"
	"github.com/dendik-creation/task-control/internal/database"
	deliveryHttp "github.com/dendik-creation/task-control/internal/delivery/http"
	"github.com/dendik-creation/task-control/internal/repository"
	"github.com/dendik-creation/task-control/internal/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()
	db, err := database.NewPostgres(cfg)
	if err != nil {
		log.Fatalf("Failed connect db: %v", err)
	}

	if cfg.AppEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	r.SetTrustedProxies([]string{"localhost"})

	userRepo := repository.NewUserRepository(db)
	columnRepo := repository.NewColumnRepository(db)
	// taskRepo := repository.NewTaskRepository(db)

	authUsecase := usecase.NewAuthUsecase(userRepo, columnRepo)

	deliveryHttp.NewAuthHandler(r, authUsecase)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Running at port:%s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Stopped because: %v", err)
	}
}
