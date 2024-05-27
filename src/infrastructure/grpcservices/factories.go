package grpcservices

import (
	usersModels "github.com/chack-check/organizations-service/domain/users/models"
	"github.com/chack-check/organizations-service/infrastructure/grpcservices/protousers"
)

func ProtoUserToModel(user *protousers.UserResponse) usersModels.User {
	return usersModels.NewUser(int(user.Id))
}
