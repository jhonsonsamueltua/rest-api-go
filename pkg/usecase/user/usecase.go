package user

import "github.com/rest-api-go/pkg/models"

type Usecase interface {
	Register(m models.User) (int64, error)
	Login(username string, password string) (models.User, error)
	GetDetailUser(userID int64) (models.User, error)
	UpdateUser(models.User) error
}
