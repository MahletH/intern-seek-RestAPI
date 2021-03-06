package service

import (
	"github.com/abdimussa87/intern-seek-RestAPI/entity"
	"github.com/abdimussa87/intern-seek-RestAPI/user"
)

// UserRoleService implements menu.UserRoleService interface
type UserRoleService struct {
	roleRepo user.UserRoleRepository
}

// NewRoleService  returns new UserRoleService
func NewUserRoleService(RoleRepo user.UserRoleRepository) user.UserRoleService {
	return &UserRoleService{roleRepo: RoleRepo}
}

//UserRole retrievs a given user role by its id
func (rs *UserRoleService) UserRole(id uint) (*entity.UserRole, []error) {
	rl, errs := rs.roleRepo.UserRole(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return rl, errs

}

// DeleteUserRole deletes a given user role
func (rs *UserRoleService) DeleteUserRole(id uint) (*entity.UserRole, []error) {

	rl, errs := rs.roleRepo.DeleteUserRole(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return rl, errs
}

// StoreRole stores a given user role
func (rs *UserRoleService) StoreUserRole(role *entity.UserRole) (*entity.UserRole, []error) {

	rl, errs := rs.roleRepo.StoreUserRole(role)
	if len(errs) > 0 {
		return nil, errs
	}
	return rl, errs
}
