package internship

import "github.com/abdimussa87/intern-seek-RestAPI/entity"

// InternshipService specifies customer comment related service
type InternshipService interface {
	Internships() ([]entity.Internship, []error)
	Internship(id uint) (*entity.Internship, []error)
	CompanyInternships(company *entity.CompanyDetail) ([]entity.Internship, []error)
	UpdateInternship(internship *entity.Internship) (*entity.Internship, []error)
	DeleteInternship(id uint) (*entity.Internship, []error)
	StoreInternship(internship *entity.Internship) (*entity.Internship, []error)
}
