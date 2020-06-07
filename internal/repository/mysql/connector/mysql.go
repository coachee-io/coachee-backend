package connector

import (
	"context"
	"fmt"
	"time"

	"github.com/caarlos0/env"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// config defines the env variables necessary to setup the connector
type config struct {
	Debug         bool          `env:"DEBUG" envDefault:"false"`
	RetryInterval time.Duration `env:"DB_RETRY_INTERVAL" envDefault:"1s"`
	RetryAttempts int           `env:"DB_RETRY_ATTEMPTS" envDefault:"6"`
	User          string        `env:"DB_USER,required"`
	Password      string        `env:"DB_PASSWORD,required"`
	Host          string        `env:"DB_HOST,required"`
	Port          int           `env:"DB_PORT" envDefault:"3306"`
	Database      string        `env:"DB_NAME,required"`
	Url           string        `env:"DB_URL"`
}

func Connect(ctx context.Context) (*gorm.DB, error) {
	cfg := config{}

	// gets the config from the env
	if err := env.Parse(&cfg); err != nil {
		return nil, fmt.Errorf("failed to parse env configuration: %s", err)
	}

	// sample url: billing:perkbox@tcp(mysql:3306)/billing?parseTime=true
	if cfg.Url == "" {
		cfg.Url = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	}
	// all fields are required, port defaults to 3306
	if cfg.Url == ":tcp(:3306)/?parseTime=true" {
		return nil, fmt.Errorf("missing required url for mysql connection")
	}

	db, err := gorm.Open("mysql", cfg.Url)
	if err != nil {
		return nil, err
	}

	db = db.LogMode(cfg.Debug)

	return db, err
}
