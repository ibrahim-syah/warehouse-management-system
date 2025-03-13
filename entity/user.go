package entity

type User struct {
	ID        int
	Email     string
	Password  string
	Role      string
	CreatedAt string
	UpdatedAt string
	DeletedAt *string
}

type InsertUser struct {
	Email    string
	Password string
	Role     string
}
