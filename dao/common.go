package dao

import (
	"errors"
	"memo-server/utils"

	"gorm.io/gorm"
)

func Create[T any](data *T) error {
	return utils.DB.Create(&data).Error
}

func GetOne[T any](query string, args ...any) (data T, err error) {
	dbErr := utils.DB.Where(query, args...).First(&data).Error
	if dbErr != nil && !errors.Is(dbErr, gorm.ErrRecordNotFound) {
		err = dbErr
	}
	return data, err
}
