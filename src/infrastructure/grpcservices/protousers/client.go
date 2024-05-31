package protousers

import (
	"fmt"

	"github.com/chack-check/organizations-service/infrastructure/settings"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"
)

var connection *grpc.ClientConn

func UsersClientConnect() UsersClient {
	if connection == nil || connection.GetState() != connectivity.Ready {
		opts := grpc.WithTransportCredentials(insecure.NewCredentials())
		dsl := fmt.Sprintf("%s:%d", settings.Settings.APP_USERS_GRPC_HOST, settings.Settings.APP_USERS_GRPC_PORT)

		newConnection, err := grpc.NewClient(dsl, opts)
		if err != nil {
			return nil
		}

		connection = newConnection
	}

	return NewUsersClient(connection)
}
