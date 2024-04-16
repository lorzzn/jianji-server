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
	"github.com/samber/lo"
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
	_ = mapstructure.Decode(updated, &e)
	i, _ := conv.Uint64(userId)
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

	i, _ := conv.Uint64(userId)
	data = &response.Profile{
		UserInfo: dao.UserDao.GetUserById(i),
	}

	return
}

func (*User) GetStatistics(c *gin.Context) (code int, message string, data *response.UserStatistics) {
	userUUID := c.MustGet("UserUUID").(uuid.UUID)

	var totalPosts int64
	var totalWords int64

	query := utils.DBQueryBegin()
	baseQuery := query.Model(&entity.Post{}).Where(&entity.Post{
		UserFK: entity.UserFK{UserUUID: userUUID},
	})

	err := baseQuery.Count(&totalPosts).Error
	if err == nil {
		err = baseQuery.Select("SUM(LENGTH(content))").Row().Scan(&totalWords)
	}
	if err != nil {
		query.Rollback()
		data = nil
		code = r.ERROR_DB_OPE
		return
	}

	data = &response.UserStatistics{
		TotalPosts: totalPosts,
		TotalWords: totalWords,
	}

	query.Commit()

	return
}

func (*User) RefreshToken(c *gin.Context) (code int, message string, data *response.RefreshToken) {
	params, _ := utils.GetRequestParams[request.RefreshToken](c)
	tokenData, err := utils.RefreshToken(c, params.Token, params.RefreshToken)
	if err != nil {
		code = r.USER_REFRESHTOKEN_FAILED
		return
	}

	data = &response.RefreshToken{
		Token:        tokenData.Token,
		RefreshToken: tokenData.RefreshToken,
	}

	return
}

func (*User) Logout(c *gin.Context) (code int, message string) {
	tokenUUID, ok := c.Get("TokenUUID")
	if !ok {
		code = r.USER_NOT_LOGIN
		return
	}

	utils.AddTokenToBlacklist(tokenUUID.(uuid.UUID))

	code = r.OK
	message = " 退出登录成功"

	return
}

func (*User) Login(c *gin.Context) (code int, message string, data *response.Login) {
	params, _ := utils.GetRequestParams[request.Login](c)
	user := getUserByEmail(params.Email)

	if user.ID == 0 {
		// 数据库中不存在用户，用户未注册
		code = r.USER_NOT_EXISTED
		return
	} else if password := getPasswordByUserUUID(user.UUID); password.ID != 0 && lo.Ternary(
		params.From == "active", // 如果是激活账户操作，直接字符串比较，否则就是登录操作，需要对比hash后的密码
		password.Password == params.Password,
		utils.CheckPassword(password.Password, params.Password),
	) {
		// 用户密码验证成功
		code = r.OK
		message = "登录成功"
		data = &response.Login{
			UserInfo: user,
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
	tokenDate, err := utils.GenToken(c, data.UserInfo.ID, data.UserInfo.UUID, params.Fingerprint)
	if err != nil {
		code = r.TOKEN_AUTHORIZATION_FAILED
		data = nil
		return
	}

	data.Token = tokenDate.Token
	data.RefreshToken = tokenDate.RefreshToken

	return
}

func (*User) Active(c *gin.Context) (code int, message string, data *response.Login) {
	params, _ := utils.GetRequestParams[request.Active](c)
	email, password, fingerprint, err := utils.GetActiveEmailStateInfo(params.Email, params.State)
	if err != nil {
		code = 500
		message = err.Error()
		return
	}

	// 开始事务
	tx := utils.DBQueryBegin()

	// 创建用户
	user := &entity.User{
		Name:   utils.GenerateRandomUserName(8),
		Avatar: utils.GetCravatarURL(email),
		Email:  email,
	}

	// 数据库报错
	err = tx.Create(&user).Error
	if err != nil {
		tx.Rollback()
		code = r.ERROR_DB_OPE
		data = nil
		return
	}

	// 创建密码记录
	pw := &entity.UserPassword{
		UniqueUserFK: entity.UniqueUserFK{UserUUID: user.UUID},
		Password:     password,
	}
	err3 := tx.Create(&pw).Error
	// 数据库报错
	if err3 != nil {
		tx.Rollback()
		code = r.ERROR_DB_OPE
		data = nil
		return
	}

	// 提交事务
	tx.Commit()

	//将登录操作的参数保存到上下文，后面交给下一个中间件Login
	c.Set(utils.ContextRequestParams, request.Login{
		Email:       email,
		Password:    password,
		Fingerprint: fingerprint,
		From:        "active",
	})

	return
}

func (*User) Signup(c *gin.Context) (code int, message string, data *response.Login) {
	params, _ := utils.GetRequestParams[request.Signup](c)
	if user := getUserByEmail(params.Email); user.ID != 0 {
		code = r.USER_EXISTED
		return
	}

	// 生成加盐并且哈希处理后的密码
	gpw, err := utils.GeneratePassword(params.Password)
	if err != nil {
		code = r.FAIL
		data = nil
		return
	}

	err = utils.SendActiveEmail(params.Email, string(gpw), params.Fingerprint)
	if err != nil {
		code = 500
		message = "发送激活邮件失败，请检查您的邮箱地址或联系我们的技术支持团队。"
		return
	}

	data = &response.Login{
		UserInfo: entity.User{Status: 0},
	}
	message = "need signup"

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
