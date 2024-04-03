package request

type CreateCategories struct {
	Label         *string `json:"label" form:"label"`
	ParentValue   *uint64 `json:"parentValue" form:"parentValue"`
	OrdinalNumber *uint64 `json:"ordinalNumber" form:"ordinalNumber"`
}

type UpdateCategoriesDatum struct {
	CreateCategories
	Value *uint64
}

type UpdateCategories struct {
	Data []UpdateCategoriesDatum `json:"data" form:"data"`
}

type DeleteCategories struct {
	Value uint64 `json:"value" form:"value"`
}
