package request

type Encrypted struct {
	Key  string `json:"key" form:"key"`
	Data string `json:"data" form:"data"`
}
