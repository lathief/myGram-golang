package repository

import (
	"myGram/entities"

	"gorm.io/gorm"
)

type SocialMediaRepository interface {
	GetSosmed() ([]entities.SocialMedia, error)
	GetSosmedByID(id uint) (entities.SocialMedia, error)
	CreateSosmed(data entities.SocialMedia) (entities.SocialMedia, error)
	UpdateSosmed(data entities.SocialMedia) (entities.SocialMedia, error)
	DeleteSosmed(socialMediaID uint) error
}

type socialMediaRepository struct {
	db *gorm.DB
}

func NewSosmedRepository(db *gorm.DB) SocialMediaRepository {
	return &socialMediaRepository{db: db}
}

func (r *socialMediaRepository) GetSosmed() ([]entities.SocialMedia, error) {
	var sosmed []entities.SocialMedia
	err := r.db.Preload("User").Find(&sosmed).Error
	if err != nil {
		return []entities.SocialMedia{}, err
	}
	return sosmed, nil
}

func (r *socialMediaRepository) GetSosmedByID(id uint) (entities.SocialMedia, error) {
	var sosmed entities.SocialMedia
	err := r.db.Preload("User").Where("id = ?", id).First(&sosmed).Error
	if err != nil {
		return entities.SocialMedia{}, err
	}
	return sosmed, nil
}

func (r *socialMediaRepository) CreateSosmed(data entities.SocialMedia) (entities.SocialMedia, error) {
	err := r.db.Create(&data).Error
	if err != nil {
		return entities.SocialMedia{}, err
	}
	return data, nil
}

func (r *socialMediaRepository) UpdateSosmed(data entities.SocialMedia) (entities.SocialMedia, error) {
	err := r.db.Updates(&data).First(&data).Error
	if err != nil {
		return entities.SocialMedia{}, err
	}
	return data, nil
}

func (r *socialMediaRepository) DeleteSosmed(socialMediaID uint) error {
	sosmed := entities.SocialMedia{}
	sosmed.ID = uint(socialMediaID)
	err := r.db.First(&sosmed).Where("id = ?", socialMediaID).Delete(&sosmed).Error
	if err != nil {
		return err
	}
	return nil
}