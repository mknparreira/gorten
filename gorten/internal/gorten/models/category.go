package models

import "time"

type Category struct {
	CategoryID string    `json:"categoryId" bson:"categoryId"`
	Name       string    `json:"name" binding:"required,min=5,max=10"`
	CreatedAt  time.Time `json:"createdAt" bson:"createdAt"`
}
