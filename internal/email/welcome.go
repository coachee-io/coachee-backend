package email

import (
	"bytes"
)

const (
	welcomeSub = "Welcome to coachee.io!"
)

// SendWelcomeEmail sends the welcome email to new clients
func (c *Client) SendWelcomeEmail(to, token string) error {
	templateData := struct {
		SecondaryURL string
	}{
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
