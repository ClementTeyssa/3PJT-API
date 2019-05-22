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

	"github.com/ClementTeyssa/3PJT-API/helper"
	"github.com/ClementTeyssa/3PJT-API/models"
	"github.com/gorilla/mux"
)

func UsersIndex(w http.ResponseWriter, r *http.Request) {
	helper.LogRequest(r)
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

func errorHandler(w http.ResponseWriter, err string) {
	log.Println("\n---------------ERROR---------------\n" + err + "\n---------------ERROR---------------")
	var error helper.MyError
	error.Error = err
	json.NewEncoder(w).Encode(error)
}

func UsersCreate(w http.ResponseWriter, r *http.Request) {
	helper.LogRequest(r)
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		errorHandler(w, "ioutil.ReadAll(r.Body)")
		return
	}

	var user models.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		errorHandler(w, "json.Unmarshal(body, &user)")
		return
	}

	if emailExist(user.Email) {
		errorHandler(w, "Email already exist !")
		return
	}

	password := []byte(user.Password)
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		errorHandler(w, "bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)")
		return
	}

	user.Password = string(hashedPasswordBytes)

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		errorHandler(w, "rsa.GenerateKey(rand.Reader, 2048)")
		return
	}

	privateEncoded := x509.MarshalPKCS1PrivateKey(privateKey)
	user.PrivateKey = privateEncoded

	adress, err := bcrypt.GenerateFromPassword(privateEncoded, bcrypt.DefaultCost)
	if err != nil {
		errorHandler(w, "bcrypt.GenerateFromPassword(privateEncoded, bcrypt.DefaultCost)")
		return
	}

	user.Adress = string(adress)

	validate := validator.New()
	errValidate := validate.Struct(user)

	if errValidate != nil {
		errorHandler(w, "validate.Struct(user)")
		return
	}

	models.NewUser(&user)
	json.NewEncoder(w).Encode(user)
}

func UsersShow(w http.ResponseWriter, r *http.Request) {
	helper.LogRequest(r)
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
	helper.LogRequest(r)
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
	helper.LogRequest(r)
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
