package entity

import "jianji-server/entity/common"

type Tag struct {
	common.Universal
	common.UserFK
	Label string `gorm:"type:varchar(32);comment:名称" json:"label"`
	Value uint64 `gorm:"auto_increment;unique;not null;autoIncrement:100;comment:值" json:"value"`
}
