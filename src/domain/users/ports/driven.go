package ports

import "github.com/chack-check/organizations-service/domain/users/models"

type UsersPort interface {
	GetById(id int) *models.User
}
