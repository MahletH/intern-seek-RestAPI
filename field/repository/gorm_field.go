package repository

import (
	"github.com/abdimussa87/intern-seek-RestAPI/entity"
	"github.com/abdimussa87/intern-seek-RestAPI/field"
	"github.com/jinzhu/gorm"
)

type FieldGormRepo struct {
	conn *gorm.DB
}

// NewUserRoleGormRepo returns a new a new object of FieldGormRepo
func NewFieldGormRepo(db *gorm.DB) field.FieldRepository {
	return &FieldGormRepo{conn: db}
}

// Roles returns all user roles stored in the database

// Field retrieves a role by its id from the database
func (fieldRepo *FieldGormRepo) Field(id uint) (*entity.Field, []error) {
	field := entity.Field{}
	errs := fieldRepo.conn.First(&field, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &field, errs
}

// DeleteRole deletes a given user role from the database
func (fieldRepo *FieldGormRepo) DeleteField(id uint) (*entity.Field, []error) {
	f, errs := fieldRepo.Field(id)
	if len(errs) > 0 {
		return nil, errs
	}
	errs = fieldRepo.conn.Delete(f, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return f, errs
}

// StoreRole stores a given user role in the database
func (fieldRepo *FieldGormRepo) StoreField(field *entity.Field) (*entity.Field, []error) {
	f := field
	errs := fieldRepo.conn.Create(f).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return f, errs
}
func (fieldRepo *FieldGormRepo) FieldInterns(field *entity.Field) ([]entity.Intern, []error) {
	interns := []entity.Intern{}
	errs := fieldRepo.conn.Model(field).Related(&interns, "Interns").GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return interns, errs
}
func (fieldRepo *FieldGormRepo) UpdateField(field *entity.Field) (*entity.Field, []error) {
	fld := field
	errs := fieldRepo.conn.Save(fld).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return fld, errs
}

// Fields return all fields from the database
func (fieldRepo *FieldGormRepo) Fields() ([]entity.Field, []error) {
	fields := []entity.Field{}
	errs := fieldRepo.conn.Find(&fields).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return fields, errs
}

func (fieldRepo *FieldGormRepo) FieldInternships(field *entity.Field) ([]entity.Internship, []error) {
	fieldInternships := []entity.Internship{}

	errs := fieldRepo.conn.Model(field).Related(&fieldInternships, "Internships").GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return fieldInternships, errs
}

// GetCompanyByUserId retrieves a company_detail by its user-id from the database
func (fieldRepo *FieldGormRepo) GetFieldbyName(name string) (*entity.Field, []error) {
	field := entity.Field{}
	errs := fieldRepo.conn.Where("name=?", name).First(&field).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &field, errs
}
