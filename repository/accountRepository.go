package repository

import (
	"bank-simulation/models"
	"database/sql"
	"errors"
	"strconv"
)

type AccountRepository struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

var accounts []models.Account

func GetAllAccounts() ([]models.Account, error) {
	return accounts, nil
}

func (repo *AccountRepository) GetUserAccounts(id string) ([]models.Account, error) {

	idInt, err := strconv.Atoi(id)
	var userAccounts []models.Account

	if err != nil {
		return nil, errors.New("Invalid User ID")
	}

	for _, acc := range accounts {
		if acc.UserId == idInt {
			userAccounts = append(userAccounts, acc)
		}
	}
	return userAccounts, nil
}

func (repo *AccountRepository) CreateUserAccounts(account models.Account) error {

	accounts = append(accounts, account)
	return nil
}

func (repo *AccountRepository) WithdrawFromAccount() {

}

func (repo *AccountRepository) DepositIntoAccount() {

}
