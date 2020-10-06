package mysql_test

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	// sqlmock "github.com/DATA-DOG/go-sqlmock.v1"
	"github.com/stretchr/testify/assert"

	"github.com/rest-api-go/pkg/models"
	userRepo "github.com/rest-api-go/pkg/repository/user/mysql"
)

func TestRegister(t *testing.T) {
	user := models.User{
		Username: "jhonson",
		Password: "hutagaol",
		Name:     "Jhonson Hutagaol",
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	t.Run("success", func(t *testing.T) {
		query := "INSERT INTO users"
		prep := mock.ExpectPrepare(query)
		prep.ExpectExec().WithArgs(user.Username, user.Password, user.Name).WillReturnResult(sqlmock.NewResult(101, 1))

		u := userRepo.InitUserRepo(db)

		userID, err := u.Register(user)
		assert.NoError(t, err)
		assert.Equal(t, userID, int64(101))
	})

	t.Run("prepare-statement-error", func(t *testing.T) {
		query := "BADQUERY"
		prep := mock.ExpectPrepare(query)
		prep.ExpectExec().WithArgs(user.Username, user.Password, user.Name).WillReturnResult(sqlmock.NewResult(101, 1))

		u := userRepo.InitUserRepo(db)

		userID, err := u.Register(user)
		assert.Error(t, err)
		assert.Equal(t, userID, int64(0))
	})

	t.Run("error-save-data", func(t *testing.T) {
		query := "INSERT INTO users"
		prep := mock.ExpectPrepare(query)
		prep.ExpectExec().WithArgs(user.Username, user.Password, user.Name).WillReturnError(errors.New("Error"))

		u := userRepo.InitUserRepo(db)

		userID, err := u.Register(user)
		assert.Error(t, err)
		assert.Equal(t, userID, int64(0))
	})
}
