package services

import (
	database "github.com/eron97/api-login.git/models/databases"
)

type UserServiceInterface interface {
	CreateUserService(condition string) (string, error)
}

type UserService struct {
	DB database.Database
}

func (s *UserService) CreateUserService() (string, error) {
	_, err := s.DB.Create()
	if err != nil {
		return "Erro ao conenctar no método Create do database", err
	}

	return "Services chamou database.Create com sucesso!", nil
}
