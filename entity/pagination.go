package entity

type PaginationParam struct {
	OrderBy        string
	OrderDirection string
	Limit          int
	Offset         int
	TotalRecords   int
}
