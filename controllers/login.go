package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/ClementTeyssa/3PJT-API/helper"
	"github.com/ClementTeyssa/3PJT-API/models"
	"golang.org/x/crypto/bcrypt"
)

func Register(w http.ResponseWriter, r *http.Request) {
	helper.LogRequest(r)

	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		helper.ErrorHandlerHttpRespond(w, "ioutil.ReadAll(r.Body)")
		return
	}

	var user models.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		helper.ErrorHandlerHttpRespond(w, "json.Unmarshal(body, &user)")
		return
	}

	user, err = UsersCreate(user, w)
	if err != nil {
		helper.ErrorHandlerHttpRespond(w, err.Error())
		return
	}

	json.NewEncoder(w).Encode(user)
}

func Login(w http.ResponseWriter, r *http.Request) {
	helper.LogRequest(r)

	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		helper.ErrorHandlerHttpRespond(w, "ioutil.ReadAll(r.Body)")
		return
	}

	var userParam models.User
	err = json.Unmarshal(body, &userParam)
	if err != nil {
		helper.ErrorHandlerHttpRespond(w, "json.Unmarshal(body, &userParam)")
		return
	}

	if !helper.EmailExist(userParam.Email) {
		helper.ErrorHandlerHttpRespond(w, "Email doesn't exist !")
		return
	}

	password := []byte(userParam.Password)
	err = verifyEmailPassword(userParam.Email, password)
	if err != nil {
		helper.ErrorHandlerHttpRespond(w, err.Error())
		return
	}

	if err != nil {
		helper.ErrorHandlerHttpRespond(w, "bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)")
		return
	}

	var user = models.FindUserByEmail(userParam.Email)

	json.NewEncoder(w).Encode(user)
}

func verifyEmailPassword(email string, password []byte) error {
	var user = models.FindUserByEmail(email)

	if bcrypt.CompareHashAndPassword([]byte(user.Password), password) != nil {
		return errors.New("Email and password doesn't match")
	}
	return nil
}
