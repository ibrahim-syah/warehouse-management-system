package entity

type (
	Product struct {
		ID         int
		Name       string
		SKU        string
		Quantity   int
		LocationID int
		CreatedAt  string
		UpdatedAt  string
		DeletedAt  *string
	}
)
