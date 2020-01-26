package entity

import "github.com/jinzhu/gorm"

type PersonalDetails struct {
	gorm.Model
	UserID      uint   `sql:"type:int REFERENCES users(ID)"`
	Field   string `gorm:"type:varchar(255);not null"`
	AcademicLevel string `gorm:"type:varchar(255);not null"`
}
