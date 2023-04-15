package service

import (
	"myGram/entities"
	"myGram/repository"
	"myGram/helpers"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Create(data *entities.User) error
	Login(data *entities.RequestLogin) (entities.ResponseLogin, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) Create(data *entities.User) error {
	err := s.repo.Create(*data)
	if err != nil {
		return err
	}
	return nil
}
func (s *userService) Login(data *entities.RequestLogin) (entities.ResponseLogin, error) {
	dataUser, err := s.repo.Login(data.Email)
	if err != nil {
		return entities.ResponseLogin{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(dataUser.Password), []byte(data.Password))
	if err != nil {
		return entities.ResponseLogin{}, err
	}

	token := helpers.GenerateToken(dataUser.ID, dataUser.Email)
	if err != nil {
		return entities.ResponseLogin{}, err
	}

	resp := entities.ResponseLogin{}
	resp.Token = token

	return resp, nil
}
