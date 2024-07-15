package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"anc-oms-input-api/common"
	"anc-oms-input-api/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load environment variables
	common.LoadEnv()

	// Connect to MongoDB
	common.ConnectDB()

	// Initialize Gin router
	r := gin.Default()

	// Define routes for order service
	orderGroup := r.Group("/orders")
	{
		orderGroup.POST("", handlers.CreateOrder)
		orderGroup.DELETE("/cancel", handlers.CancelOrder)
	}

	// Server configurations
	bindAddress := common.GetEnv("BIND_ADDRESS", ":9090")

	// Create HTTP server with timeouts
	s := &http.Server{
		Addr:         bindAddress,
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// Start the HTTP server in a goroutine
	go func() {
		log.Printf("Starting server on %s", bindAddress)
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on %s: %v", bindAddress, err)
		}
	}()

	// Trap SIGINT (Ctrl+C) and SIGTERM (termination signal)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	sig := <-quit
	log.Printf("Signal received: %v", sig)

	// Create a context for shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Shutdown the server gracefully
	if err := s.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown error: %v", err)
	}

	log.Println("Server shutdown complete")
}