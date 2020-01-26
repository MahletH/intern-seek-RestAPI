package internship

import "github.com/MahletH/intern-seek-RestAPI/entity"

// InternshipService specifies customer comment related service
type InternshipService interface {
	Internships() ([]entity.Internship, []error)
	Internship(id uint) (*entity.Internship, []error)
	UpdateInternship(internship *entity.Internship) (*entity.Internship, []error)
	DeleteInternship(id uint) (*entity.Internship, []error)
	StoreInternship(internship *entity.Internship) (*entity.Internship, []error)
}
