package http

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"github.com/rest-api-go/pkg/models"
)

func (d *user) Register(c echo.Context) error {
	var resp models.Responses
	resp.Status = models.StatusFailed

	username := c.FormValue("username")
	password := c.FormValue("password")
	name := c.FormValue("name")

	user := models.User{
		Username: username,
		Password: password,
		Name:     name,
	}

	err := d.userUsecase.Register(user)
	if err != nil {
		resp.Data = nil
		resp.Message = err.Error()
		return c.JSON(http.StatusInternalServerError, resp)
	}

	resp.Data = username
	resp.Status = models.StatusSucces
	resp.Message = "Register is successful"
	return c.JSON(http.StatusOK, resp)
}

func (d *user) Login(c echo.Context) error {
	var resp models.Responses
	resp.Status = models.StatusFailed

	username := c.FormValue("username")
	password := c.FormValue("password")

	user, err := d.userUsecase.Login(username, password)
	if err != nil {
		resp.Data = nil
		resp.Status = models.StatusFailed
		resp.Message = err.Error()
		return c.JSON(http.StatusInternalServerError, resp)
	}

	// session
	session, _ := d.cookieStore.Get(c.Request(), "SessionID")
	session.Values["Authenticated"] = true
	session.Save(c.Request(), c.Response())

	user.Password = ""
	resp.Data = user
	resp.Status = models.StatusSucces
	resp.Message = "Login is successful"
	return c.JSON(http.StatusOK, resp)
}

func (d *user) GetDetailUser(c echo.Context) error {
	var resp models.Responses
	resp.Status = models.StatusFailed

	userID := c.Param("userID")
	userIDInt, _ := strconv.ParseInt(userID, 10, 16)

	user, err := d.userUsecase.GetDetailUser(userIDInt)
	if err != nil {
		resp.Data = nil
		resp.Status = models.StatusFailed
		resp.Message = err.Error()
		return c.JSON(http.StatusInternalServerError, resp)
	}
	user.Password = ""
	resp.Data = user

	resp.Status = models.StatusSucces
	resp.Message = "Get User is successful"
	return c.JSON(http.StatusOK, resp)
}

func (d *user) UpdateUser(c echo.Context) error {
	var resp models.Responses
	resp.Status = models.StatusFailed

	userID := c.Param("userID")
	userIDInt, _ := strconv.ParseInt(userID, 10, 16)
	username := c.FormValue("username")
	name := c.FormValue("name")

	user := models.User{
		UserID:   userIDInt,
		Username: username,
		Name:     name,
	}

	err := d.userUsecase.UpdateUser(user)
	if err != nil {
		resp.Data = nil
		resp.Status = models.StatusFailed
		resp.Message = err.Error()
		return c.JSON(http.StatusInternalServerError, resp)
	}

	resp.Data = user
	resp.Status = models.StatusSucces
	resp.Message = "Update User is successful"
	return c.JSON(http.StatusOK, resp)
}
