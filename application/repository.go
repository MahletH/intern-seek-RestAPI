package application

import "github.com/MahletH/intern-seek-RestAPI/entity"

// ApplicationRepository specifies application related database operations

type ApplicationRepository interface {
	StoreApplication(application *entity.Application) (*entity.Application, []error)
	UpdateApplication(application *entity.Application) (*entity.Application, []error)
	DeleteApplication(id uint) (*entity.Application, []error)
	Applications() ([]entity.Application, []error)
	Application(id uint) (*entity.Application, []error)
}
