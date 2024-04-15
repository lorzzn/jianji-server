package request

import (
	"github.com/google/uuid"
)

type PostTags struct {
	TagValues *[]uint64 `json:"tagValues"`
}

type PostCommon struct {
	Title         string  `json:"title"`
	Content       string  `json:"content"`
	CategoryValue *uint64 `json:"categoryValue"`
	Favoured      *bool   `json:"favoured"`
	Public        *bool   `json:"public"`
	Status        *uint64 `json:"status"`
	Archived      *bool   `json:"archived"`
}

type CreatePost struct {
	PostCommon
	PostTags
}

type UpdatePost struct {
	CreatePost
	UUID uuid.UUID `json:"uuid"`
}

type ListPost struct {
	PageInfo
}

type GetPost struct {
	UUID uuid.UUID `json:"uuid"`
}

type DeletePost struct {
	GetPost
}
