package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/hazaloolu/openUp_backend/internal/router"
	"github.com/hazaloolu/openUp_backend/internal/storage"
	"github.com/joho/godotenv"
)

// load env

func LoadEnv() error {
	if err := godotenv.Load("config/.env"); err != nil {
		return err
	}
	return nil
}

func main() {

	err := LoadEnv()

	if err != nil {
		log.Fatalf("Error loading env file : %v", err)
	}

	storage.InitDB()
	r := router.SetUpRouter()

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		log.Print("Server is starting at port 8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe Failed: %v", err)
		}
	}()

	// Graceful shutdown

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c
	log.Println("Interruption signal received, shutting down.....")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("server shutdown failed: %v", err)

		log.Println("server exited")
	}

}
