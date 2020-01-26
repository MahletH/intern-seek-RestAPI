package repository
import (
"github.com/jinzhu/gorm"
"github.com/lensabillion/Project/entity"
"github.com/lensabillion/Project/user"
)

// applicationGormRepo Implements the user.ApplicationRepository interface
type ApplicationGormRepo struct {
conn *gorm.DB
}

// NewApplicationGormRepoImpl creates a new object of ApplicationGormRepo
func NewApplicationGormRepoImpl(db *gorm.DB) user.ApplicationRepository{
return &ApplicationGormRepo{conn: db}
}

// Applications return all Applications from the database
func (appRepo *ApplicationGormRepo) Applications() ([]entity.Application, []error) {
applications := []entity.Application{}
errs := appRepo.conn.Find(&applications).GetErrors()
if len(errs) > 0 {
return nil, errs
}
return applications, errs
}

// Application retrieves a Application by its id from the database
func (appRepo *ApplicationGormRepo) Application(id uint) (*entity.Application, []error) {
application := entity.Application{}
errs := appRepo.conn.First(&application, id).GetErrors()
if len(errs) > 0 {
return nil, errs
}
return &application, errs
}

// UpdateApplication updates a given Appliction in the database
func (appRepo *ApplicationGormRepo) UpdateApplication(application *entity.Application) (*entity.Application, []error) {
app := application
errs := appRepo.conn.Save(app).GetErrors()
if len(errs) > 0 {
return nil, errs
}
return app, errs
}

// DeleteApplication deletes a given Application from the database
func (appRepo *ApplicationGormRepo) DeleteApplication(id uint) (*entity.Application, []error) {
app, errs := appRepo.Application(id)
if len(errs) > 0 {
return nil, errs
}
errs = appRepo.conn.Delete(app, id).GetErrors()
if len(errs) > 0 {
return nil, errs
}
return app, errs
}

// StoreApplication stores a new Application into the database
func (appRepo *ApplicationGormRepo) StoreApplication (application *entity.Application) (*entity.Application, []error) {
app := application
errs := appRepo.conn.Create(app).GetErrors()
if len(errs) > 0 {
return nil, errs
}
return app, errs
}
