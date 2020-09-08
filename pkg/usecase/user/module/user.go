package module

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/rest-api-go/pkg/models"
)

func (u *user) Register(m models.User) (int64, error) {
	var userID int64
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(m.Password), bcrypt.DefaultCost)
	if len(hashedPassword) != 0 || err == nil {
		m.Password = string(hashedPassword[:])
		userID, err = u.userRepo.Register(m)
		if err != nil {
			return userID, err
		}
	} else {
		return userID, errors.New("Error Hash Password")
	}

	return userID, err
}

func (u *user) Login(username string, password string) (models.User, error) {
	var err error
	users, _ := u.userRepo.GetUser(username)
	err = bcrypt.CompareHashAndPassword([]byte(users.Password), []byte(password))
	if users.UserID == 0 && err != nil {
		return users, errors.New("Username or Password is wrong")
	}

	return users, nil
}

func (u *user) GetDetailUser(userID int64) (models.User, error) {
	user, err := u.userRepo.GetDetailUser(userID)
	return user, err
}

func (u *user) UpdateUser(m models.User) error {
	err := u.userRepo.UpdateUser(m)

	return err
}
