package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/markbmullins/city-developer/pkg/game"
	"github.com/markbmullins/city-developer/pkg/server"
)

func main() {
	world := game.InitializeGame()

	// Run the game loop
	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()

		for range ticker.C {
			world.Update()
		}
	}()

	// Start the server and get the server instance for graceful shutdown
	srv := server.StartServer(world)

	// Set up channel to listen for termination signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// Block until a termination signal is received
	<-quit
	log.Println("Received termination signal, shutting down gracefully...")

	// Create a context with a timeout to allow graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown: %v\n", err)
	} else {
		log.Println("Server shutdown complete")
	}
}
