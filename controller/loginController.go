package controller

import (
	"encoding/json"
	"net/http"

	"github.com/iqbalpradipta/modul-registrasi-peserta/models"
	"github.com/iqbalpradipta/modul-registrasi-peserta/utils"
)

func Login(w http.ResponseWriter, r *http.Request)  {
	body := utils.BParses(r)
	var user models.User
	err := json.Unmarshal(body, &user)
	if err != nil {
		utils.ToJSON(w, err.Error(), http.StatusUnauthorized)
		return
	}
}