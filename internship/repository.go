package internship

import "github.com/MahletH/intern-seek-RestAPI/entity"

// InternshipRepository specifies Internship related service
type InternshipRepository interface {
	Internships() ([]entity.Internship, []error)
	CompanyInternships(company *entity.CompanyDetail) ([]entity.Internship, []error)
	Internship(id uint) (*entity.Internship, []error)
	UpdateInternship(internship *entity.Internship) (*entity.Internship, []error)
	DeleteInternship(id uint) (*entity.Internship, []error)
	StoreInternship(internship *entity.Internship) (*entity.Internship, []error)
}

/*
// FieldsRepository specifies field Internship related service
type FieldsRepository interface {
	Fields() ([]entity.Fields, []error)
	Field(id uint) (*entity.Fields, []error)
	InternshipFields(interID uint) ([]entity.Fields, error)
	UpdateFields(field *entity.Fields) (*entity.Fields, []error)
	DeleteField(id uint) (*entity.Fields, []error)
}
*/
