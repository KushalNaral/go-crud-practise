package handlers

import (
	"bank-simulation/models"
	"bank-simulation/repository"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type AccountHandler struct {
	repo *repository.AccountRepository
}

func NewAccountHandler(repo *repository.AccountRepository) *AccountHandler {
	return &AccountHandler{repo: repo}
}

func (acc *AccountHandler) GetAllAccounts(w http.ResponseWriter, r *http.Request) {

	accounts, err := repository.GetAllAccounts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(accounts)
}

func (acc *AccountHandler) GetUserAccounts(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["id"]
	accounts, err := acc.repo.GetUserAccounts(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(accounts)
}

func (acc *AccountHandler) CreateUserAccount(w http.ResponseWriter, r *http.Request) {

	var account models.Account

	userId := mux.Vars(r)["id"]

	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userIdInt, _ := strconv.Atoi(userId)
	account.UserId = userIdInt

	fmt.Println("saving account here : ", account)

	err := acc.repo.CreateUserAccounts(account)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (acc *AccountHandler) HandleTransactions(w http.ResponseWriter, r *http.Request) {

	var params map[string]string
	params = mux.Vars(r)

	userId := params["userId"]
	accountId := params["accountId"]

	fmt.Println("params here", "user id", userId, "account id", accountId)

	//Check user exists
	//Check account exists
	//Check transaction type
	//Case : deposit, add the money to the account
	//Case : withdraw, check if balance is present, if yes withdraw, else err

}
