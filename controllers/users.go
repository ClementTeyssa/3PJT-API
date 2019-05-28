package controllers

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"errors"
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

func UsersCreate(user models.User, w http.ResponseWriter) (models.User, error) {

	if helper.EmailExist(user.Email) {
		return user, errors.New("Email already exist !")
	}

	password := []byte(user.Password)
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return user, errors.New("bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)")
	}

	user.Password = string(hashedPasswordBytes)

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return user, errors.New("rsa.GenerateKey(rand.Reader, 2048)")
	}

	privateEncoded := x509.MarshalPKCS1PrivateKey(privateKey)
	user.PrivateKey = privateEncoded

	adress, err := bcrypt.GenerateFromPassword(privateEncoded, bcrypt.DefaultCost)
	if err != nil {
		return user, errors.New("bcrypt.GenerateFromPassword(privateEncoded, bcrypt.DefaultCost)")
	}

	user.Adress = string(adress)

	validate := validator.New()
	errValidate := validate.Struct(user)

	if errValidate != nil {
		return user, errors.New("validate.Struct(user)")
	}

	models.NewUser(&user)
	return user, nil
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

//TODO: verify user id ok with private key
// func UsersDelete(w http.ResponseWriter, r *http.Request) {
// 	helper.LogRequest(r)
// 	w.Header().Set("Content-type", "application/json;charset=UTF-8")
// 	w.WriteHeader(http.StatusOK)
// 	vars := mux.Vars(r)
// 	// strconv.Atoi is shorthand for ParseInt
// 	id, err := strconv.Atoi(vars["id"])
// 	if err != nil {
// 		log.Panic(err)
// 	}
// 	private, err := strconv.Atoi(vars["private"])
// 	if err != nil {
// 		log.Panic(err)
// 	}
// 	user := models.FindUserById(id)

// 	err = models.DeleteUserById(id)
// }
