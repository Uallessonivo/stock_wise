package models

import "time"

type Product struct {
	ID                         string `gorm:"primary_key"`
	Name                       string
	Category                   string
	ExpiryDate                 time.Time
	Lot                        string
	Supplier                   Supplier
	Price                      float64
	StockQuantity              int
	MinimumAcceptableQuantity  int
	MaximumRecommendedQuantity int
}
