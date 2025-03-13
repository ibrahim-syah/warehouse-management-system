package dto

type (
	ProcessOrderRequest struct {
		ProductID int `json:"product_id" binding:"required"`
		Quantity  int `json:"quantity" binding:"required,gte=1"`
	}
)
