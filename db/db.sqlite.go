package db

import (
	"github.com/nothyphen/Password-Manager/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("database.sqlite"), &gorm.Config{})
	if err != nil {
		panic("can not connect database")
	}

	err = db.AutoMigrate(models.User{})
	if err != nil {
		panic("can not create users table")
	}

	err = db.AutoMigrate(models.Passwords{})
	if err != nil {
		panic("can not create passwords table")
	}

	return db
}

func CloseDB(db *gorm.DB) {
	dbsql, err := db.DB()
	if err != nil {
		panic("can not connect database")
	}

	err = dbsql.Close()
	if err != nil {
		panic("can not close database")
	}
}