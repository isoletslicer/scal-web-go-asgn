package service

import (
	"finalprojekmygram/apps/model/domain"
	"finalprojekmygram/apps/repository"
)

type PhotoService interface {
	Create(photo *domain.Photo) (err error)
	GetAll() (photos []domain.Photo, err error)
	GetOne(id uint) (photo domain.Photo, err error)
	Update(photo domain.Photo) (updatedPhoto domain.Photo, err error)
	Delete(id uint) (err error)
}

type PhotoServiceImpl struct {
	PhotoRepository repository.PhotoRepository
}

func NewPhotoService(photoRepository repository.PhotoRepository) PhotoService {
	return &PhotoServiceImpl{PhotoRepository: photoRepository}
}

func (photoService *PhotoServiceImpl) Create(photo *domain.Photo) (err error) {

	if err = photoService.PhotoRepository.Create(photo); err != nil {
		return
	}

	return
}

func (photoService *PhotoServiceImpl) GetAll() (photos []domain.Photo, err error) {

	if photos, err = photoService.PhotoRepository.GetAll(); err != nil {
		return
	}

	return
}

func (photoService *PhotoServiceImpl) GetOne(id uint) (photos domain.Photo, err error) {

	if photos, err = photoService.PhotoRepository.GetOne(id); err != nil {
		return
	}

	return
}

func (photoService *PhotoServiceImpl) Update(photo domain.Photo) (updatedPhoto domain.Photo, err error) {

	if updatedPhoto, err = photoService.PhotoRepository.Update(photo); err != nil {
		return
	}

	return
}

func (photoService *PhotoServiceImpl) Delete(id uint) (err error) {

	if err = photoService.PhotoRepository.Delete(id); err != nil {
		return
	}

	return
}
