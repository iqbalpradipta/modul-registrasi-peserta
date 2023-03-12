package auth

import (
	"github.com/iqbalpradipta/modul-registrasi-peserta/models"
)


UserNF := errors.New("User not found !")

func SignIn(email, password string) (string, error) {
	user := models.Get	
}