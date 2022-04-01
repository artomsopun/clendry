package database

import (
	"fmt"
	"github.com/artomsopun/clendry/clendry-api/internal/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func NewDB(user, password, host, port, name string) *gorm.DB {
	DB, err := gorm.Open(mysql.Open(
		fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			user, password, host, port, name),
	), &gorm.Config{})

	if err != nil {
		log.Panicln(err)
	}

	if err := DB.AutoMigrate(
		&domain.User{}, &domain.Session{}, &domain.FriendRequest{}, &domain.BlockRequest{}, &domain.Message{},
		&domain.Membership{}, &domain.Chat{}, &domain.File{},
	); err != nil {
		log.Panicln(err)
	}

	return DB
}

func GetInstanceDB() *gorm.DB {
	return DB
}
