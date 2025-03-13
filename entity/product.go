package entity

type (
	Product struct {
		ID          int
		Name        string
		CategoryID  int
		Description string
		CreatedAt   string
		UpdatedAt   string
		DeletedAt   *string
	}
)
