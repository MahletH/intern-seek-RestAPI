package user

import "github.com/abdimussa87/Intern-Seek-Version-1/entity"

// UserRepository specifies user related database operations
type UserRepository interface {
	Users() ([]entity.User, []error)
	User(id uint) (*entity.User, []error)
	UpdateUser(user *entity.User) (*entity.User, []error)
	UserByUsername(username string) (*entity.User, error)
	DeleteUser(id uint) (*entity.User, []error)
	UserByUsernameAndPassword(username string, password string) (*entity.User, error)
	StoreUser(user *entity.User) (*entity.User, []error)
}

type CompanyRepository interface {
	StoreCompany(company *entity.CompanyDetail) (*entity.CompanyDetail, []error)
	UpdateCompany(company *entity.CompanyDetail) (*entity.CompanyDetail, []error)
	DeleteCompany(id uint) (*entity.CompanyDetail, []error)
	Companies() ([]entity.CompanyDetail, []error)
	Company(id uint) (*entity.CompanyDetail, []error)
}
type ApplicationRepository interface{
	StoreApplication(application *entity.Application)(*entity.Application,[]error)
	UpdateApplication(application *entity.Application)(*entity.Application,[]error)
	DeleteApplication(id uint)(*entity.Application,[]error)
	Applications()([]entity.Application,[]error)
	Application(id uint)(*entity.Application,[]error)
}
