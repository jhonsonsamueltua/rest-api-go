package user

import (
	"github.com/labstack/echo"
)

type Delivery interface {
	Register(c echo.Context) error
	Login(c echo.Context) error
	GetDetailUser(c echo.Context) error
	UpdateUser(c echo.Context) error
}
