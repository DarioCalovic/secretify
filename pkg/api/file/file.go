package file

import (
	"errors"
	"fmt"
	"os"

	"github.com/DarioCalovic/secretify"
	"github.com/DarioCalovic/secretify/pkg/api/setting"
	"github.com/DarioCalovic/secretify/pkg/util/nanoid"
)

// Create creates a new encrypted file
func (f *File) Create(data []byte, filename string, memetype string, size uint) (secretify.File, error) {

	// Max size
	if size >= f.ServiceConfig().Policy().Storage.FileSystem.MaxFileSize {
		return secretify.File{}, errors.New("max file size reached")
	}

	identifier, err := nanoid.GenerateIdentifier(f.cfgSvc.Policy().Identifier.Size)
	if err != nil {
		return secretify.File{}, err
	}

	// Upload file to fs
	path := fmt.Sprintf("%s%s%s", f.ServiceConfig().Storage().FileSystem.Location, string(os.PathSeparator), identifier)
	err = f.storage.Upload(data, path)
	if err != nil {
		return secretify.File{}, err
	}

	// Create file on db
	file, err := f.repo.Create(f.db, secretify.File{
		Identifier: identifier,
		Name:       filename,
		Type:       memetype,
		Size:       size,
		Path:       path,
	})
	if err != nil {
		return secretify.File{}, err
	}
	return file, nil
}

// Read returns the file's content
func (f *File) Read(identifier string) (content []byte, err error) {
	return f.storage.Retrieve(identifier)
}

// View returns the file's information
func (f *File) View(identifier string) (file secretify.File, err error) {
	file, err = f.repo.ViewByIdentifier(f.db, identifier)
	if err != nil {
		return
	}
	return
}

func (f *File) ServiceConfig() setting.Service {
	return f.cfgSvc
}

func (f *File) Repo() Repository {
	return f.repo
}

func (f *File) Storage() Storage {
	return f.storage
}
