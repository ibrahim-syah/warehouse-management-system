package dto

type (
	ProcessOrderRequest struct {
		ProductID int `json:"product_id" binding:"required"`
		Quantity  int `json:"quantity" binding:"required,gte=1"`
	}

	ProcessOrderResponse struct {
		OrderID int `json:"order_id"`
	}

	GetOrdersRequest struct {
		Sort     string `form:"sort" binding:"omitempty,oneof=ASC DESC"`
		Limit    int    `form:"limit" binding:"omitempty,gte=1"`
		Page     int    `form:"page" binding:"omitempty,gte=1"`
		SortedBy string `form:"sorted_by" binding:"omitempty,oneof=created_at product_id quantity type"`
	}

	OrderItem struct {
		ID        int    `json:"id"`
		ProductID int    `json:"product_id"`
		Quantity  int    `json:"quantity"`
		Type      string `json:"type"`
		CreatedAt string `json:"created_at"`
	}
)

func (r *GetOrdersRequest) DefaultIfEmpty() {
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
