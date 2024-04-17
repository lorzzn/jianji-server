package response

import "time"

type Tag struct {
	Label string `json:"label"`
	Value uint64 `json:"value"`
}

type TagStatistics struct {
	TotalPosts int64     `json:"totalPosts"`
	CreateAt   time.Time `json:"createAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
