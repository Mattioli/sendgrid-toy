package providers

import (
	"encoding/json"
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

func (s *SendgridProvider) Send(data json.RawMessage) error {
	m := mail.NewV3Mail()
	err := json.Unmarshal(data, &m)
	if err != nil {
		return fmt.Errorf("could not unwrap json in sendgrid provider: %w", err)
	}

	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(m)

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
