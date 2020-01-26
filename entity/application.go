package entity

import "github.com/jinzhu/gorm"

type Application struct {
	gorm.Model
	internshipID    uint   `sql:"type:int REFERENCES internship(ID)"`
	applicantID     uint   `sql:"type:int REFERENCES users(ID)"`
	applicationDate string `gorm:"type:timestamp;not null"`
	status          string `gorm:"type:varchar(30);not null"`
}
