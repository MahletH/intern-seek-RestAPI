
package entity

import "github.com/jinzhu/gorm"

type PersonalDetails struct {
	gorm.Model
	UserID      uint   `sql:"type:int REFERENCES users(ID)"`
	Fields   []Field   `gorm:"many2many:field_internship"`
	AcademicLevel string `gorm:"type:varchar(255);not null"`
}
