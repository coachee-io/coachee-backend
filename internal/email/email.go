package email

import (
	"coachee-backend/gen/coachee"
	"context"
	"errors"
	"fmt"
	"html/template"

	"github.com/sendgrid/sendgrid-go/helpers/mail"

	"github.com/sendgrid/sendgrid-go"

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
	Path      string `env:"EMAIL_PATH" envDefault:"/web/tmpl/"`
	Key       string `env:"SENDGRID_API_KEY" envDefault:"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"`
	FromEmail string `env:"FROM_EMAIL" envDefault:"admin@coachee.io"`
	FromName  string `env:"FROM_NAME" envDefault:"coachee.io"`
}

type Client struct {
	appCtx                    context.Context
	config                    config
	sendgrid                  *sendgrid.Client
	hostname                  string
	from                      *mail.Email
	getWelcomeTemplate        func() (*template.Template, error)
	getConfirmBookingTemplate func() (*template.Template, error)
	getForgotPasswordTemplate func() (*template.Template, error)
}

// NewClient initializes a new email client
func NewClient(ctx context.Context, hostname string) (*Client, error) {
	client := &Client{
		appCtx: ctx,
	}

	err := env.Parse(&client.config)
	if err != nil {
		return nil, err
	}

	client.sendgrid = sendgrid.NewSendClient(client.config.Key)
	client.hostname = hostname
	client.from = mail.NewEmail(client.config.FromName, client.config.FromEmail)

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
	toEmail := mail.NewEmail("", to)
	html := mail.NewContent("text/html", body)
	message := mail.NewV3MailInit(c.from, subject, toEmail, html)
	res, err := c.sendgrid.Send(message)
	if err != nil {
		return err
	}
	if res.StatusCode > 299 {
		err = coachee.MakeValidation(errors.New(fmt.Sprintf("Code: %d Error: %s", res.StatusCode, res.Body)))
		return err
	}
	fmt.Println(res)
	return nil
}
