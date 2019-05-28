package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ClementTeyssa/3PJT-API/helper"
	"github.com/ClementTeyssa/3PJT-API/models"
)

func TransactionsIndex(w http.ResponseWriter, r *http.Request) {
	helper.LogRequest(r)
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.AllTransactions())
}

//TODO: create transaction
func TransactionsCreate(w http.ResponseWriter, r *http.Request) {

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
		helper.ErrorHandlerHttpRespond(w, "json.Unmarshal(body, &user)")
		return
	}

	//TODO: do != verifs

	userFrom := models.FindUserByAdress(transaction.AccountFrom)
	if userFrom == nil {
		helper.ErrorHandlerHttpRespond(w, "Account from doesn't exist")
		return
	}

	if(userFrom.)

	userTo := models.FindUserByAdress(transaction.AccountTo)
	if userTo == nil {
		helper.ErrorHandlerHttpRespond(w, "Account to doesn't exist")
		return
	}

	models.NewTransaction(&transaction)

	json.NewEncoder(w).Encode(transaction)
}
