package request

type CreateTagsDatum struct {
	Label *string `json:"label" form:"label"`
}
type UpdateTagDatum struct {
	CreateTagsDatum
	Value *uint64 `json:"value" form:"value"`
}

type CreateTags struct {
	Data []CreateTagsDatum `json:"data" form:"data"`
}

type UpdateTags struct {
	Data []UpdateTagDatum `json:"data" form:"data"`
}

type DeleteTag struct {
	Value uint64 `json:"value" form:"value"`
}

type DeleteTags struct {
	Value []uint64 `json:"value" form:"value"`
}
