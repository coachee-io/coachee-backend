package mysql_test

import (
	"coachee-backend/internal/model"
	"coachee-backend/internal/repository"
	"coachee-backend/internal/repository/mysql"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func testRecovery() *model.Recovery {
	return &model.Recovery{
		CustomerID: 1,
	}
}

func TestRecoveryRepository_Create(t *testing.T) {
	db := NewDatabase(t)
	repo := mysql.NewRecoveryRepository(db)
	defer db.Close()

	recovery := testRecovery()

	err := repo.Create(repository.DefaultNoTransaction, recovery)
	require.Nil(t, err)
}

func TestRecoveryRepository_GetByID(t *testing.T) {
	db := NewDatabase(t)
	repo := mysql.NewRecoveryRepository(db)
	defer db.Close()

	recovery := testRecovery()

	err := repo.Create(repository.DefaultNoTransaction, recovery)
	require.Nil(t, err)
	fmt.Println(recovery)

	recovery2, err := repo.GetByID(repository.DefaultNoTransaction, recovery.ID)
	require.Nil(t, err)
	recovery2.CreatedAt = recovery.CreatedAt
	require.Equal(t, recovery, recovery2)
}
