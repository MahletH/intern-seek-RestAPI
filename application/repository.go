package application

import "github.com/abdimussa87/intern-seek-RestAPI/entity"

// ApplicationRepository specifies application related database operations

type ApplicationRepository interface {
	StoreApplication(application *entity.Application) (*entity.Application, []error)
	UpdateApplication(application *entity.Application) (*entity.Application, []error)
	DeleteApplication(id uint) (*entity.Application, []error)
	Applications() ([]entity.Application, []error)
	Application(id uint) (*entity.Application, []error)
}

type StatusRepository interface {
	StoreStatus(status *entity.Status) (*entity.Status, []error)
	UpdateStatus(Status *entity.Status) (*entity.Status, []error)
	DeleteStatus(id uint) (*entity.Status, []error)
	Statuses() ([]entity.Status, []error)
	Status(id uint) (*entity.Status, []error)
}
