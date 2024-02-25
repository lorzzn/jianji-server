package entity

type User struct {
	Universal
	Name   string `gorm:"type:varchar(16);comment:名称" json:"name"`
	Avatar string `gorm:"type:varchar(128);comment:头像" json:"avatar"`
	Email  string `gorm:"type:varchar(64);comment:邮箱;unique;not null" json:"email"`
}
