// +build email

package email_test

import (
	"coachee-backend/internal/email"
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"

	"github.com/stretchr/testify/require"
)

func getTestEmailClient(t *testing.T) *email.Client {
	_ = os.Setenv("EMAIL_PATH", "../../web/tmpl/")

	cli, err := email.NewClient(context.Background(), "localhost")
	require.Nil(t, err)

	return cli
}

func TestNewClient(t *testing.T) {
	_ = getTestEmailClient(t)
}

func TestEmailService(t *testing.T) {
	_ = os.Setenv("SENDGRID_API_KEY", "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")

	from := mail.NewEmail("Example User", "admin@coachee.io")
	subject := "Sending with SendGrid is Fun"
	to := mail.NewEmail("Example User", "joca14@gmail.com")
	plainTextContent := "and easy to do anywhere, even with Go"
	htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}
