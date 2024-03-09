package entity

import "github.com/google/uuid"

type UserPassword struct {
	Universal
	UserUUID uuid.UUID `gorm:"type:uuid;comment:对应user表中uuid;unique;not null" json:"user_uuid"`
	Password string    `gorm:"type:varchar(256);comment:密码" json:"password"`
}
