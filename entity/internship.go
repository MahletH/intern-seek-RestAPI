package entity

import "time"

// Internship is a struct with all properties of an internship
type Internship struct {
	ID                    uint
	CompanyID             uint      `json:"company_id"`
	Name                  string    `json:"name" gorm:"type:varchar(255)"`
	RequiredAcademicLevel string    `json:"required_academic_level" gorm:"type:varchar(255)"`
	Description           string    `json:"description"`
	ClosingDate           time.Time `json:"closing_date"`
	FieldsReq             []Fields  `gorm:"one2many:fields"`

	//numOfInterns     int
	//salary           float64

}

//Fields is a struct with all required fields under every internship
type Fields struct {
	InternshipID uint   `json:"internship_id"`
	Field        string `json:"field" gorm:"type:varchar(255)"`
}

// Error represents error message
type Error struct {
	Code    int
	Message string
}
