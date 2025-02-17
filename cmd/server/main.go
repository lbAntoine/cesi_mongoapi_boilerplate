package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/lbAntoine/mongoapi_boilerplate/internal/api"
)

func main() {
	configPath := flag.String("config", "config.yaml", "path to config file")
	flag.Parse()

	cfg, err := api.LoadConfig(*configPath)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	app, err := api.NewApp(cfg)
	if err != nil {
		log.Fatalf("failed to initialize application: %v", err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := app.Run(); err != nil {
			log.Fatalf("failed to run application: %v", err)
		}
	}()

	// Wait for signal to shutdown the server
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := app.Shutdown(ctx); err != nil {
		log.Fatalf("failed to shutdown application: %v", err)
	}

	log.Println("Server exiting")
}
