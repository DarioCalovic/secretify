package mockdb

import (
	"github.com/DarioCalovic/secretify"
	utildb "github.com/DarioCalovic/secretify/pkg/util/db"
)

type Secret struct {
	CreateFn            func(db utildb.DB, secret secretify.Secret) (secretify.Secret, error)
	ViewByIdentifiersFn func(db utildb.DB, identifier string) (secretify.Secret, error)
	ViewAllExpiredFn    func(db utildb.DB) (secrets []secretify.Secret, err error)
	DeleteFn            func(db utildb.DB, identifier string) error
	DeleteExpiredFn     func(db utildb.DB) error
}

// Create mock
func (s *Secret) Create(db utildb.DB, secret secretify.Secret) (secretify.Secret, error) {
	return s.CreateFn(db, secret)
}

// ViewByIdentifier mock
func (s *Secret) ViewByIdentifier(db utildb.DB, identifier string) (secretify.Secret, error) {
	return s.ViewByIdentifiersFn(db, identifier)
}

// ViewAllExpired mock
func (s *Secret) ViewAllExpired(db utildb.DB) (secrets []secretify.Secret, err error) {
	return s.ViewAllExpiredFn(db)
}

// Delete mock
func (s *Secret) Delete(db utildb.DB, identifier string) error {
	return s.DeleteFn(db, identifier)
}

// DeleteExpired mock
func (s *Secret) DeleteExpired(db utildb.DB) error {
	return s.DeleteExpiredFn(db)
}
