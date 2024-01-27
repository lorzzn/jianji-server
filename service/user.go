package service

import (
	"memo-server/dao"
	"memo-server/entity"
	"memo-server/model/request"
	"memo-server/utils"
	"memo-server/utils/r"

	"github.com/gin-gonic/gin"
)

type User struct{}

func (*User) Signup(c *gin.Context) (code int, message string, data any) {
	params, _ := utils.GetRequestParams[request.Signup](c)
	if exist := checkUserExistByEmail(params.Email); exist {
		code = r.USER_EXISTED
		return
	}

	user := &entity.User{
		Name:   params.Name,
		Avatar: params.Avatar,
		Email:  params.Email,
	}
	err := dao.Create(&user)
	if err != nil {
		code = r.ERROR_DB_OPE
		return
	}
	code, message = r.OK, "注册成功"
	return
}

// 检查用户是否已经注册
func checkUserExistByEmail(email string) bool {
	existUser, _ := dao.GetOne[entity.User]("email = ?", email)
	return existUser.ID != 0
}
