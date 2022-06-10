package utils

import "google.golang.org/grpc"

func MuteCloseClientConn(connection *grpc.ClientConn) {
	if connection != nil {
		_ = connection.Close()
	}
}
