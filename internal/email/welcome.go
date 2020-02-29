package email

import (
	"bytes"
	"strings"
)

const (
	confirmPath = "/fakepath/{id}/fake"
	welcomeSub  = "Welcome to coachee.io!"
)

// SendWelcomeEmail sends the welcome email to new clients
func (c *Client) SendWelcomeEmail(to, token string) error {
	url := strings.Replace(c.hostname+confirmPath, "{id}", token, 1)
	templateData := struct {
		PrimaryURL   string
		SecondaryURL string
	}{
		PrimaryURL:   url,
		SecondaryURL: c.hostname,
	}

	templ, err := c.getWelcomeTemplate()
	if err != nil {
		return err
	}

	buf := &bytes.Buffer{}
	err = templ.Execute(buf, &templateData)
	if err != nil {
		return err
	}
	body := buf.String()

	return c.sendEmail(to, welcomeSub, body)
}
