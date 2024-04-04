package request

type CreateCategoriesDatum struct {
	Label         *string `json:"label" form:"label"`
	ParentValue   *uint64 `json:"parentValue" form:"parentValue"`
	OrdinalNumber *uint64 `json:"ordinalNumber" form:"ordinalNumber"`
}

type UpdateCategoriesDatum struct {
	CreateCategoriesDatum
	Value *uint64
}

type CreateCategories struct {
	Data []CreateCategoriesDatum `json:"data" form:"data"`
}

type UpdateCategories struct {
	Data []UpdateCategoriesDatum `json:"data" form:"data"`
}

type DeleteCategories struct {
	Value uint64 `json:"value" form:"value"`
}

type DeleteCategoriesBatch struct {
	Value []uint64 `json:"value" form:"value"`
}
