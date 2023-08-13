package handlers

import (
	"Auth_service-and-Api_gateway/api/config"
	"Auth_service-and-Api_gateway/api/pkg/grpc_client"
	"Auth_service-and-Api_gateway/proto"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type handlerV1 struct {
	grpcClient *grpc_client.GrpcClient
	cfg        config.Config
}

//HandlerV1Config ...
type HandlerV1Config struct {
	GrpcClient *grpc_client.GrpcClient
	Cfg        config.Config
}

func New(c *HandlerV1Config) *handlerV1 {
	return &handlerV1{

		grpcClient: c.GrpcClient,
		cfg:        c.Cfg,
	}
}

func (h *handlerV1) CreateUser(c *gin.Context) {
	var (
		payload proto.UserRequest
	)

	err := c.ShouldBindJSON(&payload)
	if err != nil {
	}

	res, err := h.grpcClient.UserService().Create(
		context.Background(),
		&payload,
	)
	if err != nil {

	}

	c.JSON(http.StatusOK, res)
}
