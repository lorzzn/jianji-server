package response

import "jianji-server/entity"

type Profile struct {
	UserInfo entity.User
}

type Login struct {
	UserInfo     entity.User
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}

type RefreshToken struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}
