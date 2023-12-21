package repository

import (
	"errors"

	"finalprojekmygram/apps/model/domain"

	utilities "finalprojekmygram/internal/utilities"

	"gorm.io/gorm"
)

type UserRepository interface {
	Register(user *domain.User) (err error)
	Login(user *domain.User) (err error)
	Update(user *domain.User) (updatedUser domain.User, err error)
	Delete(id uint) (err error)
}

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{DB: db}
}

func (userRepository *UserRepositoryImpl) Register(user *domain.User) (err error) {
	if err = userRepository.DB.Create(&user).Error; err != nil {
		return err
	}

	return
}

func (userRepository *UserRepositoryImpl) Login(user *domain.User) (err error) {
	password := user.Password

	err = userRepository.DB.Where("email = ?", user.Email).Take(&user).Error
	isValid := utilities.ComparePasswordMethod([]byte(user.Password), []byte(password))

	if err != nil || !isValid {
		return errors.New("invalid email or password")
	}

	return
}

func (userRepository *UserRepositoryImpl) Update(user *domain.User) (updatedUser domain.User, err error) {

	if err = userRepository.DB.First(&updatedUser, user.ID).Error; err != nil {
		return
	}

	if err = userRepository.DB.Model(&updatedUser).Updates(user).Error; err != nil {
		return
	}

	return
}

func (userRepository *UserRepositoryImpl) Delete(id uint) (err error) {

	if err = userRepository.DB.First(&domain.User{}, id).Error; err != nil {
		return
	}

	if err = userRepository.DB.Delete(&domain.User{}, id).Error; err != nil {
		return
	}

	return
}
