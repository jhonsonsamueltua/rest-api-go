package module_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/rest-api-go/pkg/models"
	"github.com/rest-api-go/pkg/repository/user/mocks"
	usecase "github.com/rest-api-go/pkg/usecase/user/module"
)

func TestRegister(t *testing.T) {
	user := models.User{
		Username: "jhonson",
		Password: "hutagaol",
		Name:     "Jhonson Hutagaol",
	}

	userMockRepo := new(mocks.Repository)

	t.Run("success", func(t *testing.T) {
		userMockRepo.On("Register", mock.AnythingOfType("models.User")).Return(int64(1), nil).Once()

		u := usecase.InitUserUsecase(userMockRepo)
		userID, err := u.Register(user)

		assert.NoError(t, err)
		assert.Equal(t, userID, int64(1))
		userMockRepo.AssertExpectations(t)
	})

	t.Run("error-save-data", func(t *testing.T) {
		userMockRepo.On("Register", mock.AnythingOfType("models.User")).Return(int64(0), errors.New("Error")).Once()

		u := usecase.InitUserUsecase(userMockRepo)
		userID, err := u.Register(user)

		assert.Error(t, err)
		assert.Equal(t, userID, int64(0))
		userMockRepo.AssertExpectations(t)
	})

	// t.Run("error-hash-password", func(t *testing.T) {
	// 	user.Password = ""

	// 	u := usecase.InitUserUsecase(userMockRepo)
	// 	userID, err := u.Register(user)

	// 	assert.Error(t, err)
	// 	assert.Equal(t, userID, int64(0))
	// 	userMockRepo.AssertExpectations(t)
	// })

}
