package main

import (
	"log"

	"github.com/Odvin/go-live-reloading/config"
	"github.com/Odvin/go-live-reloading/internal/adapters/db"
	"github.com/Odvin/go-live-reloading/internal/adapters/web"
	"github.com/Odvin/go-live-reloading/internal/application"
)

func main() {
	cfg := config.InitConfig()

	dbAdapter, err := db.Init(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}

	defer dbAdapter.DB.Close()

	appAdapter := application.Init(dbAdapter)

	webService := web.Init(appAdapter, cfg.Port, "1.0.0", cfg.Env)

	log.Printf("starting server on port: %d (env: %s)", cfg.Port, cfg.Env)

	webService.Serve()
}
