package secret

import (
	"fmt"
	"time"

	"github.com/DarioCalovic/secretify"
	"github.com/DarioCalovic/secretify/pkg/api/setting"
	"github.com/DarioCalovic/secretify/pkg/util/mail"
	"github.com/DarioCalovic/secretify/pkg/util/nanoid"
	mailjet "github.com/mailjet/mailjet-apiv3-go"
)

// Create creates a new encrypted secret
func (s *Secret) Create(ciphertext string, hasPassphrase bool, expiresAt time.Time, revealOnce bool, fileID int, email string, webhookAddr string) (secretify.Secret, error) {
	identifier, err := nanoid.GenerateIdentifier(s.cfgSvc.Policy().Identifier.Size)
	if err != nil {
		return secretify.Secret{}, err
	}
	scro := secretify.Secret{
		Identifier:    identifier,
		Cipher:        ciphertext,
		ExpiresAt:     expiresAt,
		RevealOnce:    revealOnce,
		Email:         email,
		WebhookAddr:   webhookAddr,
		HasPassphrase: hasPassphrase,
	}
	if fileID > 0 {
		scro.File.ID = fileID
	}
	secret, err := s.repo.Create(s.db, scro)
	if err != nil {
		return secretify.Secret{}, err
	}

	// Send mail
	if s.cfgSvc.Policy().Mail.Enabled && email != "" {
		go func() {
			htmlBody, err := mail.SecretCreatedHTMLBody("link", s.cfgSvc.Meta().Hoster.Name, s.cfgSvc.Meta().Hoster.Address)
			if err != nil {
				fmt.Println(err)
				return
			}
			textBody, err := mail.SecretCreatedTextBody("link", s.cfgSvc.Meta().Hoster.Name, s.cfgSvc.Meta().Hoster.Address)
			if err != nil {
				fmt.Println(err)
				return
			}
			err = s.mailer.Send(email, "Your secret has been created", htmlBody, textBody, "SecretifySecretCreated")
			if err != nil {
				fmt.Println(err)
				return
			}
		}()
	}

	if s.cfgSvc.Policy().Webhook.Enabled && secret.WebhookAddr != "" {
		// TODO : send webhook
		fmt.Println("Posting to webhook ", secret.WebhookAddr)
	}
	return secret, nil
}

// Create creates a new encrypted secret with an encrypted file
func (s *Secret) CreateWithFile(ciphertext string, hasPassphrase bool, expiresAt time.Time, revealOnce bool, fileIdentifier string, email string, webhookAddr string) (secretify.Secret, error) {
	file, err := s.fileSvc.Repo().ViewByIdentifier(s.db, fileIdentifier)
	if err != nil {
		return secretify.Secret{}, err
	}
	return s.Create(ciphertext, hasPassphrase, expiresAt, revealOnce, file.ID, email, webhookAddr)
}

func sendEmail() {
	mailjetClient := mailjet.NewMailjetClient("c4bf39bd29338413af37451bbd2ac507", "ebf4d1945f419fb055b003b157b126a0")
	messagesInfo := []mailjet.InfoMessagesV31{
		mailjet.InfoMessagesV31{
			From: &mailjet.RecipientV31{
				Email: "no-reply@secretify.io",
				Name:  "Secretify",
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: "dario.calovic@itdistrict.ch",
				},
			},
			Subject:  "Greetings from Mailjet.",
			TextPart: "My first Mailjet email",
			HTMLPart: "<h3>Dear passenger 1, welcome to <a href='https://www.mailjet.com/'>Mailjet</a>!</h3><br />May the delivery force be with you!",
			CustomID: "AppGettingStartedTest",
		},
	}
	messages := mailjet.MessagesV31{Info: messagesInfo}
	res, err := mailjetClient.SendMailV31(&messages)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Data: %+v\n", res)
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
		if secret.FileID > 0 {
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
	// TODO : send email
	// if s.cfgSvc.Policy().Mail.Enabled && secret.Email != "" {
	// 	var additionalInfo string
	// 	if deleted {
	// 		additionalInfo = "The secret was deleted."
	// 	}
	// 	go s.mailer.SendRevealedSecret(secret.Email, fmt.Sprintf("%s/s/%s", s.cfgSvc.Meta().UIURL, secret.Identifier), s.cfgSvc.Meta().Hoster.Name, s.cfgSvc.Meta().Hoster.Address, additionalInfo)
	// }

	// TODO : send webhook
	if s.cfgSvc.Policy().Webhook.Enabled && secret.WebhookAddr != "" {
		fmt.Println("Posting to webhook ", secret.WebhookAddr)
	}
	return
}

// Delete the secret
func (s *Secret) Delete(identifier string) error {
	secret, err := s.repo.ViewByIdentifier(s.db, identifier)
	if err != nil {
		return err
	}
	if secret.FileID > 0 {
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
		if secret.FileID > 0 {
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
