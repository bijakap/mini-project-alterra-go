package domain

import (
	"hi-story/models"
)

type AdapterRepository interface {

	// untuk user
	CreateUsers(user models.User) error
	GetAll() []models.User
	// GetOneByID(id int) (user models.User, err error)
	GetOneByEmail(email string) (user models.User, err error)
	// UpdateOneByID(id int, user models.User) error
	// DeleteByID(id int) error
}

type AdapterService interface {
	//untuk user
	CreateUserService(user models.User) error
	// UpdateUserService(id, idToken int, user model.User) error
	GetAllUsersService() []models.User
	// GetUserByID(id int) (model.User, error)
	AuthUser(email, password string) (string, int, models.User)
	// DeleteByID(id int) error
}

type AdapterRepositoryUlasan interface {
	//untuk Ulasan
	CreateUlasan(id_user, id_museum int, ulasan, image string) error
	GetUlasanByIdMuseum(id_museum int) []models.Ulasan
}

type AdapterServiceUlasan interface {
	//untuk Ulasan
	CreateUlasanService(id_user, id_museum int, ulasan, image string) error
	GetUlasanByIdMuseumService(id_museum int) []models.Ulasan
}

type AdapterRepositoryMuseum interface {
	CreateMuseum(museum models.Museum) error
	GetMuseumById(id int) models.Museum
}

type AdapterServiceMuseum interface {
	CreateMuseumService(museum models.Museum) error
	GetMuseumByIdService(id int) models.Museum
}

type AdapterRepositoryAlbum interface {
	
}

type AdapterServiceAlbum interface {

}