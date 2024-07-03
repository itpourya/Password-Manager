package services

import (
	"errors"
	"log"

	"github.com/nothyphen/Password-Manager/models"
	"github.com/nothyphen/Password-Manager/repository"
	"github.com/nothyphen/Password-Manager/serilizers"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	AddUserService(registerRequest serilizers.RegisterRequest) (string, error)
	LoginVerify(email string, password string) (string, error)
}

type authService struct {
	authRepository repository.AuthRepository
}

func NewAuthService(repository repository.AuthRepository) AuthService {
	return &authService{
		authRepository: repository,
	}
}

func (a *authService) AddUserService(registerRequest serilizers.RegisterRequest) (string, error) {
	user := models.User{}

	userExists, _ := a.authRepository.FindByEmail(registerRequest.Email)
	if userExists != "" {
		return "", errors.New("user already exists")
	}

	user.Email = registerRequest.Email
	user.FirstName = registerRequest.Firstname
	user.LastName = registerRequest.Lastname
	user.Password = registerRequest.Password

	adduser, err := a.authRepository.AddUser(user)
	if err != nil {
		return "error in mapping", nil
	}
	return adduser, nil
}

func (a authService) LoginVerify(email string, password string) (string, error) {
	user, err := a.authRepository.FindByEmail(email)
	if err != nil {
		return "error", err
	}
	isValidPassword := comparePasswords(user, []byte(password))
	if !isValidPassword {
		errors.New("failed to login, because password is not matched")
	}
	return "Ok", nil

}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {

	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}