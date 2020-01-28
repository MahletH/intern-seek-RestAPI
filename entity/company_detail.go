package entity

import "github.com/jinzhu/gorm"

type CompanyDetail struct {
	gorm.Model
	UserID      uint   `sql:"type:int REFERENCES users(ID)"`
	Country     string `gorm:"type:varchar(255);not null"`
	City        string `gorm:"type:varchar(255);not null"`
	FocusArea   string `gorm:"type:varchar(255);not null"`
	Description string `gorm:"type:varchar(255)"`
}
