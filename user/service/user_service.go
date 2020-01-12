package service

import (
	"github.com/nebyubeyene/Intern-Seek-Version-1/entity"
	"github.com/nebyubeyene/Intern-Seek-Version-1/user"
)

// UserService implements menu.UserService interface
type UserService struct {
	userRepo user.UserRepository
}

// NewUserService  returns a new UserService object
func NewUserServiceImpl(userRepository user.UserRepository) user.UserService {
	return &UserService{userRepo: userRepository}
}

// Users returns all stored application users
func (us *UserService) Users() ([]entity.User, []error) {
	usrs, errs := us.userRepo.Users()
	if len(errs) > 0 {
		return nil, errs
	}
	return usrs, errs
}

//UserByUsernameAndPassword returns a user given a username and password if it exists
func (us *UserService) UserByUsernameAndPassword(username string, password string) (*entity.User, error) {
	usr, err := us.userRepo.UserByUsernameAndPassword(username, password)
	if err != nil {
		return nil, err
	}
	return usr, nil
}

// User retrieves an application user by its id
func (us *UserService) User(id uint) (*entity.User, []error) {
	usr, errs := us.userRepo.User(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}

// UpdateUser updates  a given application user
func (us *UserService) UpdateUser(user *entity.User) (*entity.User, []error) {
	usr, errs := us.userRepo.UpdateUser(user)
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}

// DeleteUser deletes a given application user
func (us *UserService) DeleteUser(id uint) (*entity.User, []error) {
	usr, errs := us.userRepo.DeleteUser(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}

// StoreUser stores a given application user
func (us *UserService) StoreUser(user *entity.User) (*entity.User, []error) {
	usr, errs := us.userRepo.StoreUser(user)
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}
