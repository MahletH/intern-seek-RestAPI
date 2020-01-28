package repository

import (
	"github.com/abdimussa87/intern-seek-RestAPI/entity"
	"github.com/abdimussa87/intern-seek-RestAPI/user"
	"github.com/jinzhu/gorm"
)

type UserRoleGormRepo struct {
	conn *gorm.DB
}

// NewUserRoleGormRepo returns a new a new object of UserRoleGormRepo
func NewUserRoleGormRepo(db *gorm.DB) user.UserRoleRepository {
	return &UserRoleGormRepo{conn: db}
}

// Roles returns all user roles stored in the database

// UserRole retrieves a role by its id from the database
func (userRolerepo *UserRoleGormRepo) UserRole(id uint) (*entity.UserRole, []error) {
	role := entity.UserRole{}
	errs := userRolerepo.conn.First(&role, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &role, errs
}

// DeleteRole deletes a given user role from the database
func (userRolerepo *UserRoleGormRepo) DeleteUserRole(id uint) (*entity.UserRole, []error) {
	r, errs := userRolerepo.UserRole(id)
	if len(errs) > 0 {
		return nil, errs
	}
	errs = userRolerepo.conn.Delete(r, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return r, errs
}

// StoreRole stores a given user role in the database
func (userRolerepo *UserRoleGormRepo) StoreUserRole(role *entity.UserRole) (*entity.UserRole, []error) {
	r := role
	errs := userRolerepo.conn.Create(r).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return r, errs
}
