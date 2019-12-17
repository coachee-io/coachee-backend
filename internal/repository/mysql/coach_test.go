package mysql_test

import (
	"coachee-backend/internal/model"
	"coachee-backend/internal/repository"
	"coachee-backend/internal/repository/mysql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func testCoach() *model.Coach {
	return &model.Coach{
		ID:          0,
		FirstName:   "Elie",
		LastName:    "Viveiros",
		Email:       "lize@gmail.com",
		Phone:       "094712747127",
		StripeID:    "stripe_id",
		Tags:        "health",
		Description: "o meu amor",
		City:        "London",
		Country:     "UK",
		PictureUrl:  "https://nationalpostcom.files.wordpress.com/2019/02/gbr_21_blackwhitebearsiblings.jpg?quality=80&strip=all&w=780&zoom=2",
		Status:      model.StatusActive,
		Vat:         "le vat number",
		IntroCall:   time.Now(),
		Availability: model.Availabilities{
			{
				ID:    "leid",
				Day:   0,
				Start: 1000,
				End:   1440,
			},
		},
		Certifications: model.Certifications{
			{
				ID:           "leid2",
				Title:        "Joao's GF",
				Description:  "Exemplar GF",
				Institution:  "Joao",
				DateAcquired: time.Date(2008, 1, 0, 0, 0, 0, 0, time.UTC),
			},
		},
		Programs: model.Programs{
			{
				ID:               "leid3",
				Name:             "Best girlfriend course",
				NumberOfSessions: 12,
				Duration:         60,
				Description:      "you can also be the best girlfriend",
				TotalPrice:       10000000,
				TaxPercent:       2000,
			},
		},
		TextAvailability:   "text Availability",
		TextCertifications: "text Certifications",
		TextPrograms:       "text Programs",
	}
}

func TestCoachRepository_Create(t *testing.T) {
	db := NewDatabase(t)
	repo := mysql.NewCoachRepository(db)
	defer db.Close()

	coach := testCoach()

	err := repo.Create(repository.DefaultNoTransaction, coach)
	require.Nil(t, err)

	coach2, err := repo.GetByID(repository.DefaultNoTransaction, coach.ID)
	require.Nil(t, err)
	coach2.CreatedAt = coach.CreatedAt //slight diff
	coach2.UpdatedAt = coach.UpdatedAt
	coach2.IntroCall = coach.IntroCall
	require.Equal(t, *coach, *coach2)
}

func TestCoachRepository_GetByPage(t *testing.T) {
	db := NewDatabase(t)
	repo := mysql.NewCoachRepository(db)
	defer db.Close()

	coach := testCoach()
	err := repo.Create(repository.DefaultNoTransaction, coach)
	require.Nil(t, err)

	coach2 := testCoach()
	err = repo.Create(repository.DefaultNoTransaction, coach2)
	require.Nil(t, err)

	coach3 := testCoach()
	coach3.Tags = "life"
	err = repo.Create(repository.DefaultNoTransaction, coach3)
	require.Nil(t, err)

	coaches, err := repo.GetByPage(repository.DefaultNoTransaction, coach.Tags, 100, 0)
	require.Nil(t, err)
	require.Len(t, coaches, 2)

	coaches, err = repo.GetByPage(repository.DefaultNoTransaction, coach.Tags, 1, 1)
	require.Nil(t, err)
	require.Len(t, coaches, 1)
	require.Equal(t, coach2.ID, coaches[0].ID)

	coaches, err = repo.GetByPage(repository.DefaultNoTransaction, "life", 1, 0)
	require.Nil(t, err)
	require.Len(t, coaches, 1)
	require.Equal(t, coach3.ID, coaches[0].ID)

	coaches, err = repo.GetByPage(repository.DefaultNoTransaction, "life", 1, 1)
	require.Nil(t, err)
	require.Len(t, coaches, 0)
}

func TestCoachRepository_Update(t *testing.T) {
	db := NewDatabase(t)
	repo := mysql.NewCoachRepository(db)
	defer db.Close()

	coach := testCoach()
	err := repo.Create(repository.DefaultNoTransaction, coach)
	require.Nil(t, err)

	coach.Tags = "life"

	err = repo.Update(repository.DefaultNoTransaction, coach)
	require.Nil(t, err)

	coach2, err := repo.GetByID(repository.DefaultNoTransaction, coach.ID)
	require.Nil(t, err)
	require.Equal(t, "life", coach2.Tags)
}

func TestCoachRepository_Length(t *testing.T) {
	db := NewDatabase(t)
	repo := mysql.NewCoachRepository(db)
	defer db.Close()

	coach := testCoach()

	err := repo.Create(repository.DefaultNoTransaction, coach)
	require.Nil(t, err)

	coach2, err := repo.GetByID(repository.DefaultNoTransaction, coach.ID)
	require.Nil(t, err)
	coach2.CreatedAt = coach.CreatedAt //slight diff
	coach2.UpdatedAt = coach.UpdatedAt
	coach2.IntroCall = coach.IntroCall
	require.Equal(t, *coach, *coach2)

	count, err := repo.Length(repository.DefaultNoTransaction, coach.Tags)
	require.Nil(t, err)
	require.Equal(t, uint(1), count)
}
