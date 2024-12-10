package handlers

import (
	"bank-simulation/models"
	"bank-simulation/repository"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	repo *repository.UserRepository
}

func NewUserHandler(repo *repository.UserRepository) *UserHandler {
	return &UserHandler{repo: repo}
}

func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {

	users, err := h.repo.GetAllUsers()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), 500)
		return
	}

	json.NewEncoder(w).Encode(users)
	w.WriteHeader(http.StatusFound)
}

func (h *UserHandler) GetUserById(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]
	user, err := h.repo.GetUserById(id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(user)
	w.WriteHeader(http.StatusFound)
}

func (h *UserHandler) GetUserAccount(w http.ResponseWriter, r *http.Request) {
}

func (h *UserHandler) StoreUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	fmt.Println(&user)

	err := h.repo.StoreUser(user)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
