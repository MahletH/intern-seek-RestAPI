package service

import (
	"github.com/lensabillion/intern-seek-RestAPI/entity"
	"github.com/lensabillion/intern-seek-RestAPI/user"

)

// InternService implements user.InternService interface
type InternService struct {
	InternRepo user.InternRepository
}

// NewInternService  returns a new InternService object
func NewInternService(internRepository user.InternRepository) *InternService {
	return &InternService{InternRepo: internRepository}
}

// Interns returns all stored application PersonalDetsils
func (ins *InternService) Interns() ([]entity.PersonalDetails, []error) {
	ints, errs := ins.InternRepo.Interns()
	if len(errs) > 0 {
		return nil, errs
	}
	return ints, errs
}

// Intern retrieves an application PersonalDetail by its id
func (ins *InternService) Intern(id uint) (*entity.PersonalDetails, []error) {
	int, errs := ins.InternRepo.Intern(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return int, errs
}

// GetInternByUserId retrieves an application PersonalDetails by its id
func (ins *InternService) GetInternByUserId(id uint) (*entity.PersonalDetails, []error) {
	int, errs := ins.InternRepo.GetInternByUserId(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return int, errs
}

// UpdateIntern updates  a given PersonalDetails
func (ins *InternService) UpdateIntern(intern *entity.PersonalDetails) (*entity.PersonalDetails, []error) {
	int, errs :=ins.InternRepo.UpdateIntern(intern)
	if len(errs) > 0 {
		return nil, errs
	}
	return int, errs
}

// DeleteIntern deletes a given application intern_detail
func (ins *InternService) DeleteIntern(id uint) (*entity.PersonalDetails, []error) {
	int, errs := ins.InternRepo.DeleteIntern(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return int, errs
}


// StoreIntern stores a given application company_detail
func (ins *InternService) StoreIntern(intern *entity.PersonalDetails) (*entity.PersonalDetails, []error) {
	int, errs :=ins.InternRepo.StoreIntern(intern)
	if len(errs) > 0 {
		return nil, errs
	}
	return int, errs
}

