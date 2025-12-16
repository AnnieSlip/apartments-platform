package user

type User struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

type Users []User
