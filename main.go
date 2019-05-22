package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/ClementTeyssa/3PJT-API/config"
	"github.com/ClementTeyssa/3PJT-API/models"
)

func createTestUser() {
	if models.UsersSize() <= 0 {
		password := []byte("test")
		hashedPasswordBytes, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
		if err != nil {
			log.Panic(err)
		}

		privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
		if err != nil {
			log.Panic(err)
		}

		privateEncoded := x509.MarshalPKCS1PrivateKey(privateKey)

		adress, err := bcrypt.GenerateFromPassword(privateEncoded, bcrypt.DefaultCost)
		if err != nil {
			log.Panic(err)
		}

		// create test user
		models.NewUser(&models.User{Email: "test@test.fr", Password: string(hashedPasswordBytes), Adress: string(adress), PrivateKey: privateEncoded})
		log.Println("Test user created")
	}
}

func main() {
	config.DatabaseInit()
	log.Println("Database initialised")
	router := InitializeRouter()
	log.Println("Rooter initialised")
	// create a test user if it doesn't exist
	createTestUser()

	log.Panic(http.ListenAndServe(":8180", router))
}
