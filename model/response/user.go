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

type UserStatistics struct {
	TotalPosts int64 `json:"totalPosts"`
	TotalWords int64 `json:"totalWords"`
}
