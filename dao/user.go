package dao

import (
	"jianji-server/entity"
	"jianji-server/utils"
)

type User struct{}

func (*User) GetUserById(id int) (res entity.User) {
	utils.DB.Table("user").
		Where("id = ? AND status = 1", id).
		First(&res)
	return
}

func (*User) UpdateUserById(id int, updated entity.User) (res entity.User) {
	utils.DB.Table("user").
		Where("id = ? AND status = 1", id).
		Updates(updated).
		Scan(&res)
	return
}
