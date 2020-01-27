package repository

import (
	"github.com/MahletH/intern-seek-RestAPI/entity"
	"github.com/MahletH/intern-seek-RestAPI/internship"

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
	fieldIDs := []entity.FieldOfStudyRequired{}
	field := []entity.Fields{}
	errs := igr.conn.Find(&intern).GetErrors() //pass container for result to Find()
	count := len(intern)

	for i := 0; i < count; i++ {
		errs = igr.conn.Where("internship_id = ?", intern[i].ID).Find(&fieldIDs).GetErrors()
		for _, value := range fieldIDs {
			errs = igr.conn.Where("id = ?", value.FieldID).Find(&fields).GetErrors()

		}

		intern[i].FieldsReq = fields
		if len(errs) != 0 {
			return nil, errs
		}
	}

	if len(errs) > 0 {
		return nil, errs
	}

	return intern, errs
}

// Internship finds an internship with a given id
func (igr *InternshipGormRepo) Internship(id uint) (*entity.Internship, []error) {
	intern := entity.Internship{}
	fieldIDs := []entity.FieldOfStudyRequired{}
	fields := []entity.Fields{}

	errs := igr.conn.First(&intern, id).GetErrors()
	errs = igr.conn.Where("internship_id = ?", intern.ID).Find(&fieldIDs).GetErrors()
	for _, value := range fieldIDs {
		errs = igr.conn.Where("id = ?", value.FieldID).Find(&fields).GetErrors()

	}

	intern.FieldsReq = fields

	if len(errs) != 0 {
		return nil, errs
	}
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
	field := []entity.Fields{}

	errs = igr.conn.Where("internship_id = ?", id).Delete(&field).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
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
