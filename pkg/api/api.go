package api

import (
	"crypto/sha1"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/DarioCalovic/secretify/pkg/api/auth"
	atransport "github.com/DarioCalovic/secretify/pkg/api/auth/transport"
	"github.com/DarioCalovic/secretify/pkg/api/file"
	ftransport "github.com/DarioCalovic/secretify/pkg/api/file/transport"
	"github.com/DarioCalovic/secretify/pkg/api/secret"
	stransport "github.com/DarioCalovic/secretify/pkg/api/secret/transport"
	"github.com/DarioCalovic/secretify/pkg/api/setting"
	sgtransport "github.com/DarioCalovic/secretify/pkg/api/setting/transport"
	utilconfig "github.com/DarioCalovic/secretify/pkg/util/config"
	utildb "github.com/DarioCalovic/secretify/pkg/util/db"
	"github.com/DarioCalovic/secretify/pkg/util/jwt"
	"github.com/DarioCalovic/secretify/pkg/util/mail"
	authMw "github.com/DarioCalovic/secretify/pkg/util/middleware/auth"
	"github.com/DarioCalovic/secretify/pkg/util/secure"
	"github.com/DarioCalovic/secretify/pkg/util/server"
	"github.com/gorilla/mux"
)

const (
	apiVersion = "v1"
)

// Start starts the API service
func Start(cfg *utilconfig.Configuration, db utildb.DB, mailer *mail.Mailer) error {
	r := server.New()
	r.Use(loggingMiddleware)
	s := r.PathPrefix(fmt.Sprintf("%s/%s/%s", cfg.Server.BasePath, "api", apiVersion)).Subrouter()

	// Protected
	var (
		sec            = secure.New(sha1.New())
		authMiddleware mux.MiddlewareFunc
	)
	if cfg.Server.Auth.Enabled {
		jwt, err := jwt.New(cfg.Server.Auth.JWT.SigningAlgorithm, os.Getenv("JWT_SECRET"), cfg.Server.Auth.JWT.DurationMinutes, cfg.Server.Auth.JWT.MinSecretLength)
		if err != nil {
			return err
		}
		authMiddleware = authMw.Middleware(jwt)
		atransport.NewHTTP(auth.Initialize(db, jwt, sec), s, authMiddleware)
		s.Use(authMiddleware)
	}

	cfgService := setting.Initialize(cfg)
	fileService := file.Initialize(db, cfgService)
	stransport.NewHTTP(secret.Initialize(db, cfgService, fileService, mailer), s)
	ftransport.NewHTTP(fileService, s)
	sgtransport.NewHTTP(cfgService, s)

	server.Start(r, &server.Config{
		Address:             cfg.Server.Address,
		ReadTimeoutSeconds:  cfg.Server.ReadTimeout,
		WriteTimeoutSeconds: cfg.Server.WriteTimeout,
		Debug:               cfg.Server.Debug,
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
