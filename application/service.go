package application

import "github.com/abdimussa87/intern-seek-RestAPI/entity"

//ApplicationService specifies application related services

type ApplicationService interface {
	StoreApplication(application *entity.Application) (*entity.Application, []error)
	UpdateApplication(application *entity.Application) (*entity.Application, []error)
	DeleteApplication(id uint) (*entity.Application, []error)
	Applications() ([]entity.Application, []error)
	Application(id uint) (*entity.Application, []error)
}

type StatusService interface {
	StoreStatus(status *entity.Status) (*entity.Status, []error)
	UpdateStatus(Status *entity.Status) (*entity.Status, []error)
	DeleteStatus(id uint) (*entity.Status, []error)
	Statuses() ([]entity.Status, []error)
	Status(id uint) (*entity.Status, []error)
}
