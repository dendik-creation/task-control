package main

import (
	"log"

	"github.com/dendik-creation/task-control/internal/config"
	"github.com/dendik-creation/task-control/internal/database"
)

func main() {
	cfg := config.Load()

	db, err := database.NewPostgres(cfg)
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	// Contoh: Auto migrate nanti di sini
	_ = db

	log.Println("Application started in", cfg.AppEnv, "mode")
}
