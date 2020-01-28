package service

import (
	"github.com/MahletH/intern-seek-RestAPI/application"
	"github.com/MahletH/intern-seek-RestAPI/entity"
)

//ApplicationServiceImpl implements application.Applicationservice interface
type ApplicationServiceImpl struct {
	appRepo application.ApplicationRepository
}

//NewApplicationServiceImpl returns new ApplicationServiceImpl
func NewApplicationServiceImpl(applRepo application.ApplicationRepository) *ApplicationServiceImpl {
	return &ApplicationServiceImpl{appRepo: applRepo}
}

//StoreApplication stores a Application given a Application
func (usi ApplicationServiceImpl) StoreApplication(app *entity.Application) (*entity.Application, []error) {
	apps, errs := usi.appRepo.StoreApplication(app)
	if len(errs) > 0 {
		return nil, errs
	}
	return apps, errs
}

//UpdateApplication updates a Application given a Application
func (usi ApplicationServiceImpl) UpdateApplication(app *entity.Application) (*entity.Application, []error) {
	apps, errs := usi.appRepo.UpdateApplication(app)
	if len(errs) > 0 {
		return nil, errs
	}
	return apps, errs
}

//DeleteApplication deletes a Application given an id
func (usi ApplicationServiceImpl) DeleteApplication(id uint) (*entity.Application, []error) {
	apps, errs := usi.appRepo.DeleteApplication(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return apps, nil
}

//Applications returns  list of Applications
func (usi ApplicationServiceImpl) Applications() ([]entity.Application, []error) {
	apps, errs := usi.appRepo.Applications()
	if len(errs) > 0 {
		return nil, errs
	}
	return apps, nil
}

//Application returns a Application given an id
func (usi ApplicationServiceImpl) Application(id uint) (*entity.Application, []error) {
	app, err := usi.appRepo.Application(id)
	if len(err) > 0 {
		return app, err
	}
	return app, nil
}
