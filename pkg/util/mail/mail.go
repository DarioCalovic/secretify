package mail

import (
	"bytes"
	_ "embed"
	"html/template"

	mailjet "github.com/mailjet/mailjet-apiv3-go"
)

//go:embed secretCreatedHTML.html
var secretCreatedHTML string

//go:embed secretCreatedText.txt
var secretCreatedText string

//go:embed secretRevealedHTML.html
var secretRevealedHTML string

//go:embed secretRevealedText.html
var secretRevealedText string

type Mailer struct {
	smtp struct {
		publicKey  string
		privateKey string
	}
}

func NewMailer(publicKey string, privateKey string) *Mailer {
	m := &Mailer{}
	m.smtp.publicKey = publicKey
	m.smtp.privateKey = privateKey
	return m
}

func (m *Mailer) Send(recipient, subject, htmlBody, textBody, customID string) error {
	mailjetClient := mailjet.NewMailjetClient(m.smtp.publicKey, m.smtp.privateKey)
	messagesInfo := []mailjet.InfoMessagesV31{
		mailjet.InfoMessagesV31{
			From: &mailjet.RecipientV31{
				Email: "no-reply@secretify.io",
				Name:  "Secretify",
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: recipient,
				},
			},
			Subject:  subject,
			TextPart: textBody,
			HTMLPart: htmlBody,
			CustomID: customID,
		},
	}
	messages := mailjet.MessagesV31{Info: messagesInfo}
	_, err := mailjetClient.SendMailV31(&messages)
	if err != nil {
		return err
	}
	return nil
}

func SecretCreatedHTMLBody(link, hosterName, hosterAddress string) (string, error) {
	var body bytes.Buffer
	t, err := template.New("foo").Parse(secretCreatedHTML)
	if err != nil {
		return "", err
	}
	t.Execute(&body, struct {
		Link          string
		HosterName    string
		HosterAddress string
	}{
		Link:          link,
		HosterName:    hosterName,
		HosterAddress: hosterAddress,
	})
	return body.String(), nil
}

func SecretCreatedTextBody(link, hosterName, hosterAddress string) (string, error) {
	var body bytes.Buffer
	t, err := template.New("foo").Parse(secretCreatedText)
	if err != nil {
		return "", err
	}
	t.Execute(&body, struct {
		Link          string
		HosterName    string
		HosterAddress string
	}{
		Link:          link,
		HosterName:    hosterName,
		HosterAddress: hosterAddress,
	})
	return body.String(), nil
}

func SecretRevealedHTMLBody(link, hosterName, hosterAddress string) (string, error) {
	var body bytes.Buffer
	t, err := template.New("foo").Parse(secretRevealedHTML)
	if err != nil {
		return "", err
	}
	t.Execute(&body, struct {
		Link          string
		HosterName    string
		HosterAddress string
	}{
		Link:          link,
		HosterName:    hosterName,
		HosterAddress: hosterAddress,
	})
	return body.String(), nil
}

func SecretRevealedTextBody(link, hosterName, hosterAddress string) (string, error) {
	var body bytes.Buffer
	t, err := template.New("foo").Parse(secretRevealedText)
	if err != nil {
		return "", err
	}
	t.Execute(&body, struct {
		Link          string
		HosterName    string
		HosterAddress string
	}{
		Link:          link,
		HosterName:    hosterName,
		HosterAddress: hosterAddress,
	})
	return body.String(), nil
}
