package entity

import (
	"time"

	"github.com/google/uuid"
)

type Universal struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id" mapstructure:"id"`
	UUID      uuid.UUID `json:"uuid"`
	CreatedAt time.Time `json:"created_at" mapstructure:"-"`
	UpdatedAt time.Time `json:"updated_at" mapstructure:"-"`
}
