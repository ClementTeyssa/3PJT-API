package main

import (
	"crypto/rand"
	"crypto/rsa"
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
			log.Fatal(err)
		}

		privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
		if err != nil {
			log.Fatal(err)
		}
		publicKey := &privateKey.PublicKey

		// create test user
		models.NewUser(&models.User{Email: "test@test.fr", Password: string(hashedPasswordBytes), PublicKeyN: publicKey.N.String()})
	}
}

func main() {
	config.DatabaseInit()
	router := InitializeRouter()

	// create a test user if it doesn't exist
	createTestUser()

	log.Fatal(http.ListenAndServe(":8080", router))
}
