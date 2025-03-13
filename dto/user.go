package dto

type GetUsersRequest struct {
	PaginatedRequest
	SortedBy *string `form:"sorted_by" binding:"omitempty,oneof=created_at email"`
}

type UserItem struct {
	ID        int
	Email     string
	Role      string
	CreatedAt string
}
