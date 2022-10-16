package secret

import (
	"errors"
	"fmt"
	"time"

	"github.com/DarioCalovic/secretify"
	"github.com/DarioCalovic/secretify/pkg/api/setting"
	"github.com/DarioCalovic/secretify/pkg/util/nanoid"
)

// Create creates a new encrypted secret
func (s *Secret) Create(ciphertext string, hasPassphrase bool, expiresAt time.Time, revealOnce bool, destroyManual bool, fileID int) (secretify.Secret, error) {
	identifier, err := nanoid.GenerateIdentifier(s.cfgSvc.Policy().Identifier.Size)
	if err != nil {
		return secretify.Secret{}, err
	}
	scro := secretify.Secret{
		Identifier:    identifier,
		Cipher:        ciphertext,
		ExpiresAt:     expiresAt,
		RevealOnce:    revealOnce,
		DestroyManual: destroyManual,
		HasPassphrase: hasPassphrase,
	}
	if fileID > 0 {
		scro.File.ID = fileID
	}
	secret, err := s.repo.Create(s.db, scro)
	if err != nil {
		return secretify.Secret{}, err
	}
	return secret, nil
}

// Create creates a new encrypted secret with an encrypted file
func (s *Secret) CreateWithFile(ciphertext string, hasPassphrase bool, expiresAt time.Time, revealOnce bool, destroyManual bool, fileIdentifier string) (secretify.Secret, error) {
	if s.fileSvc == nil {
		return secretify.Secret{}, errors.New("could not contact file service, file storage not enabled probably")
	}
	file, err := s.fileSvc.Repo().ViewByIdentifier(s.db, fileIdentifier)
	if err != nil {
		return secretify.Secret{}, err
	}
	return s.Create(ciphertext, hasPassphrase, expiresAt, revealOnce, destroyManual, file.ID)
}

// View returns the secret's information, deletes it afterwards and
// sends an email if configured
func (s *Secret) View(identifier string, onlyMeta bool) (secret secretify.Secret, deleted bool, err error) {
	secret, err = s.repo.ViewByIdentifier(s.db, identifier)
	if err != nil {
		return
	}

	// Remove cipher and retuns instantly if only meta
	if onlyMeta {
		secret.Cipher = ""
		return
	}

	// Check if only once readable
	if secret.RevealOnce {
		err = s.repo.Delete(s.db, identifier)
		if err != nil {
			return
		}
		if secret.FileID > 0 && s.fileSvc != nil {
			go func() {
				time.Sleep(30 * time.Second)
				err = s.fileSvc.Repo().Delete(s.db, secret.File.Identifier)
				if err != nil {
					return
				}
				err = s.fileSvc.Storage().Delete(secret.File.Path)
				if err != nil {
					return
				}
				fmt.Println("File was deleted")
			}()
		}
		deleted = true
	}
	return
}

// Delete the secret
func (s *Secret) Delete(identifier string) error {
	secret, err := s.repo.ViewByIdentifier(s.db, identifier)
	if err != nil {
		return err
	}
	if !secret.DestroyManual {
		return errors.New("destroying manually is disabled for this secret")
	}
	if secret.FileID > 0 && s.fileSvc != nil {
		err = s.fileSvc.Repo().Delete(s.db, secret.File.Identifier)
		if err != nil {
			return err
		}
		err = s.fileSvc.Storage().Delete(secret.File.Path)
		if err != nil {
			return err
		}
	}
	return s.repo.Delete(s.db, identifier)
}

// DeleteExpired all expired secrets (and associated files)
func (s *Secret) DeleteExpired() error {
	// Get all expired
	secrets, err := s.repo.ViewAllExpired(s.db)
	if err != nil {
		return err
	}
	// Delete all expired
	err = s.repo.DeleteExpired(s.db)
	if err != nil {
		return err
	}
	// Delete all files
	for _, secret := range secrets {
		if secret.FileID > 0 && s.fileSvc != nil {
			err = s.fileSvc.Repo().Delete(s.db, secret.File.Identifier)
			if err != nil {
				return err
			}
			err = s.fileSvc.Storage().Delete(secret.File.Path)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *Secret) ServiceConfig() setting.Service {
	return s.cfgSvc
}
