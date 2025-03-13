package dto

type GetUsersRequest struct {
	Sort     string `form:"sort" binding:"omitempty,oneof=ASC DESC"`
	Limit    int    `form:"limit" binding:"omitempty,gte=1"`
	Page     int    `form:"page" binding:"omitempty,gte=1"`
	SortedBy string `form:"sorted_by" binding:"omitempty,oneof=created_at email"`
}

func (r *GetUsersRequest) DefaultIfEmpty() {
	if r.Limit == 0 {
		r.Limit = 10
	}
	if r.Sort == "" {
		r.Sort = "DESC"
	}
	if r.Page == 0 {
		r.Page = 1
	}
	if r.SortedBy == "" {
		r.SortedBy = "created_at"
	}
}

type UserItem struct {
	ID        int
	Email     string
	Role      string
	CreatedAt string
}
