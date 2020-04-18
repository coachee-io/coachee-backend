package mysql

import (
	"coachee-backend/internal/model"
	"coachee-backend/internal/repository"

	"github.com/jinzhu/gorm"
)

// CoachRepository is the repository to access and persist coaches
type CoachRepository struct {
	db *gorm.DB
}

// NewCoachRepository initializes a CoachRepository
func NewCoachRepository(db *gorm.DB) *CoachRepository {
	return &CoachRepository{db: db}
}

// Begin starts a new Transaction
func (r CoachRepository) Begin() repository.Transaction {
	return newTransaction(r.db)
}

// Create persists a coach
func (r CoachRepository) Create(transaction repository.Transaction, coach *model.Coach) error {
	tx := r.checkTransaction(transaction)

	// create customer
	if err := tx.Create(coach).Error; err != nil {
		return parseError(err)
	}

	return nil
}

// GetByID returns a coach by id
func (r CoachRepository) GetByID(transaction repository.Transaction, id uint) (*model.Coach, error) {
	tx := r.checkTransaction(transaction)

	var coach model.Coach
	if err := tx.First(&coach, id).Error; err != nil {
		return nil, parseError(err)
	}

	return &coach, nil
}

// GetByEmail returns a coach by id
func (r CoachRepository) GetByEmail(transaction repository.Transaction, email string) (*model.Coach, error) {
	tx := r.checkTransaction(transaction)

	var coach model.Coach
	if err := tx.Where("email = ?", email).First(&coach).Error; err != nil {
		return nil, parseError(err)
	}

	return &coach, nil
}

// GetByPage returns a page of coaches by a tag with pagination
func (r CoachRepository) GetByPage(transaction repository.Transaction, tag string, limit, page uint, showAll bool) ([]*model.Coach, error) {
	tx := r.checkTransaction(transaction)

	tag = "%" + tag + "%"

	offset := page * limit

	var coaches []*model.Coach
	var err error
	if showAll {
		err = tx.Where("tags LIKE ?", tag).Limit(limit).Offset(offset).Find(&coaches).Error
	} else {
		err = tx.Where("tags LIKE ? AND status = ?", tag, string(model.StatusActive)).Limit(limit).Offset(offset).Find(&coaches).Error
	}

	if err != nil {
		return nil, parseError(err)
	}

	return coaches, nil
}

// Update updates the changed fields in the db
func (r CoachRepository) Update(transaction repository.Transaction, coach *model.Coach) error {
	tx := r.checkTransaction(transaction)

	if err := tx.Model(coach).Update(coach).Error; err != nil {
		return parseError(err)
	}

	return nil
}

// Length gives the number of coaches available
func (r CoachRepository) Length(transaction repository.Transaction, tag string) (uint, error) {
	tx := r.checkTransaction(transaction)

	tag = "%" + tag + "%"

	var coaches []*model.Coach
	var count uint
	if err := tx.Where("tags LIKE ?", tag).Find(&coaches).Count(&count).Error; err != nil {
		return 0, parseError(err)
	}

	return count, nil
}

// checkTransaction returns either a gorm.DB pointer either as a transaction or as a direct db access
// this depends on the if the repository.Transaction is a transaction or a NoTransaction
func (r CoachRepository) checkTransaction(in repository.Transaction) *gorm.DB {
	tx, ok := in.(*transaction)
	if !ok {
		return r.db
	}
	return tx.db
}
