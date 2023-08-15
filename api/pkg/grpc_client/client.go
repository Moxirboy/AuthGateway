package grpc_client

import (
	"Auth_service-and-Api_gateway/api/config"
	"Auth_service-and-Api_gateway/proto"
	"fmt"

	"google.golang.org/grpc"
)

//GrpcClientI ...
type GrpcClientI interface {
	// user service
	UserService() proto.AuthClient
}

//GrpcClient ...
type GrpcClient struct {
	cfg         config.Config
	connections map[string]interface{}
}

//New ...
func New(cfg config.Config) (*GrpcClient, error) {

	connUser, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.UserServiceHost, cfg.UserServicePort),
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, fmt.Errorf("user service dial host: %s port:%d err: %s",
			cfg.UserServiceHost, cfg.UserServicePort, err)
	}
	return &GrpcClient{
		cfg: cfg,
		connections: map[string]interface{}{
			// user service
			"user_service": proto.NewAuthClient(connUser),
		},
	}, nil
}

// User new
func (g *GrpcClient) UserService() proto.AuthClient {
	return g.connections["user_service"].(proto.AuthClient)
}
