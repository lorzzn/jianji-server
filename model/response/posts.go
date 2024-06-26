package response

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	UserUUID      uuid.UUID `json:"userUuid"`
	UUID          uuid.UUID `json:"uuid"`
	Title         string    `json:"title"`
	Content       string    `json:"content"`
	Description   string    `json:"description"`
	CategoryValue *uint64   `json:"categoryValue"`
	Category      *Category `gorm:"foreignKey:CategoryValue;references:Value;" json:"category"`
	Tags          *[]Tag    `gorm:"many2many:post_tags;foreignKey:UUID;joinForeignKey:PostUUID;references:Value;joinReferences:TagValue" json:"tags"`
	Favoured      bool      `json:"favoured"`
	Archived      *bool     `json:"archived"`
	Public        bool      `json:"public"`
	Status        uint64    `json:"status"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

type ListPost struct {
	Data     *[]Post   `json:"data"`
	PageInfo *PageInfo `json:"pageInfo"`
}
