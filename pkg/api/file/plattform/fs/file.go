package fs

import (
	"io/ioutil"
	"os"

	"github.com/DarioCalovic/secretify/pkg/util/file"
)

// FileSystem represents the client for the filesystem storage
type FileSystem struct{}

// NewFileSystemStorage creates a new instance of FileSystem
func NewFileSystemStorage() *FileSystem {
	return &FileSystem{}
}

// Upload uploads a new file to a filesystem
func (f FileSystem) Upload(content []byte, path string) error {
	err := file.CreateFile(content, path)
	if err != nil {
		return err
	}
	return nil
}

// Retrieve retrieves the content of the file from a filesystem
func (f FileSystem) Retrieve(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

// Delete removes the file from a filesystem
func (f FileSystem) Delete(path string) error {
	return os.Remove(path)
}
