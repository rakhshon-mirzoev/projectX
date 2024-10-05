package repository

import (
	"errors"

	"gorm.io/gorm"
)

func translateError(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return gorm.ErrRecordNotFound
	}
	return err
}
