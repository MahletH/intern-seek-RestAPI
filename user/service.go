package user

import "github.com/abdimussa87/intern-seek-RestAPI/entity"

//UserService specifies user related services
type UserService interface {
	Users() ([]entity.User, []error)
	User(id uint) (*entity.User, []error)
	UpdateUser(user *entity.User) (*entity.User, []error)
	DeleteUser(id uint) (*entity.User, []error)
	StoreUser(user *entity.User) (*entity.User, []error)
	UserByUsernameAndPassword(username string, password string) (*entity.User, error)
	UserByUsername(username string) (*entity.User, error)
}
type CompanyService interface {
	StoreCompany(company *entity.CompanyDetail) (*entity.CompanyDetail, []error)
	UpdateCompany(company *entity.CompanyDetail) (*entity.CompanyDetail, []error)
	DeleteCompany(id uint) (*entity.CompanyDetail, []error)
	Companies() ([]entity.CompanyDetail, []error)
	Company(id uint) (*entity.CompanyDetail, []error)
	GetCompanyByUserId(id uint) (*entity.CompanyDetail, []error)
}
type UserRoleService interface {
	UserRole(id uint) (*entity.UserRole, []error)

	DeleteUserRole(id uint) (*entity.UserRole, []error)
	StoreUserRole(role *entity.UserRole) (*entity.UserRole, []error)
}
type InternService interface {
	StoreIntern(intern *entity.PersonalDetails) (*entity.PersonalDetails, []error)
	UpdateIntern(intern *entity.PersonalDetails) (*entity.PersonalDetails, []error)
	DeleteIntern(id uint) (*entity.PersonalDetails, []error)
	Interns() ([]entity.PersonalDetails, []error)
	Intern(id uint) (*entity.PersonalDetails, []error)
	GetInternByUserId(id uint) (*entity.PersonalDetails, []error)
	InternFields(intern *entity.PersonalDetails) ([]entity.Field, []error)
}
