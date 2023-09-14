package models

import "time"

type StockMovement struct {
	ID        string `gorm:"primary_key"`
	CreatedAt time.Time
	Type      string
	Quantity  int
	Reason    string
}
