package repository

import (
	"github.com/abdimussa87/intern-seek-RestAPI/entity"
	"github.com/abdimussa87/intern-seek-RestAPI/internship"

	"github.com/jinzhu/gorm"
)

// InternshipGormRepo implements menu.InternshipRepository interface
type InternshipGormRepo struct {
	conn *gorm.DB
}

// NewInternshipGormRepo returns new object of InternshipGormRepo
func NewInternshipGormRepo(db *gorm.DB) internship.InternshipRepository {
	return &InternshipGormRepo{conn: db}
}

// Internships returns a list of all available interships
func (igr *InternshipGormRepo) Internships() ([]entity.Internship, []error) {
	intern := []entity.Internship{}
	field := []entity.Field{}
	errs := igr.conn.Find(&intern).GetErrors() //pass container for result to Find()
	count := len(intern)

	for i := 0; i < count; i++ {
		errs := igr.conn.Model(intern[i]).Related(&field, "FieldsReq").GetErrors()
		intern[i].FieldsReq = field
		if len(errs) > 0 {
			return nil, errs
		}
	}

	if len(errs) > 0 {
		return nil, errs
	}

	return intern, errs
}

//CompanyInternships returns internships under a company
func (igr *InternshipGormRepo) CompanyInternships(company *entity.CompanyDetail) ([]entity.Internship, []error) {
	compInterns := []entity.Internship{}
	errs := igr.conn.Model(company).Related(&compInterns, "FieldsReq").GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return compInterns, errs
}

// Internship finds an internship with a given id
func (igr *InternshipGormRepo) Internship(id uint) (*entity.Internship, []error) {
	intern := entity.Internship{}
	field := []entity.Field{}

	errs := igr.conn.First(&intern, id).GetErrors()
	errs = igr.conn.Model(intern).Related(&field, "FieldsReq").GetErrors()
	intern.FieldsReq = field
	// errs = igr.conn.Where("internship_id = ?", intern.ID).Find(&field).GetErrors()
	// intern.FieldsReq = field
	// if len(errs) != 0 {
	// 	return nil, errs
	// }
	if len(errs) > 0 {
		return nil, errs
	}
	return &intern, errs
}

//UpdateInternship updates a given internship
func (igr *InternshipGormRepo) UpdateInternship(internship *entity.Internship) (*entity.Internship, []error) {
	intern := internship //the new role isn't directly transferred, but transferred to a variable and which is then transferred to another. what?
	errs := igr.conn.Save(intern).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return intern, errs
}

// DeleteInternship deletes a given internship
func (igr *InternshipGormRepo) DeleteInternship(id uint) (*entity.Internship, []error) {
	intern, errs := igr.Internship(id)
	errs = igr.conn.Delete(intern, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return intern, errs
}

//StoreInternship creates a new internship
func (igr *InternshipGormRepo) StoreInternship(internship *entity.Internship) (*entity.Internship, []error) {
	intern := internship
	errs := igr.conn.Create(intern).GetErrors() // this now has an id, auto
	if len(errs) > 0 {
		return nil, errs
	}
	return intern, errs
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
