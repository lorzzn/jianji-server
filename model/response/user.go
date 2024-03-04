package response

import "jianji-server/entity"

type Profile struct {
	UserInfo entity.User
}

type Login struct {
	UserInfo     entity.User
	Token        string
	RefreshToken string
}

type RefreshToken struct {
	Token        string
	RefreshToken string
}
