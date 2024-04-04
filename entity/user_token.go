package entity

import (
	"net"
	"time"

	"github.com/google/uuid"
)

type UserToken struct {
	Universal
	UserFK
	Token             string    `gorm:"type:text" json:"-"`
	TokenUUID         uuid.UUID `gorm:"type:uuid;comment:jwt token的uuid;unique;not null" json:"tokenUUID"`
	ClientFingerprint string    `gorm:"type:text;comment:登录设备浏览器指纹" json:"clientFingerprint"`
	UserAgent         string    `gorm:"type:text;comment:登录设备浏览器user-agent" json:"userAgent"`
	IP                net.IP    `gorm:"comment:登录设备ip地址" json:"ip"`
	Country           string    `gorm:"type:text;comment:地区" json:"country"`
	City              string    `gorm:"type:text;comment:城市" json:"city"`
	Blacklisted       bool      `gorm:"comment:是否拉黑" json:"blacklisted"`
	ExpiresAt         time.Time `gorm:"comment:过期时间" json:"expiresAt"`
	Status            int       `gorm:"type:int;comment:token状态;not null;default:1" json:"status"`
}
