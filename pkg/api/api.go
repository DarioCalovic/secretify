package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/DarioCalovic/secretify/pkg/api/file"
	ftransport "github.com/DarioCalovic/secretify/pkg/api/file/transport"
	"github.com/DarioCalovic/secretify/pkg/api/outlook"
	otransport "github.com/DarioCalovic/secretify/pkg/api/outlook/transport"
	"github.com/DarioCalovic/secretify/pkg/api/secret"
	stransport "github.com/DarioCalovic/secretify/pkg/api/secret/transport"
	"github.com/DarioCalovic/secretify/pkg/api/setting"
	sgtransport "github.com/DarioCalovic/secretify/pkg/api/setting/transport"
	utilconfig "github.com/DarioCalovic/secretify/pkg/util/config"
	utildb "github.com/DarioCalovic/secretify/pkg/util/db"
	"github.com/DarioCalovic/secretify/pkg/util/server"
)

const (
	apiVersion = "v1"
)

// Start starts the API service
func Start(cfg *utilconfig.Configuration, db utildb.DB) error {
	r := server.New()
	r.HandleFunc("/", status).
		Methods("GET")
	r.Use(loggingMiddleware)
	s := r.PathPrefix(fmt.Sprintf("%s/%s/%s", cfg.Server.BasePath, "api", apiVersion)).Subrouter()

	// Status
	{
		s.HandleFunc("", status).
			Methods("GET")
		s.HandleFunc("/healthz", status).
			Methods("GET")
	}

	cfgService := setting.Initialize(cfg)
	var fileService *file.File
	if cfg.Policy.Storage.Enabled {
		fileService = file.Initialize(db, cfgService)
		ftransport.NewHTTP(fileService, s)
	}
	stransport.NewHTTP(secret.Initialize(db, cfgService, fileService), s)
	sgtransport.NewHTTP(cfgService, s)

	// Outlook
	if cfg.Outlook.Enabled {
		otransport.NewHTTP(outlook.Initialize(cfgService), s, cfg.Outlook.AppID)
	}

	server.Start(r, &server.Config{
		Address:             cfg.Server.Address,
		ReadTimeoutSeconds:  cfg.Server.ReadTimeout,
		WriteTimeoutSeconds: cfg.Server.WriteTimeout,
		Debug:               cfg.Server.Debug,
		CORSAllowedOrigins:  cfg.Policy.Security.CORS.AllowedOrigins,
	})
	return nil
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func status(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
}
