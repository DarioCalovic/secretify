package secret

import (
	"time"

	"github.com/DarioCalovic/secretify"
	"github.com/DarioCalovic/secretify/pkg/api/file"
	"github.com/DarioCalovic/secretify/pkg/api/secret/plattform/sqlite"
	"github.com/DarioCalovic/secretify/pkg/api/setting"
	utildb "github.com/DarioCalovic/secretify/pkg/util/db"
	"github.com/DarioCalovic/secretify/pkg/util/mail"
)

// Service represents secret application interface
type Service interface {
	Create(ciphertext string, hasPassphrase bool, expiresAt time.Time, revealOnce bool, destroyManual bool, fileID int, email string, webhookAddr string) (secretify.Secret, error)
	CreateWithFile(ciphertext string, hasPassphrase bool, expiresAt time.Time, revealOnce bool, destroyManual bool, fileIdentifier string, email string, webhookAddr string) (secretify.Secret, error)
	View(identifier string, onlyMeta bool) (secret secretify.Secret, deleted bool, err error)
	Delete(identifier string) error
	DeleteExpired() error
	ServiceConfig() setting.Service
}

// Repository represents secret repository interface
type Repository interface {
	Create(utildb.DB, secretify.Secret) (secretify.Secret, error)
	ViewByIdentifier(utildb.DB, string) (secretify.Secret, error)
	Delete(utildb.DB, string) error
	ViewAllExpired(utildb.DB) ([]secretify.Secret, error)
	DeleteExpired(utildb.DB) error
}

// Secret represents secret application service
type Secret struct {
	db      utildb.DB
	repo    Repository
	cfgSvc  setting.Service
	fileSvc file.Service
	mailer  *mail.Mailer
}

// New creates new secret application service
func New(db utildb.DB, sdb Repository, cfgSvc setting.Service, fileSvc file.Service, mailer *mail.Mailer) *Secret {
	return &Secret{db, sdb, cfgSvc, fileSvc, mailer}
}

// Initialize initalizes secret application service with defaults
func Initialize(db utildb.DB, cfgSvc setting.Service, fileSvc file.Service, mailer *mail.Mailer) *Secret {
	switch db.(type) {
	case *utildb.SQLiteDB:
		return New(db, sqlite.NewSQLiteSecretRepository(), cfgSvc, fileSvc, mailer)
	}
	return New(db, nil, cfgSvc, fileSvc, mailer)
}
