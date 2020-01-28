package service

import (
	"github.com/abdimussa87/intern-seek-RestAPI/entity"
	"github.com/abdimussa87/intern-seek-RestAPI/field"
)

// FieldService implements menu.FieldService interface
type FieldService struct {
	fieldRepo field.FieldRepository
}

// NewRoleService  returns new FieldService
func NewFieldService(FieldRepo field.FieldRepository) field.FieldService {
	return &FieldService{fieldRepo: FieldRepo}
}

//Field retrievs a given user role by its id
func (fs *FieldService) Field(id uint) (*entity.Field, []error) {
	fl, errs := fs.fieldRepo.Field(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return fl, errs

}

// DeleteField deletes a given user role
func (fs *FieldService) DeleteField(id uint) (*entity.Field, []error) {

	fl, errs := fs.fieldRepo.DeleteField(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return fl, errs
}

// StoreRole stores a given user role
func (fs *FieldService) StoreField(field *entity.Field) (*entity.Field, []error) {

	fl, errs := fs.fieldRepo.StoreField(field)
	if len(errs) > 0 {
		return nil, errs
	}
	return fl, errs
}

// UpdateOrder updates a given order
func (fs *FieldService) UpdateField(field *entity.Field) (*entity.Field, []error) {
	fld, errs := fs.fieldRepo.UpdateField(field)
	if len(errs) > 0 {
		return nil, errs
	}
	return fld, errs
}

// Items returns all stored food menu items
func (fs *FieldService) Fields() ([]entity.Field, []error) {
	flds, errs := fs.fieldRepo.Fields()
	if len(errs) > 0 {
		return nil, errs
	}
	return flds, errs
}

func (fs *FieldService) FieldInternships(field *entity.Field) ([]entity.Internship, []error) {
	fldInternships, errs := fs.fieldRepo.FieldInternships(field)
	if len(errs) > 0 {
		return nil, errs
	}
	return fldInternships, errs
}

func (fs *FieldService) FieldInterns(field *entity.Field) ([]entity.Intern, []error) {
	fldInterns, errs := fs.fieldRepo.FieldInterns(field)
	if len(errs) > 0 {
		return nil, errs
	}
	return fldInterns, errs
}

// GetFieldbyName, retrieves a field its name
func (fs *FieldService) GetFieldbyName(name string) (*entity.Field, []error) {
	fld, errs := fs.fieldRepo.GetFieldbyName(name)
	if len(errs) > 0 {
		return nil, errs
	}
	return fld, errs
}
