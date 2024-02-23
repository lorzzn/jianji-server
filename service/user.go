package service

import (
	"memo-server/dao"
	"memo-server/entity"
	"memo-server/model/request"
	"memo-server/model/response"
	"memo-server/utils"
	"memo-server/utils/r"

	"github.com/gin-gonic/gin"
)

type User struct{}

func (*User) Login(c *gin.Context) (code int, message string, data response.Login) {
	params, _ := utils.GetRequestParams[request.Login](c)
	//如果用户已经注册
	if exist := getUserByEmail(params.Email); exist.ID != 0 {
		code = r.USER_EXISTED
		data = response.Login{UserInfo: exist}
		return
	}

	user := &entity.User{
		Email: params.Email,
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
func getUserByEmail(email string) entity.User {
	existUser, _ := dao.GetOne[entity.User]("email = ?", email)
	return existUser
}
