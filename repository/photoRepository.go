package repository

import (
	"myGram/entities"

	"gorm.io/gorm"
)

type PhotoRepository interface {
	GetPhotos() ([]entities.Photo, error)
	GetPhotoByID(id uint) (entities.Photo, error)
	CreatePhoto(data entities.Photo) (entities.Photo, error)
	UpdatePhoto(data entities.Photo) (entities.Photo, error)
	DeletePhoto(PhotoID uint) error
}

type photoRepository struct {
	db *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) PhotoRepository {
	return &photoRepository{db: db}
}

func (r *photoRepository) GetPhotos() ([]entities.Photo, error) {
	var photo []entities.Photo
	err := r.db.Preload("User").Find(&photo).Error
	if err != nil {
		return []entities.Photo{}, err
	}
	return photo, nil
}

func (r *photoRepository) GetPhotoByID(id uint) (entities.Photo, error) {
	var photo entities.Photo
	err := r.db.Preload("User").Where("id = ?", id).First(&photo).Error
	if err != nil {
		return entities.Photo{}, err
	}
	return photo, nil
}

func (r *photoRepository) CreatePhoto(data entities.Photo) (entities.Photo, error) {
	err := r.db.Create(&data).Error
	if err != nil {
		return entities.Photo{}, err
	}
	return data, nil
}

func (r *photoRepository) UpdatePhoto(data entities.Photo) (entities.Photo, error) {
	err := r.db.Model(&data).Where("id = ?", data.ID).Updates(entities.Photo{Title: data.Title, Caption: data.Caption, PhotoURL: data.PhotoURL}).Error
	if err != nil {
		return entities.Photo{}, err
	}
	return data, nil
}

func (r *photoRepository) DeletePhoto(PhotoID uint) error {
	photo := entities.Photo{}
	photo.ID = uint(PhotoID)
	err := r.db.First(&photo).Where("id = ?", PhotoID).Delete(&photo).Error
	if err != nil {
		return err
	}
	return nil
}
