package services

import (
	"github.com/nothyphen/Password-Manager/models"
	"github.com/nothyphen/Password-Manager/repository"
	"github.com/nothyphen/Password-Manager/serilizers"
)

type UserService interface {
	SavePassword(saveRequest serilizers.SaveRequest, useremail string) (string, error)
	ListPassword(email string) (interface{}, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func (a userService) SavePassword(saveRequest serilizers.SaveRequest, useremail string) (string, error) {
	var manager models.Passwords

	manager.UserEmail = useremail
	manager.Username = saveRequest.Username
	manager.Password = saveRequest.Password
	manager.Website = saveRequest.Website

	result, err := a.userRepository.AddPassword(manager)
	if err != nil {
		return "", err
	}

	return result, nil
}

func (a userService) ListPassword(email string) (interface{}, error) {
	result, err := a.userRepository.ListPassword(email)
	if err != nil {
		return "", err
	}

	return result, nil
}