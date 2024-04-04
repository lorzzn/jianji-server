package common

import (
	"jianji-server/entity"

	"github.com/google/uuid"
)

type UserFK struct {
	UserUUID uuid.UUID   `gorm:"type:uuid;comment:对应user表中uuid;not null" json:"userUUID"`
	User     entity.User `gorm:"foreignKey:UserUUID;references:UUID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
}

type UniqueUserFK struct {
	UserUUID uuid.UUID   `gorm:"type:uuid;comment:对应user表中uuid;unique;not null" json:"userUUID"`
	User     entity.User `gorm:"foreignKey:UserUUID;references:UUID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
}
