package models

type Supplier struct {
	ID      string `gorm:"primary_key"`
	Name    string
	Contact string
}
