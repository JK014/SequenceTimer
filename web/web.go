package web

import (
	"context"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/unrolled/secure"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// DefaultPort defines the default port the server will run on
const DefaultPort = 8080

// Global shutdown channel
var stopChannel = make(chan struct{})

// setupRouter sets up and returns a configured HTTP router
func setupRouter() http.Handler {
	secureMiddleware := secure.New(secure.Options{
		FrameDeny: true,
	})
	r := mux.NewRouter()
	r.Use(secureMiddleware.Handler)

	// Register all routes
	RegisterRoutes(r)

	return r
}

// StartServer starts the HTTP server with clean shutdown capability
func StartServer() {
	router := setupRouter()

	// Create the server with context support
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", DefaultPort),
		Handler: router,
	}

	// Create a context to handle shutdown signals
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// Run the server in a separate goroutine
	go func() {
		log.Printf("Starting server on port %d...\n", DefaultPort)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Server error: %v\n", err)
		}
	}()

	// Wait for either OS signal or /stop endpoint signal
	select {
	case <-ctx.Done():
		stop() // Stop receiving further shutdown signals
		log.Println("Shutdown signal received...")
	case <-stopChannel:
		log.Println("/stop endpoint triggered. Shutting down...")
	}

	// Gracefully shutdown the server
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Printf("Server shutdown error: %v\n", err)
	}

	log.Println("Server gracefully stopped.")
}
