package domain

import (
	"hi-story/models"
)

type AdapterRepository interface {

	// untuk user
	CreateUsers(user models.User) error
	GetAll() []models.User
	// GetOneByID(id int) (user models.User, err error)
	// GetOneByEmail(email string) (user models.User, err error)
	// UpdateOneByID(id int, user models.User) error
	// DeleteByID(id int) error
}