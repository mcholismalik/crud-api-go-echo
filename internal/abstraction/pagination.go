package abstraction

type Pagination struct {
	Page     *int    `query:"page" json:"page"`
	PageSize *int    `query:"page_size" json:"page_size"`
	SortBy   *string `query:"sort_by" json:"sort_by"`
	Sort     *string `query:"sort" json:"sort"`
}

type PaginationInfo struct {
	*Pagination
	Count       int  `json:"count"`
	MoreRecords bool `json:"more_records"`
}
