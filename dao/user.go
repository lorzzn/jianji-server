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
