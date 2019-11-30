package connector_test

import (
	"coachee-backend/internal/repository/mysql/connector"
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConnect(t *testing.T) {
	// required env variables
	os.Setenv("DB_USER", "coachee_user")
	os.Setenv("DB_PASSWORD", "coachee_pass")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_NAME", "coachee_db")
	ctx := context.Background()

	sql, err := connector.Connect(ctx)
	require.Nil(t, err)
	require.NotNil(t, sql)
}
