package user

import (
	"github.com/rest-api-go/pkg/models"
)

type Repository interface {
	Register(models.User) (int64, error)
	GetUser(username string) (models.User, error)
	GetDetailUser(userID int64) (models.User, error)
	UpdateUser(models.User) error
}
