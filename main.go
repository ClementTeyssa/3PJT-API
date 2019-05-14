package main

import (
	"crypto/rand"
	"crypto/rsa"
	"log"
	"net/http"

	"github.com/ClementTeyssa/3PJT-API/config"
	"github.com/ClementTeyssa/3PJT-API/models"
)

func main() {
	config.DatabaseInit()
	router := InitializeRouter()

	// record := strconv.Itoa(block.Index) + block.Timestamp + block.Transaction.AccountFrom + block.Transaction.AccountTo + fmt.Sprintf("%.2f", block.Transaction.Amount) + block.PrevHash
	// h := sha256.New()
	// h.Write([]byte(record))
	// hashed := h.Sum(nil)

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatal(err)
	}
	publicKey := &privateKey.PublicKey

	// Populate database
	models.NewUser(&models.User{Email: "test@test.fr", Password: "test", PublicKeyN: publicKey.N.String()})
	log.Fatal(http.ListenAndServe(":8080", router))
}
