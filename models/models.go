package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID			uint		`gorm:"primary_key:auto_increment" json:"-"`
	FirstName	string		`gorm:"type:varchar(100)" json:"-"`
	LastName	string		`gorm:"type:varchar(100)" json:"-"`
	Email		string		`gorm:"type:varchar(100);unique" json:"-"`
	Password	string		`gorm:"type:varchar(100)" json:"-"`
	Manager		[]Passwords	`gorm:"foreignKey:ID"`
}

type Passwords struct {
	gorm.Model
	ID			uint		`gorm:"primary_key:auto_increment" json:"-"`
	UserEmail	string		`gorm:"type:varchar(100)" json:"-"`
	Username	string		`gorm:"type:varchar(100)" json:"-"`
	Password    string		`gorm:"type:varchar(100)" json:"-"`
	Website		string		`gorm:"type:varchar(200)" json:"-"`
}
