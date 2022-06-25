package file

import (
	"github.com/DarioCalovic/secretify"
	"github.com/DarioCalovic/secretify/pkg/api/file/plattform/fs"
	"github.com/DarioCalovic/secretify/pkg/api/file/plattform/sqlite"
	"github.com/DarioCalovic/secretify/pkg/api/setting"
	utildb "github.com/DarioCalovic/secretify/pkg/util/db"
)

// Service represents file application interface
type Service interface {
	Create(data []byte, filename string, memetype string, size uint) (secretify.File, error)
	Read(identifier string) (content []byte, err error)
	View(identifier string) (file secretify.File, err error)
	ServiceConfig() setting.Service
	Repo() Repository
	Storage() Storage
}

// Repository represents file repository interface
type Repository interface {
	Create(utildb.DB, secretify.File) (secretify.File, error)
	ViewByIdentifier(utildb.DB, string) (secretify.File, error)
	Delete(utildb.DB, string) error
	DeleteExpired(utildb.DB) error
}

// Storage represents file storage interface
type Storage interface {
	Upload(content []byte, identifier string) error
	Retrieve(path string) ([]byte, error)
	Delete(path string) error
}

// File represents file application service
type File struct {
	db      utildb.DB
	repo    Repository
	storage Storage
	cfgSvc  setting.Service
}

// New creates new file application service
func New(db utildb.DB, sdb Repository, storage Storage, cfgSvc setting.Service) *File {
	return &File{db, sdb, storage, cfgSvc}
}

// Initialize initalizes file application service with defaults
func Initialize(db utildb.DB, cfgSvc setting.Service) *File {
	switch db.(type) {
	case *utildb.SQLiteDB:
		return New(db, sqlite.NewSQLiteFileRepository(), fs.NewFileSystemStorage(), cfgSvc)
	}
	return New(db, nil, nil, cfgSvc)
}
