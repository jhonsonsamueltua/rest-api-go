package http

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo"

	userUsecase "github.com/rest-api-go/pkg/usecase/user"
)

type user struct {
	userUsecase userUsecase.Usecase
	cookieStore *sessions.CookieStore
}

func InitUserHandler(e *echo.Echo, u userUsecase.Usecase, c *sessions.CookieStore) {
	handler := &user{
		userUsecase: u,
		cookieStore: c,
	}

	e.POST("/api/user/register", handler.Register)
	e.POST("/api/user/login", handler.Login)
	e.GET("/api/user/:userID", handler.GetDetailUser)
	e.PUT("/api/user/:userID", handler.UpdateUser)
}
