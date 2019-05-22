package helper

import (
	"log"
	"net/http"
)

type MyError struct {
	Error string `json:"error"`
}

func LogRequest(r *http.Request) {
	log.Println("Request to " + r.URL.String() + " with " + r.Method + " method")
}
