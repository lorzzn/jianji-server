package entity

import "github.com/google/uuid"

type PostTags struct {
	Universal
	PostUUID uuid.UUID `gorm:"primary_key"`
	TagValue uint64    `gorm:"primary_key"`
}

func (t PostTags) TableName() string {
	return "post_tags"
}
