package main

import (
	"github.com/ClementTeyssa/3PJT-API/controllers"
	"github.com/gorilla/mux"
)

// InitializeRouter initialize a mux router
func InitializeRouter() *mux.Router {
	// StrictSlash is true => redirect /users/ to /users
	router := mux.NewRouter().StrictSlash(true)
	router.Methods("GET").Path("/users").Name("Index").HandlerFunc(controllers.UsersIndex)
	router.Methods("POST").Path("/login").Name("Login").HandlerFunc(controllers.Login)
	router.Methods("POST").Path("/register").Name("Register").HandlerFunc(controllers.Register)
	router.Methods("GET").Path("/users/{id}").Name("Show").HandlerFunc(controllers.UsersShow)
	router.Methods("PUT").Path("/users/{id}").Name("Update").HandlerFunc(controllers.UsersUpdate)
	router.Methods("DELETE").Path("/users/{id}").Name("DELETE").HandlerFunc(controllers.UsersDelete)
	return router
}
