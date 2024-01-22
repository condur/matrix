package routes

import (
	"github.com/condur/matrix/handlers"
	"github.com/condur/matrix/handlers/healthcheck"
	"github.com/condur/matrix/handlers/noroute"
	"github.com/condur/matrix/middleware"
	"github.com/gin-gonic/gin"
)

type Options struct {
}

// Initialize - define all app routes
func Initialize(opt *Options) *gin.Engine {
	// Gen a new blank gin router
	router := gin.New()

	// Handle no route found
	router.NoRoute(noroute.NoRoute)

	// Health check
	router.GET("/", healthcheck.Check)
	router.GET("/ping", healthcheck.Pong)

	// Global panic recovery middleware
	router.Use(middleware.Recovery())

	// Gen or set the request id
	router.Use(middleware.RequestId())

	// Global logging middleware
	router.Use(middleware.Logger())

	// Set the cache control
	router.Use(middleware.CacheControl())

	// Add API endpoints
	router.POST("/echo", handlers.Echo)
	router.POST("/invert", handlers.Invert)
	router.POST("/flatten", handlers.Flatten)
	router.POST("/sum", handlers.Sum)
	router.POST("/multiply", handlers.Multiply)

	return router
}
