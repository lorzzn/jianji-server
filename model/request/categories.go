package request

type CreateCategories struct {
	Label       string `json:"label" form:"label"`
	ParentValue uint64 `json:"parentValue" form:"parentValue"`
}

type DeleteCategories struct {
	Value uint64 `json:"value" form:"value"`
}
