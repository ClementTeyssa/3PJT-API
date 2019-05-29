package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/ClementTeyssa/3PJT-API/helper"
	"github.com/ClementTeyssa/3PJT-API/models"
)

type GoodResult struct {
	Good string `json:"good"`
}

func DoVerifications(w http.ResponseWriter, r *http.Request) {
	helper.LogRequest(r)
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		helper.ErrorHandlerHttpRespond(w, "ioutil.ReadAll(r.Body)")
		return
	}

	var transaction models.Transaction
	err = json.Unmarshal(body, &transaction)
	if err != nil {
		helper.ErrorHandlerHttpRespond(w, "json.Unmarshal(body, &transaction)")
		return
	}

	var bodyPrivateKey BodyPrivate
	err = json.Unmarshal(body, &bodyPrivateKey)
	if err != nil {
		helper.ErrorHandlerHttpRespond(w, "json.Unmarshal(body, &bodyPrivateKey)")
		return
	}

	userFrom, err := models.FindUserByAdress(transaction.AccountFrom)
	if err != nil {
		helper.ErrorHandlerHttpRespond(w, err.Error())
		return
	}

	if string(userFrom.PrivateKey) != string(bodyPrivateKey.Private) {
		helper.ErrorHandlerHttpRespond(w, "Private key doesn't match !")
		return
	}

	if userFrom.Solde < transaction.Amount {
		helper.ErrorHandlerHttpRespond(w, "Account from doesn't have enough tokens")
		return
	}

	userTo, err := models.FindUserByAdress(transaction.AccountTo)
	if err != nil || userTo == nil {
		helper.ErrorHandlerHttpRespond(w, err.Error())
		return
	}

	var goodResult GoodResult
	goodResult.Good = "OK"

	log.Println("Transaction from " + transaction.AccountFrom + " ; to " + transaction.AccountTo + " ; for " + strconv.FormatFloat(float64(transaction.Amount), 'f', -1, 32) + " tokens")

	json.NewEncoder(w).Encode(goodResult)
}
