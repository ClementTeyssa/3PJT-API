package controllers

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"golang.org/x/crypto/bcrypt"
	validator "gopkg.in/go-playground/validator.v9"

	"github.com/ClementTeyssa/3PJT-API/models"
	"github.com/gorilla/mux"
)

func UsersIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.AllUsers())
}

func emailExist(email string) bool {
	if models.UserWithEmailSize(email) > 0 {
		return true
	} else {
		return false
	}
}

func UsersCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Panic(err)
		return
	}

	var user models.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Panic(err)
	}

	if emailExist(user.Email) {
		log.Panic("Email already exist !")
	}

	password := []byte(user.Password)
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)

	user.Password = string(hashedPasswordBytes)

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Panic(err)
	}

	privateEncoded := x509.MarshalPKCS1PrivateKey(privateKey)
	user.PrivateKey = privateEncoded

	adress, err := bcrypt.GenerateFromPassword(privateEncoded, bcrypt.DefaultCost)
	if err != nil {
		log.Panic(err)
	}

	user.Adress = string(adress)

	validate := validator.New()
	errValidate := validate.Struct(user)

	if errValidate != nil {
		log.Panic(errValidate)
	}

	models.NewUser(&user)
	json.NewEncoder(w).Encode(user)
}

func UsersShow(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Panic(err)
	}
	user := models.FindUserById(id)
	json.NewEncoder(w).Encode(user)
}

func UsersUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Panic(err)
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Panic(err)
	}
	user := models.FindUserById(id)
	err = json.Unmarshal(body, &user)
	models.UpdateUser(user)
	json.NewEncoder(w).Encode(user)
}

func UsersDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	// strconv.Atoi is shorthand for ParseInt
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Panic(err)
	}
	err = models.DeleteUserById(id)
}
