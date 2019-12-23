package mysql

import (
	"coachee-backend/gen/coachee"
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
)

// ErrExists record already exists - error prefix
const ErrExists = "Error 1062"

// parseError returns either a repository or a generic error
func parseError(err error) error {
	if err != nil {
		if strings.Contains(err.Error(), ErrExists) {
			return coachee.MakeValidation(fmt.Errorf("record already exists"))
		} else if gorm.IsRecordNotFoundError(err) {
			return coachee.MakeNotFound(fmt.Errorf("record not found"))
		}
		return coachee.MakeTransient(err)
	}
	return nil
}
