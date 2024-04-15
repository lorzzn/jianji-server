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
	Description   string  `json:"description"`
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
	Archived *bool   `json:"archived"`
	Favoured *bool   `json:"favoured"`
	SortBy   *string `json:"sortBy"`
	SortType *string `json:"sortType"`
}

type GetPost struct {
	UUID uuid.UUID `json:"uuid"`
}

type DeletePost struct {
	GetPost
}
