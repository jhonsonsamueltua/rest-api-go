package mysql

import (
	"database/sql"

	userRepo "github.com/rest-api-go/pkg/repository/user"
)

type user struct {
	DB *sql.DB
}

func InitUserRepo(db *sql.DB) userRepo.Repository {
	return &user{
		DB: db,
	}
}
