package email

import "bytes"

const (
	learnMore  = "coachee.io/learnmore"
	bookingSub = "Your session is booked!"
)

// SendWelcomeEmail sends the welcome email to new clients
func (c *Client) SendBookingEmail(to, programme, coachName string) error {
	templateData := struct {
		CoachingProgramme string
		CoachName         string
		LearnMore         string
	}{
		CoachingProgramme: programme,
		CoachName:         coachName,
		LearnMore:         learnMore,
	}

	templ, err := c.getConfirmBookingTemplate()
	if err != nil {
		return err
	}

	buf := &bytes.Buffer{}
	err = templ.Execute(buf, &templateData)
	if err != nil {
		return err
	}
	body := buf.String()

	return c.sendEmail(to, bookingSub, body)
}
