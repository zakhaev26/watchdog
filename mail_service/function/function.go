package function

import (
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
)

const (
	smtpAuthAddress   string = "smtp.gmail.com"
	smtpServerAddress        = "smtp.gmail.com:587"
)

// EmailSender is an interface for sending emails.
type EmailSender interface {
	SendEmail(subject, content string, to, cc, bcc, attachFiles []string) error
}

// GmailSender is a type that implements the EmailSender interface.
type GmailSender struct {
	name              string
	fromEmailAddress  string
	fromEmailPassword string
}

// NewGmailSender creates a new GmailSender instance.
func NewGmailSender(name, fromEmailAddress, fromEmailPassword string) EmailSender {
	return &GmailSender{
		name:              name,
		fromEmailAddress:  fromEmailAddress,
		fromEmailPassword: fromEmailPassword,
	}
}

// SendEmail is the implementation of the SendEmail method for GmailSender.
func (sender *GmailSender) SendEmail(subject, content string, to, cc, bcc, attachFiles []string) error {

	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", sender.name, sender.fromEmailAddress)
	e.Subject = subject
	e.HTML = []byte(content)
	e.To = to
	e.Cc = cc
	e.Bcc = bcc

	for _, f := range attachFiles {
		_, err := e.AttachFile(f)

		if err != nil {
			return fmt.Errorf("Failed to attach the file %s:%w", f, err)
		}
	}

	/*PlainAuth returns an Auth that implements the PLAIN authentication mechanism as defined in RFC 4616. The returned Auth uses the given username and password to authenticate to host and act as identity. Usually identity should be the empty string, to act as username.

	PlainAuth will only send the credentials if the connection is using TLS or is connected to localhost. Otherwise authentication will fail with an error, without sending the credentials.*/

	smtpAuth := smtp.PlainAuth("", sender.fromEmailAddress, sender.fromEmailPassword, smtpAuthAddress)
	fmt.Println("OKOK")

	return e.Send(smtpServerAddress, smtpAuth)
}
