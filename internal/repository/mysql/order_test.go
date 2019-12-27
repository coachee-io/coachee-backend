package mysql_test

import (
	"coachee-backend/internal/model"
	"coachee-backend/internal/repository"
	"coachee-backend/internal/repository/mysql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func testOrder() *model.Order {
	return &model.Order{
		CoachID:         1,
		CustomerID:      1,
		ProgramID:       "program_id",
		PaymentIntentID: "pi_id",
		Amount:          10000,
		TaxPercent:      2000,
		IntroCall:       time.Now(),
		Status:          model.OrderStatusCreated,
		Observations:    "I observe stuff",
	}
}

func TestOrderRepository_Create(t *testing.T) {
	db := NewDatabase(t)
	repo := mysql.NewOrderRepository(db)
	defer db.Close()

	order := testOrder()

	err := repo.Create(repository.DefaultNoTransaction, order)
	require.Nil(t, err)
}

func TestOrderRepository_GetByID(t *testing.T) {
	db := NewDatabase(t)
	repo := mysql.NewOrderRepository(db)
	defer db.Close()

	order := testOrder()

	err := repo.Create(repository.DefaultNoTransaction, order)
	require.Nil(t, err)

	order2, err := repo.GetByID(repository.DefaultNoTransaction, order.ID)
	require.Nil(t, err)
	order2.UpdatedAt = order.UpdatedAt
	order2.CreatedAt = order.CreatedAt
	order2.IntroCall = order.IntroCall
	require.Equal(t, order, order2)
}

func TestOrderRepository_Update(t *testing.T) {
	db := NewDatabase(t)
	repo := mysql.NewOrderRepository(db)
	defer db.Close()

	order := testOrder()

	err := repo.Create(repository.DefaultNoTransaction, order)
	require.Nil(t, err)

	order.Observations = "new observations"
	err = repo.Update(repository.DefaultNoTransaction, order)
	require.Nil(t, err)

	order2, err := repo.GetByID(repository.DefaultNoTransaction, order.ID)
	require.Nil(t, err)
	order2.UpdatedAt = order.UpdatedAt
	order2.CreatedAt = order.CreatedAt
	order2.IntroCall = order.IntroCall
	require.Equal(t, order, order2)
}
