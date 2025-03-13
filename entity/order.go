package entity

type (
	Order struct {
		ID        int
		ProductID int
		Quantity  int
		Type      string
		CreatedAt string
		UpdatedAt string
		DeletedAt *string
	}
)
