package entity

import "github.com/jinzhu/gorm"

//User represents user
type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(255);not null; unique"`
	Name     string `gorm:"type:varchar(255);not null"`
	Email    string `gorm:"type:varchar(255);not null; unique" `
	Phone    string `gorm:"type:varchar(100);not null;unique"`
	Password string `gorm:"type:varchar(255);not null"`
}
