package service

import (
	"finalprojekmygram/apps/model/domain"
	"finalprojekmygram/apps/repository"
)

type UserService interface {
	Register(user *domain.User) (err error)
	Login(user *domain.User) (err error)
	Update(user *domain.User) (updatedUser domain.User, err error)
	Delete(id uint) (err error)
}

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &UserServiceImpl{UserRepository: userRepository}
}

func (userService *UserServiceImpl) Register(user *domain.User) (err error) {
	if err = userService.UserRepository.Register(user); err != nil {
		return err
	}

	return
}

func (userService *UserServiceImpl) Login(user *domain.User) (err error) {
	if err = userService.UserRepository.Login(user); err != nil {
		return err
	}

	return
}

func (userService *UserServiceImpl) Update(user *domain.User) (u domain.User, err error) {
	if u, err = userService.UserRepository.Update(user); err != nil {
		return u, err
	}

	return u, nil
}

func (userService *UserServiceImpl) Delete(id uint) (err error) {
	if err = userService.UserRepository.Delete(id); err != nil {
		return err
	}

	return
}
