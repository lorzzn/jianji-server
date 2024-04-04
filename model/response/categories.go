package response

type Categories struct {
	Label         string  `json:"label"`
	Value         uint64  `json:"value"`
	ParentValue   *uint64 `json:"parentValue"`
	OrdinalNumber *uint64 `json:"ordinalNumber"`
}
