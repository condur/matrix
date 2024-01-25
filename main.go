package main

import (
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/condur/matrix/log"
	"github.com/condur/matrix/routes"
	"github.com/condur/matrix/server"
	"github.com/gin-gonic/gin"
)

const (
	port string = ":8080"
)

func main() {
	// Set the log configuration
	log.Config("matrix", "debug")
	log.Infof("Starting HTTP server, listening on port %s", port)

	// Set the GIN options
	gin.SetMode(gin.DebugMode)
	gin.DisableBindValidation()

	// Initialize the router
	router := routes.Initialize(&routes.Options{})

	// Create an HTTP server
	srv := server.NewServer(port, router)

	// Gracefully shutdown the server on exit
	defer server.Shutdown(srv, 10*time.Second)

	// Initializing the server in a goroutine so that it won't block the graceful shutdown.
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Errorf("Failed to serve: %s", err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server with a timeout.
	quit := make(chan os.Signal, 1)

	// Kill (no param) default sends syscall.SIGTERM
	// Kill -2 is syscall.SIGINT
	// Kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Block the quit channel is closed
	<-quit

	// Print shutting down message
	log.Info("Shutting down server...")
}
