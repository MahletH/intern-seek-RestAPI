package repository

import (
	"github.com/abdimussa87/intern-seek-RestAPI/application"
	"github.com/abdimussa87/intern-seek-RestAPI/entity"
	"github.com/jinzhu/gorm"
)

// StatusGormRepo Implements the user.ApplicationRepository interface
type StatusGormRepo struct {
	conn *gorm.DB
}

// NewStatusGormRepo creates a new object of CompanyGormRepo
func NewStatusGormRepo(db *gorm.DB) application.StatusRepository {
	return &StatusGormRepo{conn: db}
}

// Statuses return all Statuses from the database
func (statRepo *StatusGormRepo) Statuses() ([]entity.Status, []error) {
	Statuses := []entity.Status{}
	errs := statRepo.conn.Find(&Statuses).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return Statuses, errs
}

// Status retrieves a Status by its id from the database
func (statRepo *StatusGormRepo) Status(id uint) (*entity.Status, []error) {
	Status := entity.Status{}
	errs := statRepo.conn.First(&Status, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &Status, errs
}

// UpdateStatus updates a given Status in the database
func (statRepo *StatusGormRepo) UpdateStatus(status *entity.Status) (*entity.Status, []error) {
	Status := status
	errs := statRepo.conn.Save(Status).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return Status, errs
}

// DeleteStatus deletes a given Status from the database
func (statRepo *StatusGormRepo) DeleteStatus(id uint) (*entity.Status, []error) {
	status, errs := statRepo.Status(id)
	if len(errs) > 0 {
		return nil, errs
	}
	errs = statRepo.conn.Delete(status, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return status, errs
}

// StoreStatus stores a new Status into the database
func (statRepo *StatusGormRepo) StoreStatus(status *entity.Status) (*entity.Status, []error) {
	Status := status
	errs := statRepo.conn.Create(Status).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return Status, errs
}
