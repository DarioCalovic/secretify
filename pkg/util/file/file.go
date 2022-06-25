package file

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func CreateFile(b []byte, p string) error {
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
		return err
	}
	return ioutil.WriteFile(p, b, 0644)
}
