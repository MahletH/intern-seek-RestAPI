package service

import (
	"github.com/nebyubeyene/Intern-Seek-Version-1/entity"
	"github.com/nebyubeyene/Intern-Seek-Version-1/user"
)

// CompanyService implements user.CompanyService interface
type CompanyService struct {
	companyRepo user.CompanyRepository
}

// NewCompanyService  returns a new CompanyService object
func NewCompanyService(companyRepository user.CompanyRepository) *CompanyService {
	return &CompanyService{companyRepo: companyRepository}
}

// Companies returns all stored application company_details
func (cs *CompanyService) Companies() ([]entity.CompanyDetail, []error) {
	comps, errs := cs.companyRepo.Companies()
	if len(errs) > 0 {
		return nil, errs
	}
	return comps, errs
}

// Company retrieves an application company_detail by its id
func (cs *CompanyService) Company(id uint) (*entity.CompanyDetail, []error) {
	comp, errs := cs.companyRepo.Company(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return comp, errs
}

// UpdateCompany updates  a given company_detail
func (cs *CompanyService) UpdateCompany(company *entity.CompanyDetail) (*entity.CompanyDetail, []error) {
	comp, errs := cs.companyRepo.UpdateCompany(company)
	if len(errs) > 0 {
		return nil, errs
	}
	return comp, errs
}

// DeleteCompany deletes a given application company_detail
func (cs *CompanyService) DeleteCompany(id uint) (*entity.CompanyDetail, []error) {
	comp, errs := cs.companyRepo.DeleteCompany(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return comp, errs
}

// StoreCompany stores a given application company_detail
func (cs *CompanyService) StoreCompany(company *entity.CompanyDetail) (*entity.CompanyDetail, []error) {
	comp, errs := cs.companyRepo.StoreCompany(company)
	if len(errs) > 0 {
		return nil, errs
	}
	return comp, errs
}
