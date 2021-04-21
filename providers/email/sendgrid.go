package providers

import (
	"fmt"
	"log"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type SendgridProvider struct{}

func NewSendGridProvider() *SendgridProvider {
	return &SendgridProvider{}
}

func (s *SendgridProvider) Send(fromName, fromEmail, toName, toEmail, subject, msg string) error {
	from := mail.NewEmail(fromName, fromEmail)
	to := mail.NewEmail(toName, toEmail)

	plainTextContent := msg
	htmlContent := "<strong>" + msg + "</strong>"

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)

	if err != nil {
		log.Println(err)
		return err
	}

	fmt.Println(response.StatusCode)
	fmt.Println(response.Body)
	fmt.Println(response.Headers)

	return nil
}

func (s *SendgridProvider) Name() string {
	return "sendgrid"
}
