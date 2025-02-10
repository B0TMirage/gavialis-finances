package handlers

import (
	"github.com/B0TMirage/gavialis-finances/pkg/database"
	models "github.com/B0TMirage/gavialis-finances/pkg/models/users"
	"golang.org/x/crypto/bcrypt"

	_ "github.com/lib/pq"
)

func Login(u *models.User) (*models.User, error) {
	var uFromBD models.User
	err := database.DB.QueryRow("SELECT id, login, password FROM users WHERE login=$1", u.Login).Scan(&uFromBD.ID, &uFromBD.Login, &uFromBD.Password)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(uFromBD.Password), []byte(u.Password))
	if err != nil {
		return nil, err
	}

	return &uFromBD, nil
}

func Register(u *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = database.DB.Exec("INSERT INTO users(login, password) VALUES($1, $2)", u.Login, string(hashedPassword))
	if err != nil {
		return err
	}

	return nil
}
