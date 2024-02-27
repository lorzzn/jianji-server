package dao

import (
	"jianji-server/entity"
	"jianji-server/utils"
)

type User struct{}

func (*User) GetUserById(id uint64) (res entity.User) {
	utils.DB.Table("user").
		Where("id = ? AND status = 1", id).
		First(&res)
	return
}

func (*User) UpdateUserById(id uint64, updated entity.User) (res entity.User) {
	utils.DB.Table("user").
		Where("id = ? AND status = 1", id).
		Updates(updated).
		Scan(&res)
	return
}
