package mysql_test

import (
	"coachee-backend/internal/model"
	"coachee-backend/internal/repository"
	"coachee-backend/internal/repository/mysql"
	"testing"

	"github.com/stretchr/testify/require"
)

func testCustomer() *model.Customer {
	return &model.Customer{
		StripeID:  "stripeId",
		FirstName: "Lize",
		LastName:  "Viveiros",
		Email:     "email@email.com",
		Password:  "lepassword",
	}
}

func TestCustomerRepository_Create(t *testing.T) {
	db := NewDatabase(t)
	repo := mysql.NewCustomerRepository(db)
	defer db.Close()

	client := testCustomer()

	err := repo.Create(repository.DefaultNoTransaction, client)
	require.Nil(t, err)
}

func TestCustomerRepository_GetByID(t *testing.T) {
	db := NewDatabase(t)
	repo := mysql.NewCustomerRepository(db)
	defer db.Close()

	client := testCustomer()

	err := repo.Create(repository.DefaultNoTransaction, client)
	require.Nil(t, err)

	client2, err := repo.GetByID(repository.DefaultNoTransaction, client.ID)
	require.Nil(t, err)
	client2.UpdatedAt = client.UpdatedAt
	client2.CreatedAt = client.CreatedAt
	require.Equal(t, client, client2)
}

func TestCustomerRepository_GetByEmail(t *testing.T) {
	db := NewDatabase(t)
	repo := mysql.NewCustomerRepository(db)
	defer db.Close()

	client := testCustomer()

	err := repo.Create(repository.DefaultNoTransaction, client)
	require.Nil(t, err)

	client2, err := repo.GetByEmail(repository.DefaultNoTransaction, client.Email)
	require.Nil(t, err)
	client2.UpdatedAt = client.UpdatedAt
	client2.CreatedAt = client.CreatedAt
	require.Equal(t, client, client2)
}

func TestCustomerRepository_Update(t *testing.T) {
	db := NewDatabase(t)
	repo := mysql.NewCustomerRepository(db)
	defer db.Close()

	client := testCustomer()

	err := repo.Create(repository.DefaultNoTransaction, client)
	require.Nil(t, err)

	client.StripeID = "anotherStripeId"
	err = repo.Update(repository.DefaultNoTransaction, client)
	require.Nil(t, err)

	client2, err := repo.GetByID(repository.DefaultNoTransaction, client.ID)
	require.Nil(t, err)
	client2.UpdatedAt = client.UpdatedAt
	client2.CreatedAt = client.CreatedAt
	require.Equal(t, client, client2)
}
