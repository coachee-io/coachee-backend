package mysql_test

import (
	"coachee-backend/internal/model"
	"coachee-backend/internal/repository"
	"coachee-backend/internal/repository/mysql"
	"testing"

	"github.com/stretchr/testify/require"
)

func testNewsletter() *model.Newsletter {
	return &model.Newsletter{
		Email: "test@test.com",
	}
}

func TestNewNewsletterRepository_Create(t *testing.T) {
	db := NewDatabase(t)
	repo := mysql.NewNewsletterRepository(db)
	defer db.Close()

	newsletter := testNewsletter()

	err := repo.Create(repository.DefaultNoTransaction, newsletter)
	require.Nil(t, err)
}
