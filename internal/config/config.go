package config

import (
	"github.com/caarlos0/env"
)

type Config struct {
	Email Email
	Main  Main
}

type Email struct {
	Path      string `env:"EMAIL_PATH" envDefault:"/web/tmpl/"`
	Key       string `env:"SENDGRID_API_KEY" envDefault:"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"`
	FromEmail string `env:"FROM_EMAIL" envDefault:"admin@coachee.io"`
	FromName  string `env:"FROM_NAME" envDefault:"coachee.io"`
	HostName  string `env:"COACHEE_HOST_NAME" envDefault:"https://flamboyant-bohr-46e743.netlify.app"`
}

type Main struct {
	StripeKey     string `env:"STRIPE_KEY" envDefault:"sk_test_yKV7Mo9kSpokxpFvwxKRtbyd00knjXTpJh"`
	PubKey        string `env:"STRIPE_PUB_KEY" envDefault:"pk_test_bmGuB7UJfIeeeofOouGHeJcd00MQjvjYVL"`
	AdminEmail    string `env:"ADMIN_EMAIL" envDefault:"test@test.com"`
	AdminPassword string `env:"ADMIN_PASSWORD" envDefault:"rucalindo19"`
	SlackWebhook  string `env:"SLACK_WEBHOOK,required"`
}

// Parse parses the environment variables into the config struct
func (c *Config) Parse() error {
	err := env.Parse(&c.Email)
	if err != nil {
		return err
	}

	err = env.Parse(&c.Main)
	if err != nil {
		return err
	}

	return nil
}
