package service

import (
	"jianji-server/dao"
	"jianji-server/entity"
	"jianji-server/model/request"
	"jianji-server/model/response"
	"jianji-server/utils"
	"jianji-server/utils/r"

	"github.com/cstockton/go-conv"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
)

type User struct{}

func (*User) EditProfile(c *gin.Context) (code int, message string, data *response.Profile) {
	userId, ok := c.Get("UserId")
	if !ok {
		code = r.USER_NOT_EXISTED
		return
	}

	params, _ := utils.GetRequestParams[request.EditProfile](c)

	updated := &request.EditProfile{
		Name: params.Name,
	}

	var e entity.User
	mapstructure.Decode(updated, &e)
	i, _ := conv.Int(userId)
	data = &response.Profile{
		UserInfo: dao.UserDao.UpdateUserById(i, e),
	}

	return
}

func (*User) GetProfile(c *gin.Context) (code int, message string, data *response.Profile) {
	userId, ok := c.Get("UserId")
	if !ok {
		code = r.USER_NOT_EXISTED
		return
	}

	i, _ := conv.Int(userId)
	data = &response.Profile{
		UserInfo: dao.UserDao.GetUserById(i),
	}

	return
}

func (*User) RefreshToken(c *gin.Context) (code int, message string, data *response.RefreshToken) {
	params, _ := utils.GetRequestParams[request.RefreshToken](c)
	token := params.Token

	var err error
	var refreshToken string
	token, refreshToken, err = utils.RefreshToken(token, params.RefreshToken)
	if err != nil {
		code = r.USER_REFRESHTOKEN_FAILED
		return
	}

	data = &response.RefreshToken{
		Token:        token,
		RefreshToken: refreshToken,
	}

	return
}

func (*User) Login(c *gin.Context) (code int, message string, data *response.Login) {
	params, _ := utils.GetRequestParams[request.Login](c)
	user := getUserByEmail(params.Email)

	if user.ID == 0 {
		// 数据库中不存在用户就进行注册
		code, message, data = signup(params)
	} else if password := getPasswordByUserUUID(user.UUID); password.ID != 0 && utils.CheckPassword(password.Password, params.Password) {
		// 用户密码验证成功
		code = r.OK
		message = "登录成功"
		data = &response.Login{
			UserInfo:  user,
			IsNewUser: false,
		}
	} else {
		// 用户密码验证失败
		code = r.USER_PASSWORD_INCORRECT
	}

	// 如果存在错误，code就不会是r.OK
	if code != r.OK {
		data = nil
		return
	}

	// 登录成功和注册成功都要返回生成的 jwt token
	var err error
	data.Token, data.RefreshToken, err = utils.GenToken(data.UserInfo.ID, data.UserInfo.UUID)
	if err != nil {
		code = r.JWT_AUTHORIZATION_FAILED
		data = nil
		return
	}

	return
}

func signup(params request.Login) (code int, message string, data *response.Login) {
	// 开始事务
	tx := utils.DB.Begin()

	// 创建用户
	user := &entity.User{
		Name:   utils.GenerateRandomUserName(8),
		Avatar: utils.GetCravatarURL(params.Email),
		Email:  params.Email,
	}

	// 数据库报错
	err := tx.Create(&user).Error
	if err != nil {
		tx.Rollback()
		code = r.ERROR_DB_OPE
		data = nil
		return
	}

	// 生成加盐并且哈希处理后的密码
	gpw, err2 := utils.GeneratePassword(params.Password)
	if err2 != nil {
		tx.Rollback()
		code = r.FAIL
		data = nil
		return
	}

	// 创建密码
	password := &entity.UserPassword{
		UserUUID: user.UUID,
		Password: string(gpw),
	}
	err3 := tx.Create(&password).Error
	// 数据库报错
	if err3 != nil {
		tx.Rollback()
		code = r.ERROR_DB_OPE
		data = nil
		return
	}

	// 提交事务
	tx.Commit()

	code = r.OK
	message = "注册成功"
	data = &response.Login{
		UserInfo:  *user,
		IsNewUser: true,
	}

	return
}

// 通过邮箱获取用户
func getUserByEmail(email string) entity.User {
	exist, _ := dao.GetOne[entity.User]("email = ?", email)
	return exist
}

// 通过用户uuid获取密码
func getPasswordByUserUUID(uuid uuid.UUID) entity.UserPassword {
	exist, _ := dao.GetOne[entity.UserPassword]("user_uuid = ?", uuid)
	return exist
}
