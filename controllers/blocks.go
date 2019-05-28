package controllers

import (
	"encoding/json"
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
