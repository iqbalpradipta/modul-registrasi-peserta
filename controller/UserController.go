package controller

import (
	"encoding/json"
	"net/http"

	"github.com/iqbalpradipta/modul-registrasi-peserta/models"
	"github.com/iqbalpradipta/modul-registrasi-peserta/utils"
)

func CreateUser(w http.ResponseWriter, r *http.Request)  {
	body := utils.BParses(r)
	var user models.User
	err := json.Unmarshal(body, &user)
	if err != nil {
		utils.ToJSON(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	rs, err := models.CreateUser(user)
	if err != nil {
		utils.ToJSON(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	utils.ToJSON(w, rs, http.StatusCreated)
}

func GetUser(w http.ResponseWriter, r *http.Request)  {
	users := models.GetUsers()
	utils.ToJSON(w, users, http.StatusCreated)
}