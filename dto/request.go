package dto

type PaginatedRequest struct {
	// SortedBy         *string `form:"sorted_by" binding:"omitempty,oneof=<determined by derivative struct>"`
	// Search         *string `form:"search" binding:"omitempty"`
	Sort   *string `form:"sort" binding:"omitempty,oneof=asc desc"`
	Limit  *int    `form:"limit" binding:"omitempty,gt=0"`
	Page   *int    `form:"page" binding:"omitempty,gt=0"`
	Offset int     `form:"-"`
}
