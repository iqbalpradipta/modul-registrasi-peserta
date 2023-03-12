package config

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jinzhu/gorm"
)

const (
	DB_USER = "iqbalpradipta"
	DB_PASS = "mbangg12"
	DB_HOST = "localhost"
	DB_PORT = 5432
	DB_NAME = "loginRegistrasi"
)

func InitDB() *gorm.DB {
	URL := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME)
	db, err := gorm.Open("postgres", URL)
	if err != nil {
		return nil
	}
	return db
}