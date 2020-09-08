package module

import (
	userRepo "github.com/rest-api-go/pkg/repository/user"
	userUsecase "github.com/rest-api-go/pkg/usecase/user"
)

type user struct {
	userRepo userRepo.Repository
}

func InitUserUsecase(r userRepo.Repository) userUsecase.Usecase {
	return &user{
		userRepo: r,
	}
}
