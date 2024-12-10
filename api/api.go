package api

import (
	"bank-simulation/handlers"

	"github.com/gorilla/mux"
)

func SetupRouter(uHandler *handlers.UserHandler, aHandler *handlers.AccountHandler) *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/users", uHandler.GetAllUsers).Methods("GET")
	router.HandleFunc("/users/{id}", uHandler.GetUserById).Methods("GET")
	router.HandleFunc("/users/{id}/accounts", aHandler.GetUserAccounts).Methods("GET")
	router.HandleFunc("/users", uHandler.StoreUser).Methods("POST")

	router.HandleFunc("/accounts", aHandler.GetAllAccounts).Methods("GET")
	router.HandleFunc("/users/{id}/accounts", aHandler.CreateUserAccount).Methods("POST")
	router.HandleFunc("/users/{userId}/accounts/{accountId}/transactions", aHandler.HandleTransactions).Methods("POST")

	return router

}
