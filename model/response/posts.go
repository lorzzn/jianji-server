package response

import (
	"github.com/google/uuid"
)

type Post struct {
	UUID     uuid.UUID `json:"uuid"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	Category *uint64   `json:"category"`
	Tags     *[]Tag    `gorm:"many2many:post_tags;foreignKey:UUID;joinForeignKey:PostUUID;references:Value;joinReferences:TagValue" json:"tags"`
	Favoured *bool     `json:"favoured"`
	Public   *bool     `json:"public"`
	Status   *uint64   `json:"status"`
}
