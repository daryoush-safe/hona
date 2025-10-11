package entities

import "gorm.io/gorm"

type Pet struct {
	gorm.Model
	Name string
	Age  uint8
}
