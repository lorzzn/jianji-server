package entity

import "github.com/google/uuid"

type PostTags struct {
	Universal
	PostUUID uuid.UUID `gorm:"primary_key"`
	TagValue uint64    `gorm:"primary_key"`
}
