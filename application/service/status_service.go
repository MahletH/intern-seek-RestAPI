package service

import (
	"github.com/abdimussa87/intern-seek-RestAPI/application"
	"github.com/abdimussa87/intern-seek-RestAPI/entity"
)

// StatusService implements application.StatusService interface
type StatusService struct {
	statRepo application.StatusRepository
}

// NewStatusService  returns a new StatusService object
func NewStatusService(statusRepository application.StatusRepository) *StatusService {
	return &StatusService{statRepo: statusRepository}
}

// Statuses returns all stored statuses
func (ss *StatusService) Statuses() ([]entity.Status, []error) {
	statuses, errs := ss.statRepo.Statuses()
	if len(errs) > 0 {
		return nil, errs
	}
	return statuses, errs
}

// Status retrieves an application Status by its id
func (ss *StatusService) Status(id uint) (*entity.Status, []error) {
	status, errs := ss.statRepo.Status(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return status, errs
}

// UpdateStatus updates  a given status
func (ss *StatusService) UpdateStatus(status *entity.Status) (*entity.Status, []error) {
	Status, errs := ss.statRepo.UpdateStatus(status)
	if len(errs) > 0 {
		return nil, errs
	}
	return Status, errs
}

// DeleteStatus deletes a given status
func (ss *StatusService) DeleteStatus(id uint) (*entity.Status, []error) {
	Status, errs := ss.statRepo.DeleteStatus(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return Status, errs
}

// StoreStatus stores a given status
func (ss *StatusService) StoreStatus(status *entity.Status) (*entity.Status, []error) {
	Status, errs := ss.statRepo.StoreStatus(status)
	if len(errs) > 0 {
		return nil, errs
	}
	return Status, errs
}
