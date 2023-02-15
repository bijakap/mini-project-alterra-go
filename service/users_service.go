package service

import (
	"hi-story/domain"
	"hi-story/helper"
	"hi-story/models"
	"net/http"
	"os"

	"github.com/subosito/gotenv"
)

type svcUser struct {
	repo domain.AdapterRepository
}


func (s *svcUser) CreateUserService(user models.User) error {
	return s.repo.CreateUsers(user)
}

func (s *svcUser) GetAllUsersService() []models.User {
	return s.repo.GetAll()
}

func (s *svcUser) AuthUser(email, password string) (string, int, models.User){
	if (email == "" || password == ""){
		return "", http.StatusBadRequest, models.User{}
	}	

	user, err := s.repo.GetOneByEmail(email)
	if err != nil {
		if (err.Error() == "not found"){
			return "", http.StatusNotFound, user
		}
		return "", http.StatusInternalServerError, user
	} else if user.Password != password {
		return "", http.StatusBadRequest, user
	}

	gotenv.Load()
	token, err := helper.CreateToken(int(user.ID), user.Email, os.Getenv("API_SECRET"))
	if err != nil {
		return "", http.StatusUnauthorized, user
	}

	return token, http.StatusOK, user
}

func NewServiceUser(repo domain.AdapterRepository) domain.AdapterService {
	return &svcUser {
		repo: repo,
	}
}