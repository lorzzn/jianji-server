package entity

import (
	"time"

	"github.com/google/uuid"
)

type Universal struct {
	ID        uint64    `gorm:"primary_key;auto_increment;unique;not null" json:"id" mapstructure:"id"`
	UUID      uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();unique;not null" json:"uuid"`
	CreatedAt time.Time `json:"created_at" mapstructure:"-"`
	UpdatedAt time.Time `json:"updated_at" mapstructure:"-"`
}
