package dao

import (
	"memo-server/model/response"
	"memo-server/utils"
)

type User struct{}

func (*User) GetUserById(id int) (res response.User) {
	utils.DB.Table("user").
		Where("id = ? AND status = 1", id).
		First(&res)
	return
}
