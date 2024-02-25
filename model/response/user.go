package response

import "memo-server/entity"

type User struct {
	entity.User
}

type Login struct {
	UserInfo     entity.User
	IsNewUser    bool
	Token        string
	RefreshToken string
}