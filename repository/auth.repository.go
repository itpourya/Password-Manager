package repository

import (

	"github.com/nothyphen/Password-Manager/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthRepository interface {
	AddUser(user models.User) (string, error)
	FindByEmail(email string) (string, error)
	FindByID(id int64) (string, error)
}

type authRepository struct {
	conn *gorm.DB
}

func NewAuthRepository(connection *gorm.DB) AuthRepository {
	return &authRepository{
		conn: connection,
	}
}

func (a *authRepository) AddUser(user models.User) (string, error) {
	if user.Password != "" {
		user.Password = hashAndsalt(user.Password)
	}

	a.conn.Save(&user)
	return "insert to database", nil
}

func (a *authRepository) FindByEmail(email string) (string, error){
	var user models.User
	query := a.conn.Where("email = ?", email).Take(&user)
	if query.Error != nil{
		return "", query.Error
	}

	return user.Password, nil
}

func (a *authRepository) FindByID(id int64) (string, error){
	var user models.User
	query := a.conn.Where("id = ?", id).Take(&user)
	if query.Error != nil{
		return "error", query.Error
	}
	return "user has exists", nil
}

func hashAndsalt(password string) string {
	pwd := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		panic("can not hash password")
	}

	return string(hash)
}