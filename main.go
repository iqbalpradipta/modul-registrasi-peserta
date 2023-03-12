package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/iqbalpradipta/modul-registrasi-peserta/config"
	"github.com/iqbalpradipta/modul-registrasi-peserta/models"
	"github.com/iqbalpradipta/modul-registrasi-peserta/route"
)

func Run()  {
	db := config.InitDB()
	if !db.HasTable(&models.User{}){
		db.Debug().CreateTable(&models.User{})
	}
	db.Close()
	listen(5432)
}

func listen(p int) {
	port := fmt.Sprintf(":%d", p)
	fmt.Printf("Listening Port %s...\n", port)
	r := route.NewRouter()
	log.Fatal(http.ListenAndServe(port, route.LoadCors(r)))
}

func main()  {
	Run()
}