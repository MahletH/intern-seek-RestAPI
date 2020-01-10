package service

import (
	"github.com/Beemnet/internseek/entity"
	"github.com/Beemnet/internseek/internship"
)

// InternshipService implements menu.CommentService interface
type InternshipService struct {
	internshipRepo internship.InternshipRepository
}

// NewInternshipService returns a new CommentService object
func NewInternshipService(internRepo internship.InternshipRepository) internship.InternshipService {
	return &InternshipService{internshipRepo: internRepo}
}

// Internships returns all stored internships
func (is *InternshipService) Internships() ([]entity.Internship, []error) {

	intern, errs := is.internshipRepo.Internships()
	if len(errs) > 0 {
		return nil, errs
	}
	return intern, errs
}

// Internship retrievs a given user internship by its id
func (is *InternshipService) Internship(id uint) (*entity.Internship, []error) {

	intern, errs := is.internshipRepo.Internship(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return intern, errs

}

// UpdateInternship updates a given user internship
func (is *InternshipService) UpdateInternship(internship *entity.Internship) (*entity.Internship, []error) {
	intern, errs := is.internshipRepo.UpdateInternship(internship)
	if len(errs) > 0 {
		return nil, errs
	}
	return intern, errs
}

// DeleteInternship deletes a given user internship
func (is *InternshipService) DeleteInternship(id uint) (*entity.Internship, []error) {

	intern, errs := is.internshipRepo.DeleteInternship(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return intern, errs
}

// StoreInternship stores a given user internship
func (is *InternshipService) StoreInternship(internship *entity.Internship) (*entity.Internship, []error) {
	intern, errs := is.internshipRepo.StoreInternship(internship)
	if len(errs) > 0 {
		return nil, errs
	}
	return intern, errs
}
