package main

import (
	"github.com/ClementTeyssa/3PJT-API/controllers"
	"github.com/gorilla/mux"
)

// InitializeRouter initialize a mux router
func InitializeRouter() *mux.Router {
	// StrictSlash is true => redirect /users/ to /users
	router := mux.NewRouter().StrictSlash(true)
	//TODO: to delete
	router.Methods("GET").Path("/users").Name("ListUsers").HandlerFunc(controllers.UsersIndex)
	router.Methods("POST").Path("/login").Name("Login").HandlerFunc(controllers.Login)
	router.Methods("POST").Path("/register").Name("Register").HandlerFunc(controllers.Register)
	// router.Methods("GET").Path("/users/{id}").Name("Show").HandlerFunc(controllers.UsersShow)
	// router.Methods("PUT").Path("/users/{id}").Name("Update").HandlerFunc(controllers.UsersUpdate)
	// router.Methods("DELETE").Path("/users/{id}/{private}").Name("DELETE").HandlerFunc(controllers.UsersDelete)

	router.Methods("GET").Path("/transactions").Name("ListTransactions").HandlerFunc(controllers.TransactionsIndex)
	router.Methods("POST").Path("/transactions").Name("CreateTransaction").HandlerFunc(controllers.TransactionsCreate)
	router.Methods("GET").Path("/transactions/{adress}").Name("ShowTransactions").HandlerFunc(controllers.TransactionsShow)

	router.Methods("POST").Path("/solde").Name("ShowSolde").HandlerFunc(controllers.ShowSolde)

	router.Methods("POST").Path("/transactions/verify").Name("VerifyTransaction").HandlerFunc(controllers.DoVerifications)

	router.Methods("GET").Path("/blocks").Name("ListBlocks").HandlerFunc(controllers.BlocksIndex)
	router.Methods("POST").Path("/blocks").Name("CreateBlocks").HandlerFunc(controllers.BlocksCreate)
	return router
}
