package models

import "time"

type Company struct {
	CompanyID string    `json:"companyId" bson:"companyId"`
	Name      string    `json:"name" binding:"required,min=10,max=20"`
	Address   string    `json:"address"`
	Contact   string    `json:"contact"`
	Email     string    `json:"email" binding:"required,email"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
}
