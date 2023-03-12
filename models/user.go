package models

import (
	"time"

	"github.com/iqbalpradipta/modul-registrasi-peserta/config"
	"github.com/iqbalpradipta/modul-registrasi-peserta/hash"
)

type User struct {
	Id        uint32     `gorm:"primary_key;auto_increment" json:"id"`
	Username  string     `gorm:"size:20;not null;unique_index" json:"username,omitempty"`
	Email     string     `gorm:"size:35;not null;unique_index" json:"email,omitempty"`
	Password  string     `gorm:"size:60;not null" json:"password,omitempty"`
	CreatedAt *time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt *time.Time `gorm:"default:current_timestamp()" json:"updated_at"`
}

func CreateUser(user User) (interface{}, error) {
	db := config.InitDB()
	defer db.Close()
	hashedPassword, err := hash.PswdHash(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPassword)
	rs := db.Create(&user)
	return rs.Value, rs.Error
}

func GetUsers() []User {
	db := config.InitDB()
	defer db.Close()
	var users []User
	db.Order("id asc").Find(&users)
	return users
}

func GetUserByEmail(email string) User {
	db := config.InitDB()
	defer db.Close()
	var user User
	db.Where("email = ?", email).Find(&user)
	return user
}