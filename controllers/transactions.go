package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ClementTeyssa/3PJT-API/helper"
	"github.com/ClementTeyssa/3PJT-API/models"
)

type BodyPrivate struct {
	Private []byte `json:"privatekey"`
}

type AdressTransac struct {
	Adress string `json:"adress"`
}

type ToReturnTransacs struct {
	Transactions *models.Transactions `json:"transactions"`
}

func TransactionsIndex(w http.ResponseWriter, r *http.Request) {
	helper.LogRequest(r)
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.AllTransactions())
}

func TransactionsCreate(w http.ResponseWriter, r *http.Request) {
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

	if transaction.AccountFrom == transaction.AccountTo {
		helper.ErrorHandlerHttpRespond(w, "You can't send a transaction to yourself !")
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

	//TODO: traiter les retours d'erreurs
	if userFrom.Solde < transaction.Amount {
		helper.ErrorHandlerHttpRespond(w, "Account from doesn't have enough tokens")
		return
	}

	userTo, err := models.FindUserByAdress(transaction.AccountTo)
	if err != nil {
		helper.ErrorHandlerHttpRespond(w, err.Error())
		return
	}
	models.NewTransaction(&transaction)

	userFrom.Solde = userFrom.Solde - transaction.Amount
	userTo.Solde = userTo.Solde + transaction.Amount

	models.UpdateUser(userFrom)
	models.UpdateUser(userTo)

	json.NewEncoder(w).Encode(transaction)
}

func TransactionsShow(w http.ResponseWriter, r *http.Request) {
	helper.LogRequest(r)
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		helper.ErrorHandlerHttpRespond(w, "ioutil.ReadAll(r.Body)")
		return
	}

	var adress AdressTransac
	err = json.Unmarshal(body, &adress)
	if err != nil {
		helper.ErrorHandlerHttpRespond(w, "json.Unmarshal(body, &adress)")
		return
	}

	if models.CountTransactionsByAdress(adress.Adress) <= 0 {
		helper.ErrorHandlerHttpRespond(w, "No transactions for this adress")
		return
	}

	// var transactions *models.Transactions = models.FindTransactionsByAdress(adress.Adress)
	var toReturnTransacs ToReturnTransacs
	toReturnTransacs.Transactions = models.FindTransactionsByAdress(adress.Adress)

	json.NewEncoder(w).Encode(toReturnTransacs)
}
