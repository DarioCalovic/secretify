package sqlite

import (
	"github.com/DarioCalovic/secretify"
	utildb "github.com/DarioCalovic/secretify/pkg/util/db"
	"gorm.io/gorm"
)

// SQLite represents the client for file table
type SQLite struct{}

// NewSQLiteFileRepository creates a new instance of SQLite
func NewSQLiteFileRepository() *SQLite {
	return &SQLite{}
}

// Create creates a new file on database
func (s SQLite) Create(db utildb.DB, file secretify.File) (secretify.File, error) {
	sqlitedb := db.(*utildb.SQLiteDB)
	err := sqlitedb.DBConn.Create(&file).Error
	return file, err
}

// ViewByIdentifier returns a single file by identifier
func (s SQLite) ViewByIdentifier(db utildb.DB, identifier string) (secretify.File, error) {
	sqlitedb := db.(*utildb.SQLiteDB)
	var file secretify.File
	err := sqlitedb.DBConn.Where("identifier = ?", identifier).First(&file).Error
	if err == gorm.ErrRecordNotFound {
		err = secretify.ErrRecordNotFound
	}
	return file, err
}

// Delete deletes the file from the database
func (s SQLite) Delete(db utildb.DB, identifier string) error {
	sqlitedb := db.(*utildb.SQLiteDB)
	var file secretify.File
	err := sqlitedb.DBConn.Where("identifier = ?", identifier).Delete(&file).Error
	if err != nil {
		return err
	}
	return nil
}

// DeleteExpired deletes all expired files
func (s SQLite) DeleteExpired(db utildb.DB) error {
	sqlitedb := db.(*utildb.SQLiteDB)
	var file secretify.File
	err := sqlitedb.DBConn.Where("expires_at <= datetime('now', 'localtime')").Delete(&file).Error
	if err != nil {
		return err
	}
	return nil
}
