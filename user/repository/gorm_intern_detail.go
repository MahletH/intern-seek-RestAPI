package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/lensabillion/intern-seek-RestAPI/entity"
	"github.com/lensabillion/intern-seek-RestAPI/user"
)



// InternGormRepo Implements the user.InternRepository interface
type InternGormRepo struct {
	conn *gorm.DB
}
// NewInternGormRepoImpl creates a new object of InternGormRepo
func NewInternGormRepoImpl(db *gorm.DB) user.InternRepository {
	return &InternGormRepo{conn: db}
}
//storeIntern stores new personal_detail in the database
func (internRepo *InternGormRepo) StoreIntern(intern *entity.PersonalDetails) (*entity.PersonalDetails, []error) {
	int := intern
	errs := internRepo.conn.Create(int).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return int, errs

}
// UpdateIntern updates a given personal_detail in the database
func (internRepo InternGormRepo) UpdateIntern(intern *entity.PersonalDetails) (*entity.PersonalDetails, []error) {
	int := intern
	errs := internRepo.conn.Save(int).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return int, errs
}
// DeleteIntern deletes a given personal_detail from the database
func (intRepo *InternGormRepo) DeleteIntern(id uint) (*entity.PersonalDetails, []error) {
	int, errs := intRepo.Intern(id)
	if len(errs) > 0 {
		return nil, errs
	}
	errs = intRepo.conn.Delete(int, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return int, errs
}
//  Interns return all personal_details from the database
func (internRepo *InternGormRepo) Interns() ([]entity.PersonalDetails, []error) {
	interns := []entity.PersonalDetails{}
	errs := internRepo.conn.Find(&interns).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return interns, errs

}
//Intern retrieves an Intern_detail by its id from the database
func (internRepo *InternGormRepo) Intern(id uint) (*entity.PersonalDetails, []error) {
	intern := entity.PersonalDetails{}
	errs:=internRepo.conn.First(&intern,id).GetErrors()
	if len(errs) > 0{
		return nil,errs
	}
	return &intern,errs
}

// GetInternByUserId retrieves a PersonalDetails by its user-id from the database
func (intRepo *InternGormRepo) GetInternByUserId(id uint) (*entity.PersonalDetails, []error) {
	intern := entity.PersonalDetails{}
	errs := intRepo.conn.Where("user_id=?", id).First(&intern).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &intern, errs
}
