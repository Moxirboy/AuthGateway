package main

import (
	"Auth_service-and-Api_gateway/api/config"
	"Auth_service-and-Api_gateway/api/http"
	"Auth_service-and-Api_gateway/api/pkg/grpc_client"
)

var (
	cfg        config.Config
	grpcClient *grpc_client.GrpcClient
)

func main() {

	cfg = config.Load()

	grpcClient, err := grpc_client.New(cfg)
	if err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}

	server := http.New(http.Config{

		GrpcClient: grpcClient,
	})
	err = server.Run(cfg.HttpPort)
	if err != nil {
		panic(err)
	}

}
