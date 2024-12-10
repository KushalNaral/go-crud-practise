package models

type User struct {
	Id    int    `json:"id,omitempty"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Account struct {
	AccountNumber string  `json:"account_number"`
	UserId        int     `json:"user_id"`
	Name          string  `json:"name"`
	Balance       float32 `json:"balance"`
}
