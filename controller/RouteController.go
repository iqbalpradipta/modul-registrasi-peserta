package controller

import (
	"net/http"

	"github.com/iqbalpradipta/modul-registrasi-peserta/middlewares"
	"github.com/iqbalpradipta/modul-registrasi-peserta/models"
	"github.com/iqbalpradipta/modul-registrasi-peserta/utils"
)

func PublicRoute(w http.ResponseWriter, r *http.Request)  {
	utils.ToJSON(w, "Public Route", http.StatusOK)
}

func ProtectedRoute(w http.ResponseWriter, r *http.Request)  {
	jwtParams, err := middlewares.JWTExtractToken(r)
	if err != nil {
		utils.ToJSON(w, err.Error(), http.StatusUnauthorized)
	}
	email, ok := jwtParams["user_email"].(string)
	if !ok {
		utils.ToJSON(w, "token Invalid", http.StatusUnauthorized)
		return
	}
	user := models.GetUserByEmail(email)
	if user.Id == 0 {
		utils.ToJSON(w, "User Not Found", http.StatusUnauthorized)
		return
	}
	utils.ToJSON(w, user, http.StatusOK)

}