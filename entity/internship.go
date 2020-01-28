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
	FieldsReq             []Field   `gorm:"many2many:internship_req_fields"`

	//numOfInterns     int
	//salary           float64

}

// Error represents error message
type Error struct {
	Code    int
	Message string
}
