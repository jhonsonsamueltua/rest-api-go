package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo"

	"github.com/rest-api-go/middleware"
	conn "github.com/rest-api-go/pkg/common/connection"
	userDeliver "github.com/rest-api-go/pkg/delivery/user/http"
	userRepo "github.com/rest-api-go/pkg/repository/user/mysql"
	userUseCase "github.com/rest-api-go/pkg/usecase/user/module"
)

func main() {
	//DB
	db := conn.InitDB()
	defer db.Close()
	// cookieStore
	cs := middleware.InitCookieStore()
	//http
	e := echo.New()
	middL := middleware.InitMiddleware(cs)
	e.Use(middL.CORS)
	e.Use(middL.Auth)

	//module
	user(e, db, cs)

	log.Fatal(e.Start(":8000"))
}

func user(e *echo.Echo, db *sql.DB, cs *sessions.CookieStore) {
	userRepo := userRepo.InitUserRepo(db)
	userUsecase := userUseCase.InitUserUsecase(userRepo)
	userDeliver.InitUserHandler(e, userUsecase, cs)
}
