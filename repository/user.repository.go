package repository

import (

	"github.com/nothyphen/Password-Manager/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	AddPassword(password models.Passwords) (string, error)
	ListPassword(email string) (interface{}, error)
}

type userRepository struct {
	conn *gorm.DB
}


func NewUserRepository(connection *gorm.DB) UserRepository {
	return &userRepository{
		conn: connection,
	}
}

func (a userRepository) AddPassword(password models.Passwords) (string, error) {
	if password.Website == "" {
		return "enter the website", nil
	}

	if password.Password == "" {
		return "enter the password", nil
	}

	a.conn.Save(&password)
	return "save password", nil
}

func (a userRepository) ListPassword(email string) (interface{}, error) {
	var lists      []models.Passwords
	var list	   []map[string]interface{}

	result := a.conn.Where("user_email = ?", email).Find(&lists)
	if result.Error != nil {
		return "", result.Error
	}

	for _, user := range lists {
		data1 := map[string]interface{}{
			"id" : user.ID,
			"website" : user.Website,
			"username" : user.Username,
			"password" : user.Password,
		}

		list = append(list, data1)
	}

	return list, nil
}