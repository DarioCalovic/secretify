package db

import (
	"github.com/DarioCalovic/secretify"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type SQLiteDB struct {
	name   string
	DBConn *gorm.DB
}

func NewSQLiteDB(name string) *SQLiteDB {
	return &SQLiteDB{name: name}
}

func (sqldb *SQLiteDB) Initialize() error {
	db, err := gorm.Open(sqlite.Open(sqldb.name), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(err.Error())
	}
	// Migrate the schema
	db.AutoMigrate(&secretify.Secret{})
	db.AutoMigrate(&secretify.File{})
	sqldb.DBConn = db
	return err
}
