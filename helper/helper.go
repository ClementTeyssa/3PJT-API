package helper

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ClementTeyssa/3PJT-API/models"
)

type MyError struct {
	Error string `json:"error"`
}

func LogRequest(r *http.Request) {
	log.Println("Request to " + r.URL.String() + " with " + r.Method + " method")
}

func EmailExist(email string) bool {
	if models.UserWithEmailSize(email) > 0 {
		return true
	} else {
		return false
	}
}

func ErrorHandlerHttpRespond(w http.ResponseWriter, err string) {
	log.Println("\n---------------ERROR---------------\n" + err + "\n---------------ERROR---------------")
	var error MyError
	error.Error = err
	json.NewEncoder(w).Encode(error)
}
