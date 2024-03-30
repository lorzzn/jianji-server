package entity

import (
	"time"

	"github.com/google/uuid"
)

type UserToken struct {
	Universal
	UserUUID          uuid.UUID `gorm:"type:uuid;comment:对应user表中uuid;not null" json:"userUUID"`
	User              User      `gorm:"foreignKey:UserUUID;references:UUID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
	Token             string    `gorm:"type:text" json:"-"`
	TokenUUID         uuid.UUID `gorm:"type:uuid;comment:jwt token的uuid;unique;not null" json:"tokenUUID"`
	ClientFingerprint string    `gorm:"type:text;comment:登录设备浏览器指纹" json:"clientFingerprint"`
	UserAgent         string    `gorm:"type:text;comment:登录设备浏览器user-agent" json:"userAgent"`
	Country           string    `gorm:"type:text;comment:地区" json:"country"`
	City              string    `gorm:"type:text;comment:城市" json:"city"`
	Blacklisted       bool      `gorm:"comment:是否拉黑" json:"blacklisted"`
	ExpiresAt         time.Time `gorm:"comment:过期时间" json:"expiresAt"`
	Status            int       `gorm:"type:int;comment:token状态;not null;default:1" json:"status"`
}
