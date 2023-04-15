package service

import (
	"myGram/entities"
	"myGram/repository"
)

type PhotoService interface {
	GetAllPhotos() ([]entities.ResponsePhoto, error)
	GetPhotoByID(photoID uint) (entities.ResponsePhoto, error)
	CreatePhoto(data entities.Photo) (entities.Photo, error)
	UpdatePhoto(data entities.Photo, photoID uint) (entities.Photo, error)
	DeletePhoto(photoID uint) error
}

func NewPhotoService(photoRepo repository.PhotoRepository) PhotoService {
	return &photoService{photoRepo: photoRepo}
}

type photoService struct {
	photoRepo repository.PhotoRepository
}

func (service *photoService) GetAllPhotos() ([]entities.ResponsePhoto, error) {
	resPhotos, err := service.photoRepo.GetPhotos()

	if err != nil {
		return []entities.ResponsePhoto{}, err
	}
	var response []entities.ResponsePhoto
	for _, photo := range resPhotos {
		tempResp := entities.ResponsePhoto{}
		tempResp.ID = photo.ID
		tempResp.Title = photo.Title
		tempResp.Caption = photo.Caption
		tempResp.PhotoURL = photo.PhotoURL
		tempResp.CreatedAt = *photo.CreatedAt
		tempResp.UpdatedAt = *photo.UpdatedAt
		tempResp.User.Username = photo.User.Username
		tempResp.User.Email = photo.User.Email
		response = append(response, tempResp)
	}

	return response, nil
}

func (service *photoService) GetPhotoByID(photoID uint) (entities.ResponsePhoto, error) {
	resPhotos, err := service.photoRepo.GetPhotoByID(photoID)
	if err != nil {
		return entities.ResponsePhoto{}, err
	}
	var response entities.ResponsePhoto
	response.ID = resPhotos.ID
	response.Title = resPhotos.Title
	response.Caption = resPhotos.Caption
	response.PhotoURL = resPhotos.PhotoURL
	response.CreatedAt = *resPhotos.CreatedAt
	response.UpdatedAt = *resPhotos.UpdatedAt
	response.User.Username = resPhotos.User.Username
	response.User.Email = resPhotos.User.Email
	return response, nil
}

func (service *photoService) CreatePhoto(data entities.Photo) (entities.Photo, error) {
	create, err := service.photoRepo.CreatePhoto(data)
	if err != nil {
		return entities.Photo{}, err
	}
	return create, nil
}

func (service *photoService) UpdatePhoto(data entities.Photo, photoID uint) (entities.Photo, error) {
	entityPhoto := entities.Photo{}
	entityPhoto.ID = uint(photoID)
	entityPhoto.Title = data.Title
	entityPhoto.Caption = data.Caption
	entityPhoto.PhotoURL = data.PhotoURL
	getPhoto, err := service.photoRepo.GetPhotoByID(photoID)
	if err != nil {
		return entities.Photo{}, err
	}
	if data.Title == "" {
		entityPhoto.Title = getPhoto.Title
	}
	if data.Caption == "" {
		entityPhoto.Caption = getPhoto.Caption
	}
	if data.PhotoURL == "" {
		entityPhoto.PhotoURL = getPhoto.PhotoURL
	}
	update, err := service.photoRepo.UpdatePhoto(entityPhoto)
	if err != nil {
		return entities.Photo{}, err
	}
	return update, nil
}

func (service *photoService) DeletePhoto(photoID uint) error {
	err := service.photoRepo.DeletePhoto(photoID)
	if err != nil {
		return err
	}
	return nil
}
