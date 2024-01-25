package server

import (
	"context"
	"net/http"
	"time"

	"github.com/condur/matrix/log"
)

const (
	readHeaderTimeout = 1 * time.Second
	readTimeout       = 5 * time.Second
	writeTimeout      = 10 * time.Second
)

func NewServer(port string, handler http.Handler) *http.Server {
	return &http.Server{
		Addr:              port,
		Handler:           handler,
		ReadHeaderTimeout: readHeaderTimeout,
		ReadTimeout:       readTimeout,
		WriteTimeout:      writeTimeout,
	}
}

func Shutdown(srv *http.Server, timeout time.Duration) {
	// The context is used to inform the server it has a few seconds to finish
	// the requests it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	// Cancel context on exit
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Errorf("Server forced to shutdown: %s", err)
	} else {
		log.Info("Server stopped")
	}
}
