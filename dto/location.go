package dto

type LocationRequest struct {
	Name     string `json:"name" binding:"required"`
	Capacity int    `json:"capacity" binding:"required,gte=0"`
}

type GetLocationsRequest struct {
	Sort     string `form:"sort" binding:"omitempty,oneof=ASC DESC"`
	Limit    int    `form:"limit" binding:"omitempty,gte=1"`
	Page     int    `form:"page" binding:"omitempty,gte=1"`
	SortedBy string `form:"sorted_by" binding:"omitempty,oneof=created_at name"`
}

func (r *GetLocationsRequest) DefaultIfEmpty() {
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

type LocationItem struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Capacity  int    `json:"capacity"`
	CreatedAt string `json:"created_at"`
}
