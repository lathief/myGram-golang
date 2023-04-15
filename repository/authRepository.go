package repository

import (
	"myGram/entities"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(data entities.User) error
	Login(email string) (entities.User, error)
	// 	Get() ([]entities.User, error)
	// 	GetById(id uint) (entities.User, error)
	// 	Update(data entities.User) (entities.User, error)
	// 	Delete(userID uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(data entities.User) error {
	err := r.db.Create(&data).Error
	if err != nil {
		return err
	}

	return nil
}
func (r *userRepository) Login(email string) (entities.User, error) {
	user := new(entities.User)
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return entities.User{}, err
	}

	return *user, nil
}
