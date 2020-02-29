package email

import (
	"bytes"
	"strings"
)

const (
	clientRecoveryPath = "/clientrecovery/{id}/recovery"
	coachRecoveryPath  = "/coachrecovery/{id}/recovery"
	recoverySub        = "Password recovery for coachee.io"
)

// SendClientPasswordRecoveryEmail sends a password recovery email for a client
func (c *Client) SendClientPasswordRecoveryEmail(to, token string) error {
	return c.sendPasswordRecoveryEmail(to, clientRecoveryPath, token)
}

// SendCoachPasswordRecoveryEmail sends a password recovery email for a coach
func (c *Client) SendCoachPasswordRecoveryEmail(to, token string) error {
	return c.sendPasswordRecoveryEmail(to, coachRecoveryPath, token)
}

// sendPasswordRecoveryEmail sends a password recovery email
func (c *Client) sendPasswordRecoveryEmail(to, urlSuffix, token string) error {
	url := strings.Replace(c.hostname+urlSuffix, "{id}", token, 1)
	templateData := struct {
		ResetPassworURL string
	}{
		ResetPassworURL: url,
	}

	templ, err := c.getForgotPasswordTemplate()
	if err != nil {
		return err
	}

	buf := &bytes.Buffer{}
	err = templ.Execute(buf, &templateData)
	if err != nil {
		return err
	}
	body := buf.String()

	return c.sendEmail(to, recoverySub, body)
}
