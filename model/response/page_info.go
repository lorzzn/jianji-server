package response

type PageInfo struct {
	PageNo     *int   `json:"pageNo"`
	PageSize   *int   `json:"pageSize"`
	TotalCount *int64 `json:"totalCount"`
}
