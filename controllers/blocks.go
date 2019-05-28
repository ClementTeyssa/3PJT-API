package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ClementTeyssa/3PJT-API/helper"
	"github.com/ClementTeyssa/3PJT-API/models"
)

func BlocksIndex(w http.ResponseWriter, r *http.Request) {
	helper.LogRequest(r)
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.AllBlocks())
}

func BlocksCreate(w http.ResponseWriter, r *http.Request) {
	helper.LogRequest(r)
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		helper.ErrorHandlerHttpRespond(w, "ioutil.ReadAll(r.Body)")
		return
	}

	var block models.Block
	// var transaction models.Transaction
	err = json.Unmarshal(body, &block)
	if err != nil {
		helper.ErrorHandlerHttpRespond(w, "json.Unmarshal(body, &block)")
		return
	}

	if models.CountTransactionsById(block.TransactionID) != 1 {
		helper.ErrorHandlerHttpRespond(w, "Transaction doesn't exist")
		return
	}

	models.NewBlock(&block)

	json.NewEncoder(w).Encode(block)
}
