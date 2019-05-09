package main

import (
	"github.com/ClementTeyssa/New_Test/config"
	"github.com/ClementTeyssa/New_Test/models"
	"log"
	"net/http"
)

func main() {
	config.DatabaseInit()
	router := InitializeRouter()
	// Populate database
	models.NewUser(&models.User{Manufacturer: "citroen", Design: "ds3", Style: "sport", Doors: 4})
	log.Fatal(http.ListenAndServe(":8080", router))
}