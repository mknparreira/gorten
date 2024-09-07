package models

import "time"

type User struct {
	UserID    string    `json:"userId" bson:"userId"`
	Name      string    `json:"name" binding:"required,min=10,max=50"`
	Email     string    `json:"email" binding:"required,email"`
	Password  string    `json:"password" binding:"required,min=6"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
}
