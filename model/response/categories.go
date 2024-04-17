package response

import "time"

type Category struct {
	Label         string  `json:"label"`
	Value         uint64  `json:"value"`
	ParentValue   *uint64 `json:"parentValue"`
	OrdinalNumber *uint64 `json:"ordinalNumber"`
}

type CategoryStatistics struct {
	TotalPosts int64     `json:"totalPosts"`
	CreateAt   time.Time `json:"createAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
