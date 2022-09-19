package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// New instantates new mux router
func New() *mux.Router {
	r := mux.NewRouter()
	return r
}

// Config represents server specific config
type Config struct {
	Address             string
	ReadTimeoutSeconds  int
	WriteTimeoutSeconds int
	Debug               bool
	CORSAllowedOrigins  []string
}

// Start starts mux server
func Start(r *mux.Router, cfg *Config) {
	// TODO : make configurable
	cors := handlers.CORS(
		handlers.AllowedHeaders([]string{"content-type"}),
		handlers.AllowedOrigins(cfg.CORSAllowedOrigins),
		handlers.AllowCredentials(),
		handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "DELETE"}),
	)
	srv := &http.Server{
		Handler:      cors(r),
		Addr:         cfg.Address,
		ReadTimeout:  time.Duration(cfg.ReadTimeoutSeconds) * time.Second,
		WriteTimeout: time.Duration(cfg.WriteTimeoutSeconds) * time.Second,
	}
	// Start server
	go func() {
		fmt.Println("Starting server")
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal("Failed to start server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
