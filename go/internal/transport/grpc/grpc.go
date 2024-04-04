package grpc

import (
	"fmt"
	"google.golang.org/grpc"
)

func NewGRPCConnection(host string, port uint64) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		return nil, fmt.Errorf("grpc: failed to open connection: %v", err)
	}

	return conn, nil
}
