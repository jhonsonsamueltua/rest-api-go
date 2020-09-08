package mysql

import (
	"log"

	"github.com/rest-api-go/pkg/models"
)

func (r *user) Register(m models.User) (int64, error) {
	statement, err := r.DB.Prepare(queryInsertUser)
	if err != nil {
		log.Println("[Repository][User][Register][Prepare] Error : ", err)
		return 0, err
	}

	defer statement.Close()

	res, err := statement.Exec(m.Username, m.Password, m.Name)
	if err != nil {
		log.Println("[Repository][User][Register][Execute] Error : ", err)
		return 0, err
	}
	userID, _ := res.LastInsertId()
	return userID, err
}

func (r *user) GetUser(username string) (models.User, error) {
	var users = models.User{}
	statement, err := r.DB.Prepare(querySelectUser)
	if err != nil {
		log.Println("[Repository][User][GetUser][Prepare] Error : ", err)
		return users, err
	}

	defer statement.Close()

	err = statement.QueryRow(username).Scan(&users.UserID, &users.Username, &users.Password, &users.Name)
	if err != nil {
		log.Println("[Repository][User][GetUser][QueryRow] Error : ", err)
	}

	return users, err
}

func (r *user) GetDetailUser(userID int64) (models.User, error) {
	var users = models.User{}
	statement, err := r.DB.Prepare(querySelectDetailUser)
	if err != nil {
		log.Println("[Repository][GetDetailUser][Prepare] Error : ", err)
		return users, err
	}

	defer statement.Close()

	err = statement.QueryRow(userID).Scan(&users.UserID, &users.Username, &users.Password, &users.Name)
	if err != nil {
		log.Println("[Repository][User][GetDetailUser][QueryRow] Error : ", err)
	}

	return users, err
}

func (r *user) UpdateUser(m models.User) error {
	statement, err := r.DB.Prepare(QueryUpdateUser)
	if err != nil {
		log.Println("[Repository][UpdateUser][Prepare] Error : ", err)
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(m.Username, m.Name, m.UserID)
	if err != nil {
		log.Println("[Repository][UpdateUser][Execute] Error : ", err)
		return err
	}
	return err
}
