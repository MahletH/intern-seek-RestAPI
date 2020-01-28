package entity

import (
	"github.com/jinzhu/gorm"
)

type Field struct {
	gorm.Model

	Name        string       `gorm:"type:varchar(55);not null"`
	Intern      []Intern     `gorm:"many2many:field_internship"`
	Internships []Internship `gorm:"many2many:internship_req_fields"`
}
