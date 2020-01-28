package entity

import "github.com/jinzhu/gorm"

type Status struct {
	gorm.Model
	Name string
}
