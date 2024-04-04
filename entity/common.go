package entity

import (
	"time"

	"github.com/google/uuid"
)

type Universal struct {
	ID        uint64    `gorm:"primary_key;auto_increment;unique;not null" json:"id" mapstructure:"id"`
	UUID      uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();unique;not null" json:"uuid"`
	CreatedAt time.Time `json:"createdAt" mapstructure:"-"`
	UpdatedAt time.Time `json:"updatedAt" mapstructure:"-"`
}

type UserFK struct {
	UserUUID uuid.UUID `gorm:"type:uuid;comment:对应user表中uuid;not null" json:"userUUID"`
	User     User      `gorm:"foreignKey:UserUUID;references:UUID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
}

type UniqueUserFK struct {
	UserUUID uuid.UUID `gorm:"type:uuid;comment:对应user表中uuid;unique;not null" json:"userUUID"`
	User     User      `gorm:"foreignKey:UserUUID;references:UUID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
}
