package service

import (
	"github.com/nebyubeyene/Intern-Seek-Version-1/entity"
	"github.com/nebyubeyene/Intern-Seek-Version-1/user"
)

// UserRoleService implements menu.UserRoleService interface
type UserRoleService struct {
	roleRepo user.UserRoleRepository
}

// NewRoleService  returns new UserRoleService
func NewRoleService(RoleRepo user.UserRoleRepository) user.UserRoleService {
	return &UserRoleService{roleRepo: RoleRepo}
}

// Role retrievs a given user role by its id
func (rs *UserRoleService) UserRole(id uint) (*entity.UserRole, []error) {
	rl, errs := rs.roleRepo.UserRole(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return rl, errs

}

// DeleteRole deletes a given user role
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
