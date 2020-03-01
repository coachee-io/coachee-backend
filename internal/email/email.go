package email

import (
	"context"
	"html/template"
	"net/smtp"

	"github.com/caarlos0/env"
)

const (
	confirmBooking = "confirm-booking.html"
	copyright      = "copyright.html"
	footer         = "footer.html"
	forgotPassword = "forgot-password.html"
	header         = "header.html"
	logo           = "logo.html"
	welcome        = "welcome.html"
)

type config struct {
	Path     string `env:"EMAIL_PATH" envDefault:"/web/tmpl/"`
	Username string `env:"EMAIL_USERNAME,required"`
	Password string `env:"EMAIL_PASSWORD,required"`
	Host     string `env:"EMAIL_HOST" envDefault:"smtp.gmail.com"`
	Address  string `env:"EMAIL_ADDRESS" envDefault:"smtp.gmail.com:587"`
}

type Client struct {
	appCtx                    context.Context
	config                    config
	auth                      smtp.Auth
	hostname                  string
	getWelcomeTemplate        func() (*template.Template, error)
	getConfirmBookingTemplate func() (*template.Template, error)
	getForgotPasswordTemplate func() (*template.Template, error)
}

// NewClient initializes a new email client
func NewClient(ctx context.Context, hostname string) (*Client, error) {
	client := &Client{
		appCtx:   ctx,
		hostname: hostname,
	}

	err := env.Parse(&client.config)
	if err != nil {
		return nil, err
	}

	client.auth = smtp.PlainAuth("", client.config.Username, client.config.Password, client.config.Host)

	confirmBookingPath := client.config.Path + confirmBooking
	copyrightPath := client.config.Path + copyright
	footerPath := client.config.Path + footer
	forgotPasswordPath := client.config.Path + forgotPassword
	headerPath := client.config.Path + header
	logoPath := client.config.Path + logo
	welcomePath := client.config.Path + welcome

	welcomeTemplate, err := template.ParseFiles(welcomePath, headerPath, logoPath, copyrightPath, footerPath)
	if err != nil {
		return nil, err
	}
	client.getWelcomeTemplate = func() (*template.Template, error) {
		return welcomeTemplate.Clone()
	}

	confirmTemplate, err := template.ParseFiles(confirmBookingPath, headerPath, logoPath, copyrightPath, footerPath)
	if err != nil {
		return nil, err
	}
	client.getConfirmBookingTemplate = func() (*template.Template, error) {
		return confirmTemplate.Clone()
	}

	forgotPasswordTemplate, err := template.ParseFiles(forgotPasswordPath, headerPath, logoPath, copyrightPath, footerPath)
	if err != nil {
		return nil, err
	}
	client.getForgotPasswordTemplate = func() (*template.Template, error) {
		return forgotPasswordTemplate.Clone()
	}

	return client, nil
}

func (c *Client) sendEmail(to, subject, body string) error {
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	toH := "To: " + to + "\n"
	sub := "Subject: " + subject + "\n"
	msg := []byte(sub + toH + mime + body)

	return smtp.SendMail(c.config.Address, c.auth, c.config.Username, []string{to}, msg)
}
