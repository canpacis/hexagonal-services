package email

type Service interface {
	SendEmail(*SendEmailConfig) error
}

type SendEmailConfig struct {
	// Configuration
}

// SES implementation of the email service
type SES struct{}

func (*SES) SendEmail(*SendEmailConfig) error {
	return nil
}

// Mailgun implementation of the email service
type Mailgun struct{}

func (*Mailgun) SendEmail(*SendEmailConfig) error {
	return nil
}
