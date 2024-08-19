package models

type User struct {
	UserID   string `json:"userId" bson:"userId"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
