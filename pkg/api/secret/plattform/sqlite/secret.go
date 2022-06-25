package sqlite

import (
	"github.com/DarioCalovic/secretify"
	utildb "github.com/DarioCalovic/secretify/pkg/util/db"
	"gorm.io/gorm"
)

// SQLite represents the client for secret table
type SQLite struct{}

// NewSQLiteSecretRepository creates a new instance of SQLite
func NewSQLiteSecretRepository() *SQLite {
	return &SQLite{}
}

// Create creates a new secret on database
func (s SQLite) Create(db utildb.DB, secret secretify.Secret) (secretify.Secret, error) {
	sqlitedb := db.(*utildb.SQLiteDB)
	err := sqlitedb.DBConn.Create(&secret).Error
	return secret, err
}

// ViewByIdentifier returns a single secret by identifier
func (s SQLite) ViewByIdentifier(db utildb.DB, identifier string) (secretify.Secret, error) {
	sqlitedb := db.(*utildb.SQLiteDB)
	var secret secretify.Secret
	err := sqlitedb.DBConn.Joins("File").Where("secrets.identifier = ?", identifier).First(&secret).Error
	if err == gorm.ErrRecordNotFound {
		err = secretify.ErrRecordNotFound
	}
	return secret, err
}

// Delete deletes the secret from the database
func (s SQLite) Delete(db utildb.DB, identifier string) error {
	sqlitedb := db.(*utildb.SQLiteDB)
	var secret secretify.Secret
	err := sqlitedb.DBConn.Where("identifier = ?", identifier).Delete(&secret).Error
	if err != nil {
		return err
	}
	return nil
}

// ViewAllExpired returns all expired secrets
func (s SQLite) ViewAllExpired(db utildb.DB) (secrets []secretify.Secret, err error) {
	sqlitedb := db.(*utildb.SQLiteDB)
	err = sqlitedb.DBConn.Where("expires_at <= datetime('now', 'localtime')").Find(&secrets).Error
	if err != nil {
		return
	}
	return
}

// DeleteExpired deletes all expired secrets
func (s SQLite) DeleteExpired(db utildb.DB) error {
	sqlitedb := db.(*utildb.SQLiteDB)
	var secret secretify.Secret
	err := sqlitedb.DBConn.Where("expires_at <= datetime('now', 'localtime')").Delete(&secret).Error
	if err != nil {
		return err
	}
	return nil
}
