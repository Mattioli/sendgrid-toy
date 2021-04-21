package providers

type EmailProvider interface {
	Send(fromName, fromEmail, toName, toEmail, subject, msg string) error
	Name() string
}
