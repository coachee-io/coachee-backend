package mysql_test

import (
	"coachee-backend/internal/model"
	"coachee-backend/internal/repository/mysql/connector"
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

// NewDatabase returns a new database for testing
func NewDatabase(t *testing.T) *gorm.DB {
	// required env variables
	os.Setenv("DB_USER", "coachee_user")
	os.Setenv("DB_PASSWORD", "coachee_pass")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_NAME", "coachee_db")
	ctx := context.Background()

	sql, err := connector.Connect(ctx)
	require.Nil(t, err)

	// clean database
	sql.DropTableIfExists(
		model.Coach{},
		model.Client{},
	)

	sql.AutoMigrate(
		model.Coach{},
		model.Client{},
	)
	return sql
}

// Test_Migration tests running the database migrations using gorm (for getting the sql query and copy it to goose)
func Test_Migration(t *testing.T) {
	assert.Nil(t, NewDatabase(t).Error)
}
