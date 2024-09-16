package models

import "time"

type Product struct {
	ProductID   string    `json:"productId" bson:"productId"`
	Name        string    `json:"name" binding:"required,min=3,max=10"`
	Description string    `json:"description"`
	Price       float64   `json:"price" binding:"required" validate:"gt=0"`
	CategoryID  string    `json:"categoryId" bson:"categoryId" binding:"required"`
	CompanyID   string    `json:"companyId" bson:"companyId" binding:"required"`
	CreatedAt   time.Time `json:"createdAt" bson:"createdAt"`
}
