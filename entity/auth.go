package entity

type EmailPassword struct {
	Email    string
	Password string
}

type LoginToken struct {
	UserID      int
	AccessToken string
}
