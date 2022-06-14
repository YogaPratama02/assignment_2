package models

import "time"

type Item struct {
	Id          int       `json:"id" gorm:"primaryKey"`
	ItemCode    string    `json:"item_code"`
	Description string    `json:"description"`
	Quantity    int       `json:"quantity"`
	OrderId     int       `json:"order_id"`
	Order       *Order    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"order"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}
