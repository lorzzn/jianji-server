package entity

import "github.com/google/uuid"

type UserPassword struct {
	Universal
	UserUUID uuid.UUID `gorm:"type:uuid;comment:对应user表中uuid;unique;not null" json:"userUUID"`
	User     User      `gorm:"foreignKey:UserUUID;references:UUID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
	Password string    `gorm:"type:varchar(256);comment:密码" json:"-"`
}
