package model

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Age      int    `json:"age,omitempty"`
	Tell     string `json:"tell"`
	Gender   string `json:"gender"`
}
