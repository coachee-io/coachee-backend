package mysql_test

import (
	"coachee-backend/internal/model"
	"coachee-backend/internal/repository"
	"coachee-backend/internal/repository/mysql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func testClient() *model.Customer {
	return &model.Customer{
		StripeID:  "stripeId",
		FirstName: "Lize",
		LastName:  "Viveiros",
		Email:     "email@email.com",
		BirthDate: time.Date(1991, 8, 29, 0, 0, 0, 0, time.UTC),
		Password:  "lepassword",
	}
}

func TestClientRepository_Create(t *testing.T) {
	db := NewDatabase(t)
	repo := mysql.NewClientRepository(db)
	defer db.Close()

	client := testClient()

	err := repo.Create(repository.DefaultNoTransaction, client)
	require.Nil(t, err)
}

func TestClientRepository_GetByID(t *testing.T) {
	db := NewDatabase(t)
	repo := mysql.NewClientRepository(db)
	defer db.Close()

	client := testClient()

	err := repo.Create(repository.DefaultNoTransaction, client)
	require.Nil(t, err)

	client2, err := repo.GetByID(repository.DefaultNoTransaction, client.ID)
	require.Nil(t, err)
	client2.UpdatedAt = client.UpdatedAt
	client2.CreatedAt = client.CreatedAt
	client2.BirthDate = client.BirthDate
	require.Equal(t, client, client2)
}

func TestClientRepository_GetByEmail(t *testing.T) {
	db := NewDatabase(t)
	repo := mysql.NewClientRepository(db)
	defer db.Close()

	client := testClient()

	err := repo.Create(repository.DefaultNoTransaction, client)
	require.Nil(t, err)

	client2, err := repo.GetByEmail(repository.DefaultNoTransaction, client.Email)
	require.Nil(t, err)
	client2.UpdatedAt = client.UpdatedAt
	client2.CreatedAt = client.CreatedAt
	client2.BirthDate = client.BirthDate
	require.Equal(t, client, client2)
}

func TestClientRepository_Update(t *testing.T) {
	db := NewDatabase(t)
	repo := mysql.NewClientRepository(db)
	defer db.Close()

	client := testClient()

	err := repo.Create(repository.DefaultNoTransaction, client)
	require.Nil(t, err)

	client.StripeID = "anotherStripeId"
	err = repo.Update(repository.DefaultNoTransaction, client)
	require.Nil(t, err)

	client2, err := repo.GetByID(repository.DefaultNoTransaction, client.ID)
	require.Nil(t, err)
	client2.UpdatedAt = client.UpdatedAt
	client2.CreatedAt = client.CreatedAt
	client2.BirthDate = client.BirthDate
	require.Equal(t, client, client2)
}
