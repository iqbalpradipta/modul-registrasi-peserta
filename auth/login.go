package auth

import (
	"errors"

	"github.com/iqbalpradipta/modul-registrasi-peserta/hash"
	"github.com/iqbalpradipta/modul-registrasi-peserta/models"
	"github.com/iqbalpradipta/modul-registrasi-peserta/middlewares"
)


var ErrUserNF = errors.New("user not found")

func SignIn(email, password string) (string, error) {
	user := models.GetUserByEmail(email)
	if user.Id == 0 {
		return "", ErrUserNF
	}
	err := hash.PswdVerif(user.Password, password)
	if err != nil {
		return "", nil
	}
	token, err := middlewares.GenerateJWT(user)
	if err != nil {
		return "", err 
	}
	return token, nil
}