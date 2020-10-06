package http_test

// import (
// 	"testing"
// 

// 	"github.com/gorilla/securecookie"
// 	"github.com/labstack/echo"
// 	"github.com/gorilla/sessions"

// 	"github.com/rest-api-go/pkg/models"
// 	"github.com/rest-api-go/pkg/usecase/user/mocks"
// 	delivery "github.com/rest-api-go/pkg/delivery/user/http"
// )

// func InitCookieStore() *sessions.CookieStore {
// 	authKey := securecookie.GenerateRandomKey(64)
// 	encryptionKey := securecookie.GenerateRandomKey(32)

// 	store := sessions.NewCookieStore(authKey, encryptionKey)
// 	store.Options.MaxAge = 86400 * 1
// 	store.Options.HttpOnly = true

// 	return store
// }

// func TestRegister(t *testing.T) {
// 	e := echo.New()
// 	cs := InitCookieStore()
// 	user := models.User{
// 		Username: "jhonson",
// 		Password: "hutagaol",
// 		Name:     "Jhonson Hutagaol",
// 	}

// 	userMockUcase := new(mocks.Usecase)
// 	userMockUcase.On("Register", mock.AnythingOfType("models.User")).Return(int64(1), nil).Once()

// 	handler := delivery.InitUserHandler(e, userMockUcase, cs )
// 	err := handler.Register(c)

// }
