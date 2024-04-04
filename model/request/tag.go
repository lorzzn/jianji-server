package request

type CreateTagDatum struct {
	Label *string `json:"label" form:"label"`
}
type UpdateTagDatum struct {
	CreateTagDatum
	Value *uint64 `json:"value" form:"value"`
}

type CreateTag struct {
	Data []CreateTagDatum `json:"data" form:"data"`
}

type UpdateTag struct {
	Data []UpdateTagDatum `json:"data" form:"data"`
}

type DeleteTags struct {
	Value uint64 `json:"value" form:"value"`
}

type DeleteTagBatch struct {
	Value []uint64 `json:"value" form:"value"`
}
