package route

import (
	"github.com/gorilla/mux"
	"github.com/iqbalpradipta/modul-registrasi-peserta/controller"
	"github.com/iqbalpradipta/modul-registrasi-peserta/middlewares"
)

func NewRouter() *mux.Router  {
	r := mux .NewRouter().StrictSlash(true)
	r.HandleFunc("/", controller.PublicRoute).Methods("GET")
	r.HandleFunc("/admin", middlewares.IsAuth(controller.ProtectedRoute)).Methods("GET")
	r.HandleFunc("/users", controller.GetUser).Methods("GET")
	r.HandleFunc("/users", controller.CreateUser).Methods("POST")
	r.HandleFunc("/login", controller.Login).Methods("POST")
	return r

}