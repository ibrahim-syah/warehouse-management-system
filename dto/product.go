package dto

type ProductRequest struct {
	Name       string `json:"name" binding:"required"`
	SKU        string `json:"sku" binding:"required"`
	Quantity   int    `json:"quantity" binding:"required,gte=0"`
	LocationID int    `json:"location_id" binding:"required"`
}

type GetProductsRequest struct {
	Sort     string `form:"sort" binding:"omitempty,oneof=ASC DESC"`
	Limit    int    `form:"limit" binding:"omitempty,gte=1"`
	Page     int    `form:"page" binding:"omitempty,gte=1"`
	SortedBy string `form:"sorted_by" binding:"omitempty,oneof=created_at name sku quantity location_id"`
}

func (r *GetProductsRequest) DefaultIfEmpty() {
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

type ProductItem struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	SKU        string `json:"sku"`
	Quantity   int    `json:"quantity"`
	LocationID int    `json:"location_id"`
	CreatedAt  string `json:"created_at"`
}
