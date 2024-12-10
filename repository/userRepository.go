package repository

import (
	"bank-simulation/models"
	"database/sql"
	"errors"
	"strconv"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

var users = []models.User{}

func (repo *UserRepository) GetAllUsers() ([]models.User, error) {

	rows, err := repo.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, errors.New("An error occured")
	}

	var userList []models.User

	for rows.Next() {

		var user models.User
		err := rows.Scan(&user.Id, &user.Name, &user.Email)
		if err != nil {
			return nil, errors.New("An error occured")
		}

		userList = append(userList, user)
	}
	if err := rows.Err(); err != nil {
		return nil, errors.New("An error occured")
	}

	return userList, nil
}

func (repo *UserRepository) GetUserById(id string) (*models.User, error) {

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return nil, errors.New("Id error")
	}

	for _, u := range users {
		if u.Id == idInt {
			return &u, nil
		}
	}

	return nil, errors.New("404 not found user")
}

func (repo *UserRepository) StoreUser(u models.User) error {

	_, err := repo.db.Query("INSERT INTO users( name, email) values  (?, ?)", u.Name, u.Email)
	if err != nil {
		return err
	}
	return nil
}
