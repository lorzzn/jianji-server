package entity

import (
	"jianji-server/entity/common"
)

type UserPassword struct {
	common.Universal
	common.UniqueUserFK
	Password string `gorm:"type:varchar(256);comment:密码" json:"-"`
}
