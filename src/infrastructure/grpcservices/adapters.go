package grpcservices

import (
	"context"

	usersModels "github.com/chack-check/organizations-service/domain/users/models"
	"github.com/chack-check/organizations-service/infrastructure/grpcservices/protousers"
)

type UsersAdapter struct {
	client protousers.UsersClient
}

func (adapter UsersAdapter) GetById(userId int) *usersModels.User {
	user, err := adapter.client.GetUserById(context.Background(), &protousers.GetUserByIdRequest{Id: int32(userId)})
	if err != nil {
		return nil
	}

	response := ProtoUserToModel(user)
	return &response
}

func NewUsersAdapter() UsersAdapter {
	return UsersAdapter{
		client: protousers.UsersClientConnect(),
	}
}
