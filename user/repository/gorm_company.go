package repository

import (
	"github.com/abdimussa87/Intern-Seek-Version-1/entity"
	"github.com/abdimussa87/Intern-Seek-Version-1/user"
	"github.com/jinzhu/gorm"
)

// CompanyGormRepo Implements the user.CompanyRepository interface
type CompanyGormRepo struct {
	conn *gorm.DB
}

// NewCompanyGormRepoImpl creates a new object of CompanyGormRepo
func NewCompanyGormRepoImpl(db *gorm.DB) user.CompanyRepository {
	return &CompanyGormRepo{conn: db}
}

// Companies return all company_details from the database
func (compRepo *CompanyGormRepo) Companies() ([]entity.CompanyDetail, []error) {
	companies := []entity.CompanyDetail{}
	errs := compRepo.conn.Find(&companies).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return companies, errs
}

// Company retrieves a company_detail by its id from the database
func (compRepo *CompanyGormRepo) Company(id uint) (*entity.CompanyDetail, []error) {
	company := entity.CompanyDetail{}
	errs := compRepo.conn.First(&company, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &company, errs
}

// UpdateCompany updates a given company_detail in the database
func (compRepo *CompanyGormRepo) UpdateCompany(company *entity.CompanyDetail) (*entity.CompanyDetail, []error) {
	comp := company
	errs := compRepo.conn.Save(comp).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return comp, errs
}

// DeleteCompany deletes a given company_detail from the database
func (compRepo *CompanyGormRepo) DeleteCompany(id uint) (*entity.CompanyDetail, []error) {
	comp, errs := compRepo.Company(id)
	if len(errs) > 0 {
		return nil, errs
	}
	errs = compRepo.conn.Delete(comp, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return comp, errs
}

// StoreCompany stores a new Company_detail into the database
func (compRepo *CompanyGormRepo) StoreCompany(company *entity.CompanyDetail) (*entity.CompanyDetail, []error) {
	comp := company
	errs := compRepo.conn.Create(comp).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return comp, errs
}
