package models

import "time"

type Order struct {
	Id           int       `json:"id" gorm:"primaryKey"`
	CustomerName string    `json:"customer_name"`
	OrderedAt    time.Time `json:"ordered_at"`
	Item         []Item    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"items"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
}
